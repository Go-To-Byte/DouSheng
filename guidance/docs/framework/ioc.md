---
title: 花3分钟写的简易IoC，放在Golang的项目中太好用了
author: Ciusyan
date: '2023-2-23'
---

# 花3分钟写的简易IoC，放在Golang的项目中太好用了~



我给朋友`（简称小cher算了🕵️‍♂️🕵️‍♂️）`看了我在我`Go Project`中实现的简易版IOC，实现思路如图所示：



![image-20230223212057924](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302232121034.png)



什么？小cher说她看不懂！！！好吧，我的锅。那我来解释解释，什么是IoC？



## 一、什么是IOC

注：**本文是利用了Ioc的一些思想，实现了简易版的IoC容器**，和IoC的原思想并不完全一致。可以小马过河，因人而异~



~~🕵️‍♂️🕵️‍♂️*小cher说，你用一句话解释I一下`IoC`的核心思想：*~~


一句话解释：将对象利用控制反转的方式，在容器中创建出Bean`（项目的依赖对象）`，并且可以自动为每一个`Bean`注入所需的依赖。

再用一句话解释简易版IoC：将依赖对象注入容器中后，用控制反转的方式，从IoC容器中获取所需要的依赖。



~~🎅🎅*我还补充了一句：*~~

当然，一般从IoC中取出的依赖，也是为了注入给容器中的其它依赖。



~~🕵️‍♂️🕵️‍♂️*这一下小cher就懵了啊：*~~

什么是控制反转？什么又是依赖？注入之后怎么取出来勒？为什么会有IoC勒？



ps：如果这几个概念你很清楚，那咱不浪费大哥们的时间，直接跳第二节去！但是为了照顾像小cher这样cer头cer脑的小朋友，咱们来看看这几个问题~~~



### （1）什么是依赖注入？



~~🎅🎅*小cher啊，我先跟你说官方一点的啊：*~~



#### 1、官方一点的

`DI（Dependency Injection）`称为依赖注入。意思是：如果A实例依赖B实例，如下代码所示：



```go
type A struct {
    // A对象依赖对象 B
    b B
}

type B struct {
    name string
}
```



在程序启动时的时候，会去初始化IoC容器，去初始化对象A的时候，扫描到它需要依赖对象B。IoC会从自身取出B，给A对象中的`b字段`赋值。



~~🕵️‍♂️🕵️‍♂️*我说完这句话后：小cher说：*~~

这个过程，如果给A赋值的时候，B对象还未初始化呢？那之后如果调用A.b的方法，不就相当于`A.nil.xxx`了嘛？



是的，所以我们要求，在进行依赖注入的时候，必须要能够在容器中**能找到被依赖**的对象。



***



~~🕵️‍♂️🕵️‍♂️*说完，小cher加上了自己的思考，已经理解大概的意思了，那还有没有简单一点的描述？*~~



#### 2、简易版

简单说明了上面的那种解释，再来强化一下理解，因为我们要实现的简易版，依赖注入指的是：



如果将程序**刚开始启动**时分为两个阶段

* 阶段一：通过一定的手段，将依赖放入IoC对应的容器中
* 阶段二：去初始化容器中的对象，也就是给容器中每一个对象的属性赋值。



也就是图中的这个过程：

![image-20230223215943567](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302232159618.png)



🕵~~️‍♂️🕵️‍♂️*图还没看完，小cher说，欧，原来这两个步骤就是依赖注入的过程啊！！！这个懂了，但是我还有个疑惑：你说**通过一定的手段**。具体指的是什么呢？*~~



### （2）怎么注入容器？

是啊，怎么通过一定的手段将依赖放入容器呢？这里有几个可借鉴的手段：**可以是配置文件、可以是注解、注释...**



~~🎅🎅*撇开上面这句话，小cher啊，如果有一个盒子，让你把这个球放入一个盒子里，你会怎么放入呢？*~~



~~🕵️‍♂️🕵️‍♂️*我啊，我会大致会有两个主线思路吧：*~~

1. 自己手动放进去咯
2. 跟我女朋友说，当我这个盒子打开的时候，她就把球自动放进去~🐶🐶



~~🎅🎅*？？？满脸羡慕，有女朋友真幸福是吧~ 咳咳，先来继续看这个问题：*~~



是的，你挺会思考的啊。主线思路也就是自动放入和手动放入咯，但是我们作为”高级“程序员，会手动操作吗？



所以，我们注入容器的过程，就选择自动注册，那怎么注册呢？既然我这里是`Go Project`，`Go语言`的程序员围过来，其他语言围观，我们来看一段 [Go官方的Mysql驱动包](https://github.com/go-sql-driver/mysql/) 的代码（主要看标注的）：



![image-20230223221933007](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302232219052.png)



~~🎅🎅*小cher，相信你猜出我想说什么了吧！*~~



~~🕵️‍♂️🕵️‍♂️*那我浅猜一下？你是不是想说：*~~



在程序**刚开始启动**时，如果你的启动入口导入了对应的包，那么就会去加载那个包的一些东西。我在[网上看到了一副图片](https://learnku.com/go/t/47178)，分享给你看看：

![image-20230223223444898](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302232234979.png)



而使用匿名导入，是因为在导入包的地方不需要使用它，只是想去初始化那个包。你看看我说的对吗？？



~~🎅🎅*哇偶，可以啊，小cher你是懂我的！！！来，碰个杯🍻那我接着你的思路说咯*~~



我们知道了这个原理，那么我们就可以将要放入容器的操作写到`init()`函数里，比如下面两段伪代码：

```go
package api

// 这个包有依赖需要放入IoC容器

// ...

func init() {
    // 放入IoC容器的逻辑
    ioc.DI(需要放入的依赖)
}
```



那么在启动程序的时候：

```go
package main

import (
	_ "test/api" // 去初始化 api 包
)

func main() {
    // 启动程序 
    // ....
}
```



~~🕵️‍♂️🕵️‍♂️*小cher看完了这两段代码，豁然开朗，原来是这样啊！！！那我想问问，说控制反转，控制怎么就反转了呢？*~~



### （2）控制怎么就反转了？



~~🎅🎅*好的，小cher，看这个问题之前，你再回想一下，你在项目中一般是怎么使用依赖对象的？最好给我看看简单的代码*~~

~~🕵️‍♂️🕵️‍♂️*呐，你看*~~

```go
// 对象A
type A struct {
   id int
   b  *B // 这里依赖 B 对象
}

// 对象B的构造函数
func NewB(name string) *B {
   return &B{name: name}
}

// 对象B
type B struct {
   name string
}

func main() {
   a := A{
      id: 20,
      b:  NewB("ciusyan"), // 自己控制B对象的初始化
   }
   fmt.Println(a.b.name)
}
```



是啊，你看看，你写的这个，你的A对象依赖B对象，你要是给A对象初始化。你需要自己去写初始化代码。我觉得麻烦的地方就是：

1. 如果你这个类似的代码在100个地方用了，那么你就会写一百遍类似的代码。如果你的参数突然变化了，那么你又要到那用到的一百个地方修改代码
2. 需要自己理清所依赖的对象



~~🕵️‍♂️🕵️‍♂️*也是偶，那我可以封装一下啊！！！哎，不对，你不就是在利用`IoC`的思想封装吗*~~



是的，我们使用IoC的方式封装后（你先别管具体怎么封装的）：如何初始化的操作，都被放入容器中了。需要使用对象的时候，直接从容器中取出来即可，比如：

```go
func main() {
    // 从IOC中获取A对象
   a := ioc.Get("A")
    // 它是如何初始化A和B的，我们根本不需要关心
   fmt.Println(a.b.name)
}
```



再给你[找一副图](https://umajs.github.io/%E5%9F%BA%E7%A1%80%E5%8A%9F%E8%83%BD/IOC.html#%E4%BB%80%E4%B9%88%E6%98%AF-ioc)看看：

![image-20230223232055761](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302232320825.png)



~~🕵️‍♂️🕵️‍♂️*好吧，确实是，用`IoC`封装后，使用起来好方便啊，都不需要自己管理依赖的对象了！！！再看你开始给我看的图，好像清晰了很多*~~



![image-20230223235532757](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302232355799.png)



相信你看到这里，你大概也知道为什么会出现IoC了，那我们再来总结一下~



## 二、为什么会有IOC



一句话解释：方便管理项目的依赖的对象。



刚刚所述的，如果很多重复性很大的代码，那一个点咱们不在重复了，我们下面来看看，如果一个对象的依赖很多，那么你可能去理清这个对象的依赖，会很麻烦。



### （1）没有IoC时



比如下面没有IoC时的一段伪代码：



```java
// 对象A
type A struct {
   id int
   b  *B // 这里依赖 B 对象
   c  *C // 这里依赖 C 对象
   d  *D // 这里依赖 D 对象
   e  *E // 这里依赖 E 对象
   f  *F // 这里依赖 F 对象
   g  *G // 这里依赖 G 对象
   h  *H // 这里依赖 H 对象
   // ....
}

// 使用对象A
a := A{
      id: 20,
      b:  NewB(), // 自己控制B对象的初始化
      c:  NewC(), // 自己控制C对象的初始化
      d:  NewD(), // 自己控制D对象的初始化
      e:  NewE(), // 自己控制E对象的初始化
      f:  NewF(), // 自己控制F对象的初始化
      // ....
   }

```



~~🕵️‍♂️🕵️‍♂️*小cher看完了这段代码，笑道：之前对象的依赖没那么多，好像确实没多少感觉，看你这个依赖多了后，好像是有点麻烦勒*~~



### （2）有IoC之后



看完了上面的一段伪代码，我想，你肯定觉得很ex吧！那来看看有IoC之后呢？

```go
// 定义这些对象的时候，将这些对象全部注入IoC容器中
ioc.DI(A, B, C, D, E, F, G, H)

// 使用对象A
a := ioc.Get("A")
```



当然啊，并不是说没有IoC就不行啊，但若你还没有想到好的方案，不妨可以试试用IoC容器化的思想去封装一些东西？



~~🎅🎅*小cher，聊到现在，相信你应该了解了IoC的一些核心概念，那咱们开始制作装依赖的容器，看看3分钟能不能写完吧！*~~



## 三、开始实践了



一般会用很多种容器，来装不一样的依赖。



~~🎅🎅*小cher，问你个问题，你家里的大衣柜，你不可能把所有衣物，一啪啦的全扔进去吧。*~~



~~🕵️‍♂️🕵️‍♂️*肯定不会啊，要不然到时候取的时候太麻烦了，可以分开管理一下的*~~



是我的话，我可能会把它分成很多个装衣服的容器：有装衬衫的、装领带的、装西装、装裤子...



~~🎅🎅*哈哈哈，是的，既然你也这样认为。那我们多做几个容器，分开管理怎么样。就先来实现一个用来装内部服务对象的容器吧，其余的也是类似思路！*~~



**// TODO：可先看项目代码，之后再完善代码**

[项目IoC相关代码](https://github.com/Go-To-Byte/DouSheng/tree/main/dou_kit/ioc)

### （1）初阶简易版本
// TODO

### （2）完整简易版本
// TODO

### （3）拓展IOC
// TODO
