// @Author: Ciusyan 2023/2/5
package conf

import (
	"fmt"
	"github.com/hashicorp/consul/api"

	// consul 用于服务发现解析出对应的服务
	_ "github.com/mbobakov/grpc-consul-resolver"
)

//=====
// consul配置对象
//=====

type consul struct {
	RegistryName string   `toml:"registry_name" env:"CONSUL_REGISTRY_NAME"`
	Host         string   `toml:"host" env:"CONSUL_HOST"`
	Port         int      `toml:"port" env:"CONSUL_PORT"`
	Tags         []string `toml:"tags" env:"CONSUL_TAGS"`
}

func NewDefaultConsul() *consul {
	return &consul{
		RegistryName: "注册名称",
		Host:         "127.0.0.1",
		Port:         8500,
		Tags:         []string{"测试标签", "test_tag"},
	}
}

// Addr 获取配置中心地址
func (c *consul) Addr() string {
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
