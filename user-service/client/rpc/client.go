// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
)

// 用户中心 rpc 服务的 SDK

var (
	// 由自己服务的对外提供SDK
	discoverName = conf.C().Consul.Register.RegistryName
)

type UserCenterClient struct {
	userService user.ServiceClient

	l logger.Logger
}

// NewUserCenterClientFromCfg 从配置文件读取注册中心配置
func NewUserCenterClientFromCfg() (*UserCenterClient, error) {
	// 注册中心配置 [从配置文件中读取]
	cfg := conf.C().Consul.Discovers[discoverName]

	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// NewUserCenterClientFromEnv 从环境变量读取注册中心配置
func NewUserCenterClientFromEnv() (*UserCenterClient, error) {
	// 注册中心配置 [从环境变量文件中读取]

	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 去发现 user-service 服务
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *UserCenterClient {
	conn := clientSet.Conn()
	return &UserCenterClient{
		l: zap.L().Named("USER_CENTER_RPC"),

		// User 服务
		userService: user.NewServiceClient(conn),
	}
}

func (c *UserCenterClient) UserService() user.ServiceClient {
	if c.userService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.userService
}
