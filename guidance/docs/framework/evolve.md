---
title: 做了2次架构演变，算是入了微服务的门吧~
author: Ciusyan
date: '2023-2-20'
---

# 做了2次架构演变，算是入了微服务的门吧~


## 一、初见Dousheng


### （1）架构思路

因为自己以前是一个`Javer`，对传统的SSM三层架构还比较熟悉。就巨石架构而言，模块还算是比较清晰了。又因为接触了一门新的语言`GoLang`，利用一些熟悉的事物过渡到不太熟悉的领域。所以借鉴了MVC的思想，引入似MVC的架构方式。


也就是：

1. 控制层`（Handler）`：用于控制网络请求。
2. 业务层`（Service）`：用于处理具体业务，还有简单的数据库操作。
3. 持久层`（Dao）`：用于进行数据库的操作。


又因为没有类似Spring的框架来管理依赖，我们这里并没有严格的区分业务层和持久层。所以我们最初的架构是"两层半"：`Handler -> Service + Dao`。



了解了初次架构的设计思路后，我们来看一些直观的表达。



#### 1、架构图

简单画一幅图来表示就是：

![image-20230220003052506](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302200030665.png)



这样比较传统的单体架构，较容易理解，就不多解释了。直接来看看拆分后的目录结构。



#### 2、目录结构

再来看看架构的目录结构，其他的一些辅助包，暂时不需要关注，可查看单独的文档。现在只需要关注业务模块的分层即可。



##### 目录结构概览[解读]

```tex
DouSheng            # 极简版抖音 APP
├── apps            # 所有服务模块
│   ├── all         # 统一管理所有模块实例的注册[驱动加载的方式]
│   ├── comment     # ===评论模块===
│   │   ├── api     # 控制层(Handler)
│   │   ├── impl    # 业务层(Service) + 持久层(Dao)
│   │   └── pb	    # interface 、model 
│   ├── user        # ===用户模块===
│   │   ├── api     
│   │   ├── impl
│   │   └── pb
│   └── video       # ===视频模块===
│       ├── api
│       ├── impl
│       └── pb
├── cmd             # CLI
├── common.pb       # 放置公共的protobuf文件[可抽离]
├── conf            # 项目配置对象
├── docs            # 项目相关文档
├── etc             # 项目具体配置
├── ioc             # IoC容器[可抽离]
├── protocol        # 提供协议
├── utils           # 工具包
└── version         # 版本信息
```



##### 部分主要文件概览[解读]

```tex
├── apps                            # 所有的业务模块
│   ├── all                         # 驱动注册所有的IOC容器实例
│   │   └── auto_register.go
│   ├── user                        # 以用户模块举例
│   │   ├── api                     # 提供的 API 接口
│   │   │   ├── http.go             # 使用 HTTP 的方式暴露 控制层逻辑
│   │   │   └── user.go             # user服务模块暴露的方法
│   │   ├── app.go                  # user模块的结构体方法
│   │   ├── impl                    # user.ServerService 的实现
│   │   │   ├── dao.go              # 可以看作是 持久层逻辑
│   │   │   ├── impl.go             # 可以看作是 业务层逻辑
│   │   │   ├── user.go             # user.ServerService 接口方法的实现
│   │   │   └── user_test.go        # 此模块测试用例【注：必写，一般用于测试本模块CURD的功能】
│   │   ├── pb                      # 此模块的protobuf文件，里面有（接口方法、请求model、响应model、本模块model）
│   │   │   └── user.proto      
│   │   ├── README.md               # 本模块说明
│   │   ├── user.pb.go              # 利用 protoc 生成（结构体）
│   │   └── user_grpc.pb.go         # 利用 protoc 生成（接口）
├── cmd                             # 用于启动项目
│   ├── root.go                     
│   └── start.go                    # 启动逻辑在这
├── common                          # 定义的公共的protobuf文件，可抽离
│   ├── common.pb.go
│   └── pb
│       └── common.proto
├── conf                            # 项目配置对象
│   ├── app.go                      # 此项目的配置
│   ├── config.go                   # 统一配置
│   ├── config_test.go              
│   ├── load.go                     # 加载所有配置
│   ├── log.go                      # 日志相关配置
│   └── mysql.go                    # mysql相关配置
├── etc
│   ├── dousheng.toml               # 项目配置文件位置【可换成其他的，用其他库解析】[禁止上传github]
│   └── dousheng.toml.template      # 配置文件模板[可上传github]
├── ioc                             # IoC容器
│   ├── all.go                      # 统一所有容器
│   ├── gin.go                      # Gin HTTP 服务容器
│   ├── grpc.go                     # GRPC 服务容器
│   └── internal.go                 # 内部服务容器
├── Makefile                        # 利用Makefile管理项目[相当于一个脚手架]
├── utils                           # 放置一些通用的工具
│   └── md5.go  
```



看完了目录结构，应该能很清晰的看出分了三个模块`（user、comment、video）`，并且每一个模块都有自己完全独立的“两层半”架构。



既然还算清晰的对模块进行了划分。那为什么还要演变呢？



### （2）遇到的问题

尽管也是分模块开发，但是最终还是会打包并部署，还是为单体应用。不是说不行，但是可会遇到一些问题：

* 其中最主要的问题就是，这个应用最终会太复杂，以至于任何单个开发者都可能搞不懂它。

* 应用无法扩展、可靠性低，一炸全炸。

* 最终，想要实现应用的敏捷性开发和部署变得很难。



当业务体量不大的时候，单体架构可能会更受人们青睐，也不会引入更多额外的资源、技术复杂度...

但是业务体量、用户体量一旦增长了起来，单体架构很难稳定的抗住冲击。再加上也想了解一下微服务开发。



所以，我们进行了架构的第一次演变...



## 二、第一次演变

### （1）架构思路



人类自古就有**化繁为简、分而治之**的思想，我们可以将一个复杂而庞大的业务，抽象成一个个简单的服务，然后单独的分开处理。我觉得这也是微服务的核心思路。



但是，在每一个单独的服务中，我们还是保留了MVC的”两层半架构“。再来看看一些直观的表达：

#### 1、架构图

我们原先根据业务，对模块进行了垂直划分，然后在划分出来的模块中，进行了水平划分，如下图所示：



![image-20230220020709261](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302200207325.png)



从图中可以发现，拆分出来的每一个服务，我们都用不一样的端口，不一样的进程，运行了起来。对外部提供的服务，通过HTTP的方式暴露出去。而内部服务间的调用，就不再是通过文件路由引用了，而是通过GRPC协议暴露出去。值得一提的是，为了开发方便，每个服务的数据库并没有拆分开。



看完了架构图，我们来看看大致的目录结构。



#### 2、目录结构



##### 总目录结构概览[解读]

还是以用户中心、视频服务、评论服务举例。

```tex
DouSheng
├── dou_kit				         # ===简单的分Kit公共包===
│	.....
├── user_center				     # ===用户服务===
│	.....
└── video_service                # ===视频服务===
│	.....
└── comment_service              # ===评论服务===
│	.....
```





##### 详细一些的结构概览[解读]

这里以用户中心为例，展开目录结构：

```tex
DouSheng
├── dou_kit				         # ===简单的分Kit公共包===
│   ├── conf			         # 配置文件
│   ├── constant		         # 常量
│   ├── docs.sql		         # 部分文档
│   ├── exception			     # 统一error处理
│   └── ioc					     # IOC容器
├── user_center				     # ===用户服务===
│   ├── apps                     # 包含的模块
│   │   ├── token                # token模块
│   │   │   ├── impl
│   │   │   └── pb
│   │   ├── user                 # 用户模块
│   │   │   ├── api
│   │   │   ├── impl
│   │   │   └── pb
│   ├── client.rpc.middlerware   # 用户中心提供的客户端 
│   ├── cmd                      # 命令行工具
│   ├── common                   # 模块内公共工具
│   │   ├── constant
│   │   └── utils
│   ├── docs                     # 模块内文档
│   │   ├── example
│   │   ├── sql
│   │   └── static.image
│   ├── etc                      # 用户中心的配置文件
│   ├── protocol                 # 对外暴露的协议     
│   └── version                  # 用于注入版本信息
└── video_service                # ===视频服务===
│	.....
└── comment_service              # ===评论服务===
│	.....
```



看完了演进后的架构图和目录结构。其实这就是一个简单的微服务拆分了。核心就是化繁为简，分而治之的思想。我们这里仅对项目架构简单说明，很多微服务的知识并未在这一节体现。



这样进行简单的拆分之后，分出了若干服务，并且服务间通过rpc调用，每个服务可以单独部署、单独编写、本来已经解决了单体架构的很多问题了。而且是通过功能模块划分的，更容易理解了。那为什么还有一次架构演进呢？我们又遇到了什么问题呢？



### （2）遇到的问题

我们在这里，首先遇到的问题就是：**对外暴露的接口不统一**，比如官方提供的测试APP，需要配置后端接口的主机地址+端口。只能访问一个进程内的接口。而我们这样的拆分方式，会同时启动很多个对外暴露HTTP服务的进程。若想要完整的通过APP测试，是几乎不可能的事情。



必行之事，何必问天。光是因为上面所述的一个理由，我们的架构，就不得不再一次演变。还不谈会遇到的其他问题。



那我们来看看是如何进行第二次架构演变的。



## 三、第二次演变

### （1）架构思路

“没有什么是加一层解决不了的事情，如果有，那就两层”，相信大家都听过这句话。是啊，我们遇到了上面的问题之后，尝试加入了一层：`Api Rooter`，加入了这一层。

对外暴露的HTTP接口，就可以统一在这一层做了。而由这一层，可以通过GRPC去调用。实际的业务逻辑。



先来看一些较为直观的表达，在继续探讨。



#### 1、架构图



主要呈现的是服务的拆分关系。



![image-20230220130423149](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302201304286.png)



如图所示，对外暴露的HTTP服务，全是经过`Api Rooter`这一层出去的。在这一层，只做两件事情。

1. 管理`Token`的认证
2. 组装Api，对外提供HTTP服务



因为Token相当于是用户的身份凭证，以前是放在用户中心的，现在是放在`Api Rooter`的，因为放在这里，当有请求过来的时候，若需要校验信息，直接调用方法即可。就不需要额外走GRPC去调用`user_center`的方法了。



我们这里其实并没有太多组合Api的接口。我们的接口大多数是已经在内部服务组装好的。然后在这一层直接暴露出去即可。相当于这是各个服务Handler的聚集地。在这里聚集，然后统一暴露给外界。



值得一提的是，这一层，是通过GRPC去调用内部服务的，并不是通过HTTP协议去调用的。主要是因为这是自定义的Api组合层，支持GRPC去调用自己的服务。



#### 2、目录结构

加入了Api这一层、把一些公共模块更进一步的抽离出来后，现在的目录结构是这样的：



```tex
DouSheng
├── .github.workflows
├── api_rooter              # ===简易版网关===
│   ├── apps                
│   │   ├── token           # Token的 RPC Server
│   │   │   ├── impl        
│   │   │   └── pb
│   │   ├── user.api        # 用户中心的HTTP接口
│   │   └── video.api       # 视频服务的HTTP接口
│   ├── client.rpc          # Token的RPC Client
│   ├── common  
│   │   ├── all
│   │   └── utils
│   ├── docs
│   ├── etc
│   └── protocol
├── dou_kit                 # ===封装的公共库===
│   ├── client
│   ├── cmd
│   ├── conf
│   ├── constant
│   ├── docs
│   │   ├── sql
│   │   └── static
│   ├── exception.custom
│   ├── ioc
│   ├── protocol
│   └── version
├── guidance.docs           # ===项目文档===
├── user_center             # ===用户中心===
│   ├── apps.user
│   │   ├── impl
│   │   └── pb
│   ├── client.rpc
│   ├── common
│   │   ├── all
│   │   └── utils
│   ├── docs
│   │   ├── example
│   │   ├── sql
│   │   └── static.image
│   └── etc
└── video_service           # ===视频服务===
    ├── apps.video
    │   ├── impl
    │   └── pb
    ├── client.rpc
    ├── common
    │   ├── all
    │   ├── pb
    │   └── utils
    ├── docs.sql
    ├── etc
    └── store.aliyun
```



在加入这一层后，对外暴露接口的方式、样式、和端口，都统一了。这下就完事了嘛？未来真的不会出问题了吗？



### （2）可能会遇到的问题

我们现在是通过`Api Rooter`来统一暴露接口的。其中最致命的就是整个 `App Rooter` 属于 `single point of failure`，若在这一层出现严重的代码缺陷，或者流量洪峰，可能会引发集群宕机，出现单点故障。这个故障并不是说某一个服务宕机了，而是对外提供的HTTP接口会崩掉。



但是由于一些原因：如项目进度、未学习的知识、技术成本....等问题。目前还没有办法再次演进。所以Dousheng最终的架构，暂定为这样了。



## 四、未来的设想

### 未来架构演进思路

既然每一个API服务太庞大了，那我们继续利用大禹治水，分而治之的思想。将其拆分成多个服务独立的网关小组。这样就算某一服务提供的API宕机了，也不会导致所有服务宕机。也就是解决了单体故障的问题。



在引入一层真正的网关技术`（API Geteway）`，来处理转发用户的请求。而且将一些横切面的逻辑放置到这一层。比如日志监控、安全认证等等



大致画一幅图，也就是这个样子的：

![image-20230220145446543](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302201454623.png)



至此，我们通过两次架构的演进，相信你已经基本掌握了Dousheng的架构思路。也入了微服务的们了~



那在来看看，我们是如何管理Dousheng应用的生命周期的。
