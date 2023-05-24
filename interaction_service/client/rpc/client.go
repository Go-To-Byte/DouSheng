// Created by yczbest at 2023/02/19 19:02

package rpc

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"
)

// 视频互动服务
var discoverName = "interaction_service"

type InteractionServiceClient struct {
	favoriteService favorite.ServiceClient
	l               logger.Logger
	commentService  comment.ServiceClient
}

// 构建初始视频互动RPC服务客户端
func newDefault(clientSet *client.ClientSet) *InteractionServiceClient {
	conn := clientSet.Conn()
	return &InteractionServiceClient{
		l:               zap.L().Named("Interaction_Service_RPC"),
		favoriteService: favorite.NewServiceClient(conn),
		commentService:  comment.NewServiceClient(conn),
	}
}

// NewInteractionServiceClientFromConfig 从配置读取视频互动RPC服务配置，构建客户端
func NewInteractionServiceClientFromConfig() (*InteractionServiceClient, error) {
	// 注册中心配置
	cfg := conf.C().Consul.Discovers[0]

	// 根据注册中心的配置，获取Api路由的客户端
	clientSet, err := client.NewClientSet(&cfg)
	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg, err.Error())
	}
	return newDefault(clientSet), nil
}

// NewInteractionServiceClientFromEnv 从环境变量读取视频互动服务RPC服务，构建客户端
func NewInteractionServiceClientFromEnv() (*InteractionServiceClient, error) {
	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 根据注册中心的配置，获取视频互动服务的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取视频互动服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// FavoriteService 视频点赞RPC客户端实例
func (c *InteractionServiceClient) FavoriteService() favorite.ServiceClient {
	//构建客户端失败，返回错误信息
	if c.favoriteService == nil {
		c.l.Errorf("获取视频点赞客户端[favorite 服务]失败")
		return nil
	}
	return c.favoriteService
}

// CommentService 视频评论RPC客户端
func (c *InteractionServiceClient) CommentService() comment.ServiceClient {
	if c.commentService == nil {
		c.l.Errorf("获取视频评论客户端[comment 服务]失败")
		return nil
	}
	return c.commentService
}
