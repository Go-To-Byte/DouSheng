// Package rpc
// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package rpc

import (
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
)

// 用户中心 rpc 服务的 SDK

const (
	discoverName = "favorite"
)

type FavoriteClient struct {
	Client favorite.FavoriteClient

	l logger.Logger
}

// NewFavoriteClientFromCfg 从配置文件读取注册中心配置
func NewFavoriteClientFromCfg() (*FavoriteClient, error) {
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

// NewFavoriteClientFromEnv 从环境变量读取注册中心配置
func NewFavoriteClientFromEnv() (*FavoriteClient, error) {
	// 注册中心配置 [从环境变量文件中读取]

	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 去发现 user_center 服务
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *FavoriteClient {
	conn := clientSet.Conn()
	return &FavoriteClient{
		l: zap.L().Named("USER_CENTER_RPC"),

		// User 服务
		Client: favorite.NewFavoriteClient(conn),
	}
}

func (c *FavoriteClient) UserService() favorite.FavoriteClient {
	if c.Client == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.Client
}
