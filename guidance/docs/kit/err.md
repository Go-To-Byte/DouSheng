---
title: 统一error处理
author: Ciusyan
date: '2023-2-25'
---



# GRPC官方的error处理，我放在HTTP中1样的好用



`Go`的error机制，可能被吐槽得算多的了。这里就不吐槽了。既然要来做统一的error处理。



在这之前，你总的告诉我，为什么需要做异常处理吧！不可能说因为某某说：~~*一个后端项目中，不能缺少异常处理哟，没有的话，可...*~~



这样可不行，咱们还是来看看，为什么需要统一error处理。



## 一、为什么需要统一error处理？



既然我们主要是处理`HTTP Handler的error`，那先来看看，我认为的，优雅的Handler层的代码长什么样。



### （1）我认为的优雅的Handler层

先说结论，我认为，Handler层核心应该只做三件事：

1. 处理什么uri
2. 接收请求参数&提取额外参数，去调用业务方法，执行具体的业务逻辑。
3. 拿到上一步调用返回的结果，进行响应。



为什么呢？我的理由如下，听我细细道来。



与其说是处理HTTP的Handler，还不如说是处理HTTP的报文`（HTTP Message）`。简单来了解一下[HTTP的报文格式](https://juejin.cn/post/7117192792446599199)：



![image-20230224151222317](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302241512407.png)



看了图，有点懵没关系，我就想说：HTTP 报文格式，核心就是这三部分`（start-line、header、body）`



既然报文长这样，那么HTTP的Handler，核心至少也得处理这三部分吧，用一幅图片简单对应一下：

![image-20230224152826933](https://ciusyan-picgo.oss-cn-shenzhen.aliyuncs.com/images/202302241528969.png)



好了，我的理由叙述完了，再来看看我一开始叙述的结论：

1. 处理什么uri：



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



2. 接收请求参数&提取额外参数，去调用业务方法，执行具体的业务逻辑。



比如拿登录这个简单的业务来举例，用户要登陆，至少要提供用户名和密码这两个参数吧。那这两个参数放在哪里呢？：是请求头？请求体？还是query string中呢？

反正不管是哪里，你总的把他从HTTP的报文中解析出来，这不就是在处理Header和Body吗？



拿到了这些参数过后，就可以去调用你的业务方法了。



ps：当然啊，我也见过有一些代码，在Handler里进行了大量的业务逻辑的处理，甚至还有的直接在这里去操纵数据库。不是说不行啊，但的确不太推荐~



3. 拿到上一步调用返回的结果，进行响应



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



谁想有事没事就碰到error啊，我们想要的流程明明是：接收用户请求 -> 处理业务逻辑 -> 返回成功响应



可偏偏半路杀出来一个`error`很嘲讽的说：你在调用我的时候出错了，你自己想想怎么处理我吧。要不然我可不让你好过！！！



于是就出现了上面的代码。那我们有没有办法，让他稍微优雅一些呢？



答案是有的，这就引出了我们今天的究极主题：**统一`error`处理**



## 二、我们是这样做的



//  TODO：要先写其他文档。暂时没有时间编写，可以看看仓库代码

* [统一error处理相关代码](https://github.com/Go-To-Byte/DouSheng/tree/main/dou_kit/exception)



### （1）HTTP出现异常



### （2）GRPC出现异常



