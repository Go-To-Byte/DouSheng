---
title: 如何优雅的管理应用的生命周期？我是这样做的
author: Ciusyan
date: '2023-2-21'
---



# 如何优雅的管理应用的生命周期？我是这样做的



## 一、什么时候要注意管理应用的生命周期？



先来看一段代码：（假设无 err 值）



```go
func main() {
    
    // 1、启动HTTP服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
    
    
    // 2、启动GRPC服务
    server := grpc.NewServer()
    listener, _ := net.Listen("tcp", ":1234")
	server.Serve(listener)    
}

```



这一段代码，相信你一眼就能看出问题，因为在启动HTTP后，进程会堵塞住，下面启动GRPC服务的代码，压根就不会执行。



但是，要是想要同时启动`GRPC`服务呢？自己没有时间，那么就请一个帮手咯，让它来为我们启动GRPC服务，而这个帮手，就是go的携程。



* 来看一段伪代码，也就是调整成这样，

```go
func main() {
 
    // 1、将HTTP服务放在后台启动
    go start http
    
    // 2、将GRPC服务放在前台启动
    start grpc  
}
```



但是调整成这样之后，我们想要的情况就是，HTTP成功启动后、GRPC也要启动成功。HTTP意外退出后，GRPC也需要退出服务，他们俩需要共存亡。



但若出现了 HTTP 意外退出、GRPC还未退出，那么就会浪费资源。还可能出现其他的问题。比如接口异常。这样会很危险。那我们该利用什么方式，让同一服务内，启动多个线程。并且让他们共同存亡的呢？



了解了上面的问题，我们再来重新描述总结一下出现的问题。



一个服务，可能会启动多个进程，比如说 HTTP API、GRPC API、服务的注册，这些模块都是独立的，都是需要在程序启动的时候进行启动。



而且如果需要关闭掉这个应用，还需要处理很多关闭的问题。比如说

* HTTP、GRPC 的优雅关闭
* 关闭数据库链接
* 完成注册中心的注销操作
* ...



而且，启动的多个进程间，该如何通信呢？某些服务意外退出了，按理来说要关闭整个应用，该如何监听到呢？



## 二、我们是如何做的



### （1）利用面向对象的方式来管理应用的生命周期



定义一个管理者对象，来管理我们应用所需要启动的所有服务，比如这里需要被我们启动的服务有：HTTP、GRPC



这个管理者核心有两个方法：`start、stop`



```go
// 用于管理服务的开启、和关闭
type manager struct {
	http *protocol.HttpService // HTTP生命周期的结构体[自定义]
	grpc *protocol.GRPCService // GRPC生命周期的结构体[自定义]
	l    logger.Logger		   // 日志对象
}
```



不用关心这里依赖的 `http、grpc`结构体是什么，我们在后面的章节，会详细解释。只需要知道，我们用`manager`这个结构体，用于管理`http、grpc`服务即可。





### （2）处理start



`start`这个函数，核心只做了两件事，分别启动`HTTP、GRPC`服务。



```go
func (m *manager) start() error {

	// 打印加载好的服务
	m.l.Infof("已加载的 [Internal] 服务: %s", ioc.ExistingInternalDependencies())
	m.l.Infof("已加载的 [GRPC] 服务: %s", ioc.ExistingGrpcDependencies())
	m.l.Infof("已加载的 [HTTP] 服务: %s", ioc.ExistingGinDependencies())

	// 如果不需要启动HTTP服务，需要才启动HTTP服务
	if m.http != nil {
		// 将HTTP放在后台跑
		go func() {
			// 注：这属于正常关闭："http: Server closed"
			if err := m.http.Start(); err != nil && err.Error() != "http: Server closed" {
				return
			}
		}()
	}

    // 将GRPC放入前台启动
	m.grpc.Start()
	return nil
}
```



又因为开头说过了，启动这两任一服务，都会将进程堵塞住。



所以我们找了一个帮手`（携程）`来启动`HTTP`服务，然后将`GRPC`服务放在前台运行。



那为什么我要将`GRPC`服务放在前台运行呢？其实理论上放谁都行，但由于我们的架构原因。我们有的服务不需要启动`HTTP`服务，而每一个服务都会启动`GRPC`服务。所以，将GRPC放置在前台，会更合适。



至于里面如何使用`HTTP、GRPC`的服务对象启动它们的服务。在这一节就不多赘述了。在之后的章节会有详细的介绍~



看完了统一管理启动的`start`方法，那我们来看看如何停止服务吧



### （3）处理stop



#### 1、什么时候才去Stop？



我们开启了多个服务，并且有的还是放在后台运行的。这就涉及到了多个携程的间通信的问题了



用什么来通信吶？我怎么知道`HTTP`服务挂没挂？是意外挂的还是主动挂的？我们怎么能够优雅的统一关闭所有服务呢？



其实这一切的问题，`Go`都为我们想好了：那就是使用`Channels`。一个`channel`是一个通信机制，它可以让一个携程通过它给另一个携程发送值信息。每个`channel`都有一个特殊的类型，也就是`channels`可发送数据的类型。



我们把一个`go程`当作一个人的化，那么`main` 方法启动的主`go程`就是你自己。在你的程序中使用到的其他`go程`，都是你的好帮手，你的好朋友，它们有给你去处理耗时逻辑的、有给你去执行业务无关的切面逻辑的。而且是你的好帮手，按理来说最好是由你自己去决定，要不要请一个好帮手。



当你请来了一个好帮手后，它们会在你的背后为你做你让他们做的事情。那么多个人之间的通信，比较现代的方法，那可以是：打个电话？发个消息？所以用到了一个沟通的信道：`Channel`



好了，当你了解了这些后，也就是接收到一些电话后，我们才需要去`stop`。我们再回到Dousheng使用的情景：



#### 2、Dousheng的应用场景



主携程是`GRPC`服务这个人，我们请了一个帮手，给我启动HTTP服务。这个时候，如果HTTP服务这个帮手意外出事了。既然是帮我么你做事，那我们肯定得对别人负责是吧。但是我们也不知道它出不出意外啊，怎么办呢？这时候你想了两个方法：



1. 跟你的帮手HTTP，发了如下消息

![image-20230221190716875](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302211907029.png)





这就需要HTTP自己告诉我们，按理来说，应该是可以的。但是如果HTTP遇到了重大问题，根本来不及告诉我们呢？咱们又是一个负责的男人。为了避免这种情况发生，又请一个人，专门给我们看HTTP有没有遇到重大问题。于是有了第二种方式：



2. 在请一个帮手`signal.Notify`，帮助我们监听HTTP可能会遇到的重大问题

![image-20230221192046573](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302211920621.png)



就这样，我们就几乎不怕出现，HTTP出事了，我们还不知道的情况了。当我们收到HTTP出事的信号后，那我们就可以统一的去优雅关闭服务了



相信你已经了解了核心的思想，我们来看看，用代码该如何实现



#### 3、代码实现



* 启动`signal.Notify`，用于监听系统信号



我们已经分析过了，我们需要再请一个帮手，来给我们处理HTTP可能会遇到的重大事故：`（syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT）`



```go
// WaitSign 等待退出的信号，实现优雅退出
func (m *manager) waitSign() {
   // 用于接收信号的信道
   ch := make(chan os.Signal, 1)
   // 接收这几种信号
   signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)

   // 需要在后台等待关闭
   go m.waitStop(ch)
}
```



当`signal.Notify`收到上面所列举的信号后，那么就可以去做关闭的事情了，那如何关闭呢？



* 读取信号，执行优雅关闭逻辑

```go
// WaitStop 中断信号，比如Terminal [关闭服务的方法]
func (m *manager) waitStop(ch <-chan os.Signal) {

   // 等待信号，若收到了，我们进行服务统一的关闭
   for v := range ch {
      switch v {
      default:
         m.l.Infof("接受到信号：%s", v)

         // 优雅关闭HTTP服务
         if m.http != nil {
            if err := m.http.Stop(); err != nil {
               m.l.Errorf("优雅关闭 [HTTP] 服务出错：%s", err.Error())
            }
         }
          
		// 优雅关闭GRPC服务
         if err := m.grpc.Stop(); err != nil {
            m.l.Errorf("优雅关闭 [GRPC] 服务出错：%s", err.Error())
         }
      }
   }
}
```



这里的逻辑比较简单，就是当接收到信号的时候，对`HTTP、GRPC`做优雅关闭的逻辑。至于为什么要进行优雅关闭，而不是直接`os.Exit()`？我们在下一节讲~



这里值得一提的是，我们从chanel里获取数据，因为我们这里只和单个携程间进行通信了，使用的是 `for range`，并没有使用`for select`



好了，这样我们应用的生命周期算是被我们优雅的拿捏了。我们一直在讲优雅关闭这个词，我们来解释一下什么是优雅关闭？为什么需要优雅关闭？



## 三、什么是优雅关闭



既然HTTP服务和GRPC服务都需要优雅关闭，我们这里用HTTP服务来举例。



先来看这张图，假设有三个并行的请求至我们的HTTP服务。它们都期望得到服务器的`response`。HTTP服务器正常运行的情况下，多半是没问题的。



![image-20230221212855436](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302212128484.png)



请求已发出，若提供的HTTP服务突然异常关闭了呢？我们继续来把HTTP服务比作一个人。看看它是否优雅呢？



### （1）没有优雅关闭



如果HTTP这个人不太优雅，是一个做事不怎么负责的渣男。当自己异常over了之后，也不解决完自己的事情，就让别人`（request）`，找不到资源了。真的很不负责啊。

大致用一幅图表示：

![image-20230222005543821](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302220055924.png)



这个不优雅的HTTP服务，当有还未处理的请求时，自己就异常关闭了，那么它根本不会理会原先的请求是否完成了。它只管自己退出程序。



### （2）有了优雅关闭



看完了那个渣男HTTP`（没有优雅关闭）`，我们简直想骂它了。那我们来看，当一个优雅的谦谦君子`（有优雅关闭）`，又是如何看待这个问题的。



![image-20230222011138886](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302220111933.png)



这是一个负责人的人，为什么说他负责人、说它优雅呢？因为当它自己接收到异常关闭的信号后。它不会只顾自己关闭。它大概还会做两件事：

1. 关闭建立连接的请求通道，防止还会接收到新的请求
2. 处理完以请求的，但是还未响应的请求。保证资源得到响应，哪怕是错误的`response`。



正是因为它主要做了这两件事，我们才说此时的HTTP服务，是一个优雅的谦谦君子。



而当有很多个请求到时候，我们怎么知道是否会不会突然异常关闭呢？如果遇到了这种情况，我们应该处理完未完成的响应，拒绝新的请求建立连接，因为我们是一个优雅的人。



通过这里的讲解，再配合上刚刚的管理者对象，我们就可以优雅的控制应用的生命周期了。那我们下面来了解一下如何将IoC使用在架构中。



