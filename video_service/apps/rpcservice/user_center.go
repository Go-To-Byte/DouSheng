// @Author: Ciusyan 2023/2/8
package rpcservice

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc"

	"github.com/Go-To-Byte/DouSheng/user_center/conf"

	videoCfg "github.com/Go-To-Byte/DouSheng/video_service/conf"
)

// 用户中心的 rpc 服务

var (
	discoverName = "user_center"
)

type userCenter struct {
	tokenService token.ServiceClient
	userService  user.ServiceClient

	l logger.Logger
}

// NewUserCenterFromCfg 从配置文件读取注册中心配置
func NewUserCenterFromCfg() (*userCenter, error) {
	// 注册中心配置 [从配置文件中读取]
	consulCfg := videoCfg.C().Consul
	// 去发现 user_center 服务
	rpcCfg := rpc.NewConfig((*conf.Consul)(consulCfg), discoverName)

	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := rpc.NewClientSet(rpcCfg)

	if err != nil {
		return nil,
			exception.WithMsg("获取服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// NewUserCenterFromEnv 从环境变量读取注册中心配置
func NewUserCenterFromEnv() (*userCenter, error) {
	// 注册中心配置 [从环境变量文件中读取]
	consulCfg := videoCfg.NewDefaultConsul()
	// 去发现 user_center 服务
	rpcCfg := rpc.NewConfig((*conf.Consul)(consulCfg), "user_center")
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := rpc.NewClientSet(rpcCfg)

	if err != nil {
		return nil,
			exception.WithMsg("获取服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *rpc.ClientSet) *userCenter {
	return &userCenter{
		l: zap.L().Named("USER_CENTER_RPC"),

		tokenService: clientSet.Token(),
		// TODO：客户端还未暴露此模块
		// tokenService: clientSet.User(),
	}
}

func (c *userCenter) TokenService() token.ServiceClient {
	if c.tokenService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.tokenService
}

func (c *userCenter) UserService() user.ServiceClient {
	if c.userService == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.userService
}
