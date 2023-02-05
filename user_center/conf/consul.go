// @Author: Ciusyan 2023/2/5
package conf

import (
	"fmt"
	"github.com/hashicorp/consul/api"

	_ "github.com/mbobakov/grpc-consul-resolver"
)

//=====
// consul配置对象
//=====

type consul struct {
	Name string   `toml:"name" env:"NAME"`
	Host string   `toml:"address" env:"CONSUL_ADDRESS"`
	Port int      `toml:"port" env:"CONSUL_PORT"`
	Tags []string `toml:"tags" env:"TAGS"`
}

func NewDefaultConsul() *consul {
	return &consul{
		Name: "user_center",
		Host: "127.0.0.1",
		Port: 8500,
		Tags: []string{"用户中心", "user_center_service"},
	}
}

// Addr 获取配置中心地址
func (c *consul) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *consul) GrpcDailUrl() string {
	return fmt.Sprintf("consul://%s:%d/%s?wait=14s", c.Host, c.Port, c.Name)
}

var globalConsulClient *api.Client

func ConsulClient() *api.Client {
	if globalConsulClient == nil {
		panic("加载全局consul配置失败")
	}
	return globalConsulClient
}
