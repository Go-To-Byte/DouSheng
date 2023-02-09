// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

// 用户中心 rpc 服务的 SDK

var (
	discoverName = "user_center"
)

type UserCenterClient struct {
	tokenService token.ServiceClient
	userService  user.ServiceClient

	l logger.Logger
}

// NewUserCenterClientFromCfg 从配置文件读取注册中心配置
func NewUserCenterClientFromCfg() (*UserCenterClient, error) {
	// 注册中心配置 [从配置文件中读取]
	consulCfg := conf.C().Consul
	// 去发现 user_center 服务
	rpcCfg := client.NewConfig(consulCfg, discoverName)

	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(rpcCfg)

	if err != nil {
		return nil,
			exception.WithMsg("获取服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// NewUserCenterClientFromEnv 从环境变量读取注册中心配置
func NewUserCenterClientFromEnv() (*UserCenterClient, error) {
	// 注册中心配置 [从环境变量文件中读取]
	consulCfg := conf.NewDefaultConsul()
	// 去发现 user_center 服务
	rpcCfg := client.NewConfig(consulCfg, discoverName)
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(rpcCfg)

	if err != nil {
		return nil,
			exception.WithMsg("获取服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *UserCenterClient {
	conn := clientSet.Conn()
	return &UserCenterClient{
		l: zap.L().Named("USER_CENTER_RPC"),

		// Token 服务
		tokenService: token.NewServiceClient(conn),
		// User 服务
		userService: user.NewServiceClient(conn),
	}
}

func (c *UserCenterClient) TokenService() token.ServiceClient {
	if c.tokenService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.tokenService
}

func (c *UserCenterClient) UserService() user.ServiceClient {
	if c.userService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.userService
}
