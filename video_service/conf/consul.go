// @Author: Ciusyan 2023/2/7
package conf

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
)

//=====
// consul配置对象
//=====

type Consul struct {
	Name string   `toml:"name" env:"CONSUL_NAME"`
	Host string   `toml:"host" env:"CONSUL_HOST"`
	Port int      `toml:"port" env:"CONSUL_PORT"`
	Tags []string `toml:"tags" env:"CONSUL_TAGS"`
}

func NewDefaultConsul() *Consul {
	return &Consul{
		Name: "user_center",
		Host: "127.0.0.1",
		Port: 8500,
		Tags: []string{"用户中心", "user_center_service"},
	}
}

// Addr 获取配置中心地址
func (c *Consul) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Consul) GrpcDailUrl() string {
	return fmt.Sprintf("Consul://%s:%d/%s?wait=14s", c.Host, c.Port, c.Name)
}

var globalConsulClient *api.Client

func ConsulClient() *api.Client {
	if globalConsulClient == nil {
		panic("加载全局consul配置失败")
	}
	return globalConsulClient
}
