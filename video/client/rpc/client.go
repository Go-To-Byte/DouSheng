// Package rpc
// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package rpc

import (
	"github.com/Go-To-Byte/DouSheng/video/apps/video"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
)

// 用户中心 rpc 服务的 SDK

const (
	discoverName = "video"
)

type VideoClient struct {
	Client video.VideoClient

	l logger.Logger
}

// NewCommentClientFromCfg 从配置文件读取注册中心配置
func NewCommentClientFromCfg() (*VideoClient, error) {
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

// NewCommentClientFromEnv 从环境变量读取注册中心配置
func NewCommentClientFromEnv() (*VideoClient, error) {
	// 注册中心配置 [从环境变量文件中读取]

	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 去发现 服务
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *VideoClient {
	conn := clientSet.Conn()
	return &VideoClient{
		l: zap.L().Named("USER_CENTER_RPC"),

		// User 服务
		Client: video.NewVideoClient(conn),
	}
}

func (c *VideoClient) CommentService() video.VideoClient {
	if c.Client == nil {
		c.l.Errorf("获取用户中心[Token Client]失败")
		return nil
	}
	return c.Client
}
