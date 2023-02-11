# rpc服务通用的 ClientSet

我们在启动的时候，已经将GRPC服务注册到Consul了，那我们该如何进行GRPC调用呢？

## 一、定义配置对象

先说结论：配置以下的 `DiscoverConfig` 对象
```go
type DiscoverConfig struct {
    DiscoverName string `json:"discover_name"`
    Host         string `json:"host"`
    Port         int    `json:"port"`
}
```

既然服务已经注册到Consul了，我们需要知道内部调用时，要先知道Consul在哪， 所以需要设置Consul的地址：

```go
// SetAddr 设置Consul的地址
func (c *DiscoverConfig) SetAddr(host string, port int) {
    c.Host = host
    c.Port = port
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
		cfg.GrpcDailUrl(cfg.DiscoverName),
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

const (
	Addr         = "127.0.0.1:8500"
	DiscoverName = "user_center"
)

// rpc服务通用客户端
// 需要配置注册中心的[地址、服务名称]
func TestClient(t *testing.T) {
	should := assert.New(t)

	// 配置Consul[地址、服务名称]
	cfg := client.NewDefaultDiscoverCfg()
	cfg.SetAddr(Addr)
	cfg.SetDiscoverName(DiscoverName)

	// 比如这里去发现 user_center 服务
	// 根据注册中心的配置，获取用户中心的客户端
	client, err := client.NewClientSet(cfg)

	// 下面就可以使用user_center提供的SDK了
	if should.NoError(err) {
		t.Log(client)
	}
}

```
