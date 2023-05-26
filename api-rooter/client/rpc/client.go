// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
)

// Api路由 rpc 服务的 SDK

const (
	discoverName = "api-rooter"
)

type ApiRooterClient struct {
	tokenService token.ServiceClient

	l logger.Logger
}

// NewApiRooterClientFromCfg 从配置文件读取注册中心配置
func NewApiRooterClientFromCfg() (*ApiRooterClient, error) {
	// 注册中心配置 [从配置文件中读取]
	cfg := conf.C().Consul.Discovers[discoverName]

	// 根据注册中心的配置，获取Api路由的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// NewApiRooterClientFromEnv 从环境变量读取注册中心配置
func NewApiRooterClientFromEnv() (*ApiRooterClient, error) {
	// 注册中心配置 [从环境变量文件中读取]

	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 去发现 api-rooter 服务
	// 根据注册中心的配置，获取Api路由的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *ApiRooterClient {
	conn := clientSet.Conn()
	return &ApiRooterClient{
		l: zap.L().Named("USER_CENTER_RPC"),

		// Token 服务
		tokenService: token.NewServiceClient(conn),
	}
}

func (c *ApiRooterClient) TokenService() token.ServiceClient {
	if c.tokenService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.tokenService
}
