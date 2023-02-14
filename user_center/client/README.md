# 如何提供用户中心的[user_center]SDK —— GRPC调用方式


## 一、定义配置文件、通过dou_kit公共库，获取对应服务客户端

需要先定义好服务发现的配置，如：

```toml
[consul.discovers.user_center]
discover_name = "user_center"
address = "127.0.0.1:8500"
```

p.s.配置对象：

```go
type Discover struct {
	DiscoverName string `toml:"discover_name" env:"CONSUL_DISCOVER_NAME"`
	Addr         string `toml:"address" env:"CONSUL_ADDR"`
}
```

然后可通过两种方式获取用户中心的客户端
1. NewUserCenterClientFromCfg() —— 配置文件
2. NewUserCenterClientFromEnv() —— 环境变量

```go
    // 从配置文件中获取 user_center 的Client
	client, err := rpc.NewUserCenterClientFromCfg()

	if err != nil {
		return nil, err
	}
```

## 二、利用用户中心的客户端调用它的SDK

若提供配置无误，经过第一步后，我们已经获取了用户中心的客户端，现在就可以利用客户端来调用对外暴露的GRPC服务了，
可简单理解成调用 user_center 的SDK

使用此SDK的方式如下测试代码：

```go
var (
	userCenter *rpc.UserCenterClient
)

func TestUserCenter(t *testing.T) {
	should := assert.New(t)

	tokenReq := token.NewValidateTokenRequest("xxx")
	// 这里主要是为了获取 用户ID
	validatedToken, err := userCenter.TokenService().ValidateToken(context.Background(), tokenReq)
	if should.NoError(err) {
		t.Log(validatedToken)
	}

}
```

详细使用方式请看 client_test.go 文件

## 三、如何编写用户中心的客户端？

如果仅仅是调用方，看看上面即可。

那作为编写用户中心的coder，该如何对外提供SDK呢？

这是用户中心中，能对外暴露的所有接口（利用protobuf文件生成）：
1. `user.ServiceClient`

那咱们可以用面向对象的方式来编写用户中心的客户端，新建结构体`UserCenterClient`

```go
type UserCenterClient struct {
	userService  user.ServiceClient

	l logger.Logger
}
```

利用第一步，从dou_kit公共库中获取的ClientSet，连接对象，给接口附上实现：

```go
func newDefault(clientSet *client.ClientSet) *UserCenterClient {
	conn := clientSet.Conn()
	return &UserCenterClient{
		l: zap.L().Named("USER_CENTER_RPC"),

		// Token 服务
		tokenService: token.NewServiceClient(conn),
		// User 服务
		userService: user.NewServiceClient(conn),
	}
}
```

对外暴露Get方法，即可直接调用接口的具体实现：

```go

func (c *UserCenterClient) UserService() user.ServiceClient {
	if c.userService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.userService
}
```

当然，若直接对外暴露所有接口的实现，可能不是很好，可以仅暴露部分方法
