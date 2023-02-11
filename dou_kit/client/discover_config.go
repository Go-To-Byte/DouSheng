// @Author: Ciusyan 2023/2/12
package client

import "fmt"

// DiscoverConfig ：DiscoverConfig Consul 用于服务发现
// 为什么服务发现和注册的配置对象要分开？
// 因为服务注册和服务发现的地址可能不一样：
// [比如：A部门的服务A是放在注册中心A的，B部门的服务B是放在注册中心B的，
// 然后A部门想要去内部调用B服务，它注册中心的地址总不能是自己的吧！]
// 当然，也可能是放在一起的。
// 既然是配置，为什么不放在配置文件中？
// 因为服务A可能需要内部调用服务B、C、D，放在配置对象中不方便配置多个服务发现的地址
type DiscoverConfig struct {
	DiscoverName string `json:"discover_name" env:"DISCOVER_NAME"`
	Addr         string `json:"address" env:"DISCOVER_ADDRESS"`
}

func NewDefaultDiscoverCfg() *DiscoverConfig {
	return &DiscoverConfig{
		DiscoverName: "发现名称",
		Addr:         "127.0.0.1:8500",
	}
}

func NewDiscoverCfg(name, addr string) *DiscoverConfig {
	return &DiscoverConfig{
		DiscoverName: name,
		Addr:         addr,
	}
}

// SetAddr 设置Consul的地址
func (c *DiscoverConfig) SetAddr(addr string) {
	c.Addr = addr
}

// SetDiscoverName 设置Consul中服务发现的名称
func (c *DiscoverConfig) SetDiscoverName(name string) {
	c.DiscoverName = name
}

// GrpcDailUrl 获取待发现服务的 URL [用于grpc解析出对应的服务]
func (c *DiscoverConfig) GrpcDailUrl(discoverName string) string {
	return fmt.Sprintf("Register://%s/%s?wait=14s", c.Addr, discoverName)
}
