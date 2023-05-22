// @Author: Ciusyan 2023/2/5
package conf

import (
	"fmt"
	"github.com/hashicorp/consul/api"

	// register 用于服务发现解析出对应的服务
	_ "github.com/mbobakov/grpc-consul-resolver"
)

//=====
// consul配置对象
//=====

type consul struct {
	Register  *register  `mapstructure:"register" json:"register" yaml:"register"`
	Discovers []Discover `mapstructure:"discovers" json:"discovers" yaml:"discovers"`
}

func NewDefaultConsul() *consul {
	return &consul{
		Register:  NewDefaultRegister(),
		Discovers: NewDefaultDiscover(),
	}
}

// register Consul 用于服务注册
type register struct {
	RegistryName string   `mapstructure:"registry_name" json:"registry_name" yaml:"registry_name"`
	Host         string   `mapstructure:"host" json:"host" yaml:"host"`
	Port         int      `mapstructure:"port" json:"port" yaml:"port"`
	Tags         []string `mapstructure:"tags" json:"tags" yaml:"tags"`
}

func NewDefaultRegister() *register {
	return &register{
		RegistryName: "注册名称",
		Host:         "127.0.0.1",
		Port:         8500,
		Tags:         []string{"测试标签", "test_tag"},
	}
}

// Discover ：Discover Consul 用于服务发现
// 为什么服务发现和注册的配置对象要分开？
// 因为服务注册和服务发现的地址可能不一样：
// [比如：A部门的服务A是放在注册中心A的，B部门的服务B是放在注册中心B的，
// 然后A部门想要去内部调用B服务，它注册中心的地址总不能是自己的吧！]
// 当然，也可能是放在一起的。
type Discover struct {
	DiscoverName string `toml:"discover_name" env:"CONSUL_DISCOVER_NAME"`
	Addr         string `toml:"address" env:"CONSUL_ADDR"`
}

func NewDefaultDiscover() []Discover {
	return []Discover{}
}

func (d *Discover) SetDiscoverName(name string) {
	d.DiscoverName = name
}

func (d *Discover) SetAddr(addr string) {
	d.DiscoverName = addr
}

// GrpcDailUrl 获取待发现服务的 URL [用于grpc解析出对应的服务]
func (c *Discover) GrpcDailUrl() string {
	return fmt.Sprintf("Consul://%s/%s?wait=14s", c.Addr, c.DiscoverName)
}

// Addr 获取配置中心地址
func (c *register) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

var globalConsulClient *api.Client

// ConsulClient 获取Consul的Client
func ConsulClient() *api.Client {
	if globalConsulClient == nil {
		panic("加载全局consul配置失败")
	}
	return globalConsulClient
}
