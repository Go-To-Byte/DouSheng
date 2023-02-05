// @Author: Ciusyan 2023/1/29
package conf

import "fmt"

//=====
// App配置对象
//=====

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	HTTP *http  `toml:"http" env:"HTTP"`
	GRPC *grpc  `toml:"grpc" env:"GRPC"`
}

// HTTP 服务配置
type http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port string `toml:"port" env:"HTTP_PORT"`
}

// GRPC 服务配置
type grpc struct {
	Host string `toml:"host" env:"GRPC_HOST"`
	Port int    `toml:"port" env:"GRPC_PORT"`
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
		Port: "8050",
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
