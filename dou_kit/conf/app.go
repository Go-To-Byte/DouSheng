// @Author: Ciusyan 2023/2/7
package conf

import "fmt"

//=====
// App配置对象
//=====

type app struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	HTTP *http  `mapstructure:"http" json:"http" yaml:"http"`
	GRPC *grpc  `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
}

// HTTP 服务配置
type http struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

// GRPC 服务配置
type grpc struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func NewDefaultApp() *app {
	return &app{
		Name: "dousheng",
		HTTP: newDefaultHTTP(),
		GRPC: newDefaultGRPC(),
	}
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8080",
	}
}

func newDefaultGRPC() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: 8505,
	}
}

// Addr 获取 HTTP 服务配置的 IP + 端口
func (h *http) Addr() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}

// Addr 获取 GRPC 服务配置的 IP + 端口
func (g *grpc) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
