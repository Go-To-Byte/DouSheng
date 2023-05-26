# rpc服务通用的 ClientSet

我们在启动的时候，已经将GRPC服务注册到Consul了，那我们该如何进行GRPC调用呢？

## 一、定义配置对象

先说结论：配置以下的 `Discover` 对象
```go
type Discover struct {
    DiscoverName string `toml:"discover_name" env:"DISCOVER_NAME"`
    Addr         string `toml:"address" env:"DISCOVER_ADDRESS"`
}
```

既然服务已经注册到Consul了，我们需要知道内部调用时，要先知道Consul在哪， 所以需要设置Consul的地址：

```go
// SetAddr 设置Consul的地址
func (c *DiscoverConfig) SetAddr(addr string) {
    c.Addr = addr
}

// SetDiscoverName 设置Consul的名称
func (c *DiscoverConfig) SetDiscoverName(name string) {
    c.DiscoverName = name
}
```

知道Consul在哪了， 还需要知道要去调用谁，所以需要设置服务发现的名称：

```go
// SetDiscoverName 设置Consul中服务发现的名称
func (c *DiscoverConfig) SetDiscoverName(name string) {
    c.DiscoverName = name
}
```

## 二、根据配置对象，去获取待调用服务的客户端连接对象

```go
conn, err := grpc.Dial(
		cfg.GrpcDailUrl(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

// GrpcDailUrl 获取待发现服务的 URL [用于grpc解析出对应的服务]
func (c *DiscoverConfig) GrpcDailUrl(discoverName string) string {
    return fmt.Sprintf("Register://%s/%s?wait=14s", c.Addr, discoverName)
}
```

## 三、使用方式

测试案例详见：[client_test.go]

```go
// rpc服务通用客户端
// 需要配置注册中心的[地址、服务名称]
func TestClient(t *testing.T) {
	should := assert.New(t)

	// 配置Consul[地址、服务名称]
	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("DISCOVER_ADDRESS"))
	cfg.SetDiscoverName("DISCOVER_NAME")

	// 比如这里去发现 user-service 服务
	// 根据注册中心的配置，获取用户中心的客户端
	client, err := client.NewClientSet(cfg)

	// 下面就可以使用user_center提供的SDK了
	if should.NoError(err) {
		t.Log(client)
	}
}

func init() {
	conf.LoadConfigFromEnv()
}

```
