---
title: 统一error处理
author: Ciusyan
date: '2023-2-25'
---



# GRPC官方的error处理，我放在HTTP中1样的好用



`Go`的error机制，可能被吐槽得算多的了。这里就不吐槽了。既然要来做统一的error处理。

既然要来做统一的`error`处理。在这之前，你总的告诉我，为什么需要做异常处理吧！不可能说因为某某说：~~*一个后端项目中，不能缺少异常处理哟，没有的话，可...*~~


这样可不行，咱们还是来看看，为什么需要统一error处理。



## 一、为什么需要统一error处理？



既然我们主要是处理`HTTP Handler的error`，那先来看看，我认为的，优雅的Handler层的代码长什么样。



### （1）我认为的优雅的Handler层



#### 1、是这样的



先说结论，我认为，Handler层核心应该只做三件事：

1. 处理什么uri
2. 接收请求参数&提取额外参数，去调用业务方法，执行具体的业务逻辑。
3. 拿到上一步调用返回的结果，进行响应。



为什么呢？我的理由如下，听我细细道来。



与其说是处理HTTP的Handler，还不如说是处理HTTP的报文`（HTTP Message）`。简单来了解一下[HTTP的报文格式](https://juejin.cn/post/7117192792446599199)：



![image-20230224151222317](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/86dce1f72aae4ec2861cc3b8396d1f80~tplv-k3u1fbpfcp-zoom-1.image)



看了图，有点懵没关系，我就想说：HTTP 报文格式，核心就是这三部分`（start-line、header、body）`



既然报文长这样，那么HTTP的Handler，核心至少也得处理这三部分吧，用一幅图片简单对应一下：

![image-20230224152826933](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c2ea8844f312440f934eb9113fc8c15e~tplv-k3u1fbpfcp-zoom-1.image)



好了，我的理由叙述完了，再来分别看看我给的结论：

#### 2、处理什么URI



这个是在定义HandlerFunc的时候，就需要定义的，此方法到底负责哪个请求？负责什么URI？比如我们项目使用的是Gin封装的`HTTP Servce`：

```go
// Registry 用于注册Handler所需要暴露的路由
func (h *Handler) Registry(r gin.IRoutes) {
	r.POST("/register/", h.Register)
	r.POST("/login/", h.Login)
	r.GET("/", h.GetUserInfo)
}
```



如上代码所示，在添加路由的时候，不就是在处理这个函数处理什么URI吗？



#### 3、接收请求参数&提取额外参数，去调用业务方法，执行具体的业务逻辑。



比如拿登录这个简单的业务来举例，用户要登陆，至少要提供用户名和密码这两个参数吧。那这两个参数放在哪里呢？：是请求头？请求体？还是query string中呢？

反正不管是哪里，你总的把他从HTTP的报文中解析出来，这不就是在处理Header和Body吗？



拿到了这些参数过后，就可以去调用你的业务方法了。



ps：当然啊，我也见过有一些代码，在Handler里进行了大量的业务逻辑的处理，甚至还有的直接在这里去操纵数据库。不是说不行啊，但的确不太推荐~



#### 4、拿到上一步调用返回的结果，进行响应



经过前两个步骤后，不管调用业务逻辑成功与否，它都会它返回一些数据。那么，我们就需要用这些数据，作为对应的响应。



而响应也是一则HTTP报文，那它也需要组装HTTP中的三个核心：`（start-line、header、body）`



大致就是：进行响应的状态码是什么啊。需不需要添加数据到响应的Header啊。有没有Body数据要返回啊...





### （2）那现在优雅了吗？



好了，介绍完了我认为比较优雅的HTTP Handler代码，我们还是以`Login的HandlerFunc`为例：



```go
func (h *Handler) Login(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		msg := err.Error()
        // 3、若失败、响应错误结果
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		msg := err.Error()
        // 3、若失败、响应错误结果
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}
	resp.StatusCode = 0
	
    // 3、响应正确结果
	c.JSON(http.StatusBadRequest, resp)
}
```



上面的代码，看大致逻辑就好。你觉得优雅了吗？好像满足了我所说的`HTTP的Handler`只做三件事哎：

1. 处理什么uri
2. 接收请求参数&提取额外参数，去调用业务方法，执行具体的业务逻辑。
3. 拿到上一步调用返回的结果，进行响应。



emm，相比于大多数在控制层写了大量业务逻辑的代码，大体上的确还算比较优雅了。但你有没有发现，其中有好几个地方，如果`err != nil`，我们都需要进行形如这样的响应：



```go
	if err != nil {
		msg := err.Error()
        // 3、若失败、响应错误结果
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}
```



谁想有事没事就碰到error啊，我们想要的流程明明是：**接收用户请求 -> 处理业务逻辑 -> 返回成功响应**



可偏偏半路杀出来一个`error`很嘲讽的说：你在调用我的时候出错了，你自己想想怎么处理我吧。要不然我可不让你好过！！！



于是就出现了上面的代码。那我们有没有办法，让他稍微优雅一些呢？



答案是有的，这就引出了我们今天的究极主题：**统一的`error`处理**



## 二、我们是这样做的





### （1）处理前准备一些东西



#### 1、响应信息

一般与外界交互响应信息中，都会有响应码和提示消息。并且我们还可以预定义一些一一对应的code和msg。



1. **定义结构体**



为了方便封装，我们也采用面向对象的方式来做了，定义如下的结构体。

```go
type CodeMsg struct {
   // 状态码
   StatusCode int32 `json:"status_code"`
   // 消息
   StatusMsg string `json:"status_msg"`
}

```



此响应提示，就包含了 code 和 msg，为了方便使用，我们来给一些初始化的构造方法吧



2. **好用的构造方法**



都是些简单的构造方法，就不一一贴出来了，感兴趣的可以看看项目源码鸭~



但是为了之后使用方便，这两个方法还是值得一提。我们刚提到，可以预定义一些的CodeMsg。就像使用Java的枚举类一样。



而这两个方法，就是用来配合预定义的枚举使用的。这两个方法可以双向解析 `code <——> msg`。从而使用预定义的值。



```go
// NewWithMsg 根据msg初始化。
// 注：如果传入的msg不是预定义的msg，那么code：1
func NewWithMsg(msg string) *CodeMsg {
   return New(constant.Msg2Code(msg), msg)
}

// NewWithCode 根据StatusCode初始化。
// 注：如果传入的StatusCode不是预定义的code，那么msg：fmt.Sprintf("未知错误，code = %d", code)
func NewWithCode(code constant.StatusCode) *CodeMsg {
   return NewWithMsg(constant.Code2Msg(code))
}
```



ps：顺带提一嘴，`Go语言`不能写重载方法，有时候真的好难想名字啊。😒😒



#### 2、预定义枚举



本来一开始是直接使用对象来写的枚举，但是后面发现。这些预定义的枚举值，都是一一对应的，用Map做可能会更好。当然，这里借鉴了GRPC的异常处理。往后看看你就知道了。



所以做了形如下面的枚举值。



```go
// 约束枚举值
type StatusCode int32

const (
	OPERATE_OK StatusCode = 0
    
	ERROR_OPERATE        StatusCode = 40001
)

var msgToCode = map[string]StatusCode{
	"操作成功": OPERATE_OK,

	"操作失败":      ERROR_OPERATE,
}
```



#### 3、异常对象



如果出现错误，一般会进行错误的响应，也需要用到上面的枚举，来看看如何定义异常对象的吧。



1. **定义一个异常对象**



主要就是提供CodeMsg，当然，这里模仿`GRPC`的方式，可以让在抛err的时候，携带一些额外的信息`Details`。



```go
type Exception struct {
   S *CodeMsg
   // 可携带额外消息
   Details []interface{} `json:"details"`
}
```



2. 给他构造方法



核心就是根据异常对象，去构建出`CodeMsg`对象，就不一一贴出来了，感兴趣可以看看哟



```go
// 传递预定义的枚举 Msg
func WithStatusMsg(msg string) *custom.Exception {
   return WithCodeMsg(custom.NewWithMsg(msg))
}

// 传递预定义的枚举 Code
func WithStatusCode(code constant.StatusCode) *custom.Exception {
   return WithStatusMsg(constant.Code2Msg(code))
}
```



3. **给他一些好用的方法**



既然想要作为异常对象，核心只需要实现error接口即可，那它都有什么方法呢？



```go
// 标准库的 error 接口
type error interface {
   Error() string
}
```



跟进去看看，呀！只有一个方法，那也太好实现了啊！！！那就好办咯~



```go
// Error 实现 error 接口
func (e *Exception) Error() string {
   return e.S.String()
}
```



至此，我们已经定义好用于自定义error的一些结构体了。可是我们还不知道如何使用这些方法，去处理异常勒。



我们先来看如何统一处理 HTTP Handler 中出现的异常。



### （2）HTTP 出现异常



#### 1、如何处理



因为是用`Gin封装的HTTP Serveice`，所以，我这里仅以处理Gin的Handler出现的异常为例，其余的处理思路类似：



如果按网上所说xx设计模式来说的话，这里的核心思想就是用到了：**装饰器模式**



什么是装饰器模式？不知道没关系，不用管那些高大上的名字，只要知道核心思想，我怎么用不都行，是吧~



而装饰器的核心思想，还是我自己理解的啊：**搞清楚把什么装饰成什么**



而我们这里，要处理HTTTP的异常，那肯定是用来装饰和`HTTP的Handler`相关的东西咯。那先看看`Gin框架封装的HTTP Handler`长什么样：



```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)
```



知道它的HandlerFunc 长这样后，那咱们至少可以得出一个信息：经过装饰后，长成这个样子。



加上Gin的包名，也就是：

![image-20230226165219439](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/73428cabe28246f8bf75ba564da5ab3c~tplv-k3u1fbpfcp-zoom-1.image)



看完了图，如果我们再知道要装饰啥，是不是就可以写代码了。是的，继续跟上我的思路。



如果你再倒回去看上面我们觉得还不够优雅的的代码，你肯定又能想起那个 `err != nil`，那个狰狞的 error 。所以，我们能不能把那个，伤我很深，却又惹不起的`err != nil`的情形，放入装饰器处理呢？按这样的想法，再来画一幅图看看：



![image-20230226170026645](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0acb1582b4ff47ca96a9c3a976eb1644~tplv-k3u1fbpfcp-zoom-1.image)



看完了图，知道了大致思路。那我们就来写代码咯~



```go
// 方便下文使用，自定义一个函数类型
type AppHandler func(c *gin.Context) error

// GinErrWrapper 用于统一处理控制层 error
func GinErrWrapper(handler AppHandler) func(c *gin.Context) {
   return func(c *gin.Context) {
      log := zap.L().Named("GinErrWrapper")
       
      // 1、调用 Handler 方法的逻辑
      err := handler(c)

      defer func() {
         if r := recover(); r != nil {
            log.Infof("Panic: %v", r)
            c.JSON(http.StatusInternalServerError, custom.Bad(constant.INTERNAL))
         }
      }()

      if err == nil {
         return
      }

      log.Errorf("拦截到异常：%s", err.Error())

      // 2、来到这里。err 肯定是有值的咯
      switch e := err.(type) {
      case *custom.Exception:
         c.JSON(http.StatusBadRequest, e.CodeMsg())

         // 还可进行其他 Case 因为我给grpc调用的err用方法GrpcErrWrapper包装了
         // 所以从 控制层的error，基本上都是 custom.Exception
      default:
         c.JSON(http.StatusInternalServerError, custom.Bad(constant.INTERNAL))
      }
   }
}
```



看完了代码，其实也没多难是吧。可以看到：

1. 装饰器的入参和出参，是不是就是再搞清楚，要**把谁装饰成谁**？
2. 刚进入装饰器，我们就去调用了**真正的Handler逻辑**，会返回一个 err 值。如果没有，那就说明正常咯，那就不需要装饰了嘛。
3. 能到后面装饰 error 的逻辑，那说明**肯定有error了**。那我们就可以判断，具体**是哪一种 err**，是自定义的吗？如果是，那就可以取出自定义异常的消息了咯。不是呢？又该怎么办呢？


这下，算是了解思路了吧，来看看处理后，我们是如何使用的。


#### 2、处理后



来看看，现在我们的`HTTP Handler`中，是怎样的，还是拿登录来举例：



```go
// Login 需要经过装饰：r.POST("/login/", exception.GinErrWrapper(h.Login))
func (h *Handler) Login(c *gin.Context) error {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

    // 3、没有err，正常响应
	c.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       custom.NewWithCode(constant.OPERATE_OK),
			TokenResponse: resp,
		})
	return nil
}
```



和之前的代码比起来，是不是就少了很多 err 处理了？不信你往上看看。



至于我们的返回的err值，不就用上了我们自定义的异常对象了吗？当然，你不使用，也是没问题的。甚至你还可以再继续定义其他的异常对象。你只需要在装饰器中，继续**添加 case** 即可



至此，应该了解了我们异常处理的思路。我们趁热打铁，把进行**GRPC调用的异常**也处理一下。



### （3）GRPC出现异常



#### 1、如何处理

这里也是使用装饰器的方式来做的。从上面的使用你应该已经发现了。



```go
// GrpcErrWrapper 用于包装 GPRC 调用产生的 err
func GrpcErrWrapper(err error) *custom.Exception {
   if err == nil {
      return nil
   }
   // 1、通过此方法，一定能够将err转换为 grpc 调用产生的异常
   s := status.Convert(err)

   log := zap.L().Named("GrpcErrWrapper")
   log.Errorf("拦截到异常：%s", s.String())

   // 在通过包装成 自定义异常，统一在 gin 中进行拦截处理
   return WithStatusMsg(s.Message())
}
```



看方法实现，我这里也是给它**转换为了自定义的异常**。然后进入上面HTTP的处理err的装饰器中，统一处理。代码简单，看思路即可~



但是看完了代码，你可能有点疑惑，进行GRPC调用业务方法。如果在编写业务方法时，就以自定义异常`（code + msg）`返回了，那就会返回自定义error鸭，就算error真的 不是 nil，它的值不也会在我们的Handler中吗？不是已经统一处理了吗？为什么还需要我们自己处理一遍呢？



哈哈哈，来碰个杯，我再继续说~🍻。你这简直和我当时的想法一模一样啊。那我为什么还要装饰一遍呢？这就得提到使用GRPC框架进行rpc调用时，它的统一异常处理了。



#### 2、GRPC框架异常处理的大致思路



我这里只是说异常处理的大致思路，并不是说GRPC的原理的哈`（当然，这不主要是因为老弟也不是很了解吗~）`。



那从哪里开始呢？就从它处理一系列逻辑后，返回的 err 说起吧。在进行了一系列的处理后，不管成功与否吧，终于要开始返回err了。



1. 返回的err，会经过此方法进行**包装**：

![image-20230226175216552](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/87b66e32be4c45b5be905863302dc782~tplv-k3u1fbpfcp-zoom-1.image)



看这个方法的思路，是不是感觉有那么一丝丝熟悉的感觉？哈哈哈，是的，我们那里装饰器的思路，就是借鉴了这里的实现。



在经过一系列处理后，返回的err，它都会把他转换成 `status.Error`，也就是说。不管你当时编写业务逻辑的时候，你返回的是什么error。它都会把他变成`status.Error` ，当然，如果你了解到这一点，你可能编写业务逻辑的时候，如果有 err，你就会返回`status.Error`的错误。



就比如当你用`protoc-gen-go-grpc`工具根据`protobuf`定义生成代码的时候。它会有一个默认的`UnimplementedServiceServer`对象。

这个对象会默认实现你定义的所有接口，你看看它的 err 返回值，这不就是`status.Error`吗？

```go
func (UnimplementedServiceServer) Login(context.Context, *LoginAndRegisterRequest) (*TokenResponse, error) {
   return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
```



哈哈哈，所以，你看看，如果你在实现GRPC方法的时候，不是按这样的格式返回 err 的，是不是还要麻烦别人给我们处理一遍呢？



当然，它并不是完全替换掉你的 err 信息，他会根据你的 err 信息，来构建一个 `status.Error`，所以，如果在处理业务逻辑的时候，定义一些信息，会被它中间处理一下。可能会丢失部分信息，比如说，原本定义的状态码。

所以，我们这里在进行GRPC调用的时候，会将返回 err 的异常，用我们自定义的枚举的消息提示，那么之后我们就能够用哪个消息提示，去找出对应的状态码。进而给出正确的提示。

当然啊，也可以将此消息放入status.Error的额外信息中。说到这，你发现我一直提到的 `status.Error`，这玩意到底是啥嘛！



2. status包下面的有什么？



如果你点进它的 status 包中，你会看到好些方法。比如：



![image-20230226180931049](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c0291a7f97d34a2a9f3d42bd61e33aeb~tplv-k3u1fbpfcp-zoom-1.image)



你看看啊，他这些方法，是不是也很眼熟呢？你往上看看，我们定义的结构体，是不是提供了和他很类似的方法？

只不过他这里引入的是`Status`对象，是什么呢？可以看到，是它`internal`内部包中的Status对象。如果你再跟进一步：除了看到那些好用的方法外，最主要有两个结构体：



```go
// Status represents an RPC status code, message, and details.  It is immutable
// and should be created with New, Newf, or FromProto.
type Status struct {
	s *spb.Status
}

// Error wraps a pointer of a status proto. It implements error and Status,
// and a nil *Error should never be returned by this package.
type Error struct {
	s *Status
}
```



你发现，这里的 Status 对象，还依赖了一个 `*spb.Status`对象，如果再跟入，你会发现：



![image-20230226181700271](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/4469c96553e24cb8b4fb1cda1b44fa25~tplv-k3u1fbpfcp-zoom-1.image)



这是官方`Protobuf`生成的代码，看到了什么？是不是也很眼熟？`Code和Massage`。



好了，大致跟你介绍完了这好几个对象。来与我们自定义的异常对象一起看看，总结一下。



3. 总结一下



* 我们的**CodeMsg**对象，和GRPC框架的Status对象很类似，里面主要有`Code+Msg`。
* 我们的**Exception**对象，和这里的`Error`对象很类似，都是去实现了`error`接口。
* 我们的**StatusCode**枚举，和这里面有一个`Code`枚举很类似，都是使用 Map 来定义的。
* 我们**统一异常处理的装饰器**思路，和这里面的 `toRPCErr`很类似。



**所以，至此，我就能说，GRPC官方的error处理，我放在HTTP的Handler中，一样的好用了吧~**




### （4）还是很麻烦的地方



回到文章的开头，为什么说go的error被很多人吐槽，我自己认为的，最核心的点就是：



**error 是就是一个 value**，除非你使用 `panic(err)`，可能会打印一些堆栈信息以外，几乎很难看到堆栈信息。和很多其他的语言相比起来，也就很难定位到问题的根本原因。会增加程序员的调试成本、时间成本。



当然啊，也有很多好的地方，和其他地方被众多coder讨论的。这里只是说说我的个人看法，不喜勿喷哟😁~



所以我认为，这样的异常处理，只是在很大程度上让代码优雅了许多，根本的原因还是需要自己花时间定位。后期可能会想办法加上能够将堆栈信息暴露出来的库。携带上一些根因。尽量减少出现问题的调试成本~

欢迎大家讨论鸭~
