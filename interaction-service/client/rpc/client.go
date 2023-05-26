// Created by yczbest at 2023/02/19 19:02

package rpc

import (
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/message"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
)

// 视频互动服务
var (
	discoverName = conf.C().Consul.Register.RegistryName
)

// InteractionServiceClient 互动服务SDK
type InteractionServiceClient struct {
	l logger.Logger

	// 评论服务
	commentService comment.ServiceClient
	// 消息服务
	messageService message.ServiceClient
}

// 构建初始视频互动RPC服务客户端
func newDefault(clientSet *client.ClientSet) *InteractionServiceClient {
	conn := clientSet.Conn()
	return &InteractionServiceClient{
		l: zap.L().Named("Interaction_Service_RPC"),

		// 评论服务
		commentService: comment.NewServiceClient(conn),
		// Message 服务
		messageService: message.NewServiceClient(conn),
	}
}

// NewInteractionServiceClientFromConfig 从配置文件读取视频互动RPC服务配置，构建客户端
func NewInteractionServiceClientFromConfig() (*InteractionServiceClient, error) {
	//注册中心配置读取
	cfg := conf.C().Consul.Discovers[discoverName]

	clientSet, err := client.NewClientSet(cfg)
	if err != nil {
		return nil, exception.WithStatusMsgf("获取视频互动服务[%s]失败：%s", cfg.DiscoverName, err.Error())
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

// CommentService 视频评论RPC客户端
func (c *InteractionServiceClient) CommentService() comment.ServiceClient {
	if c.commentService == nil {
		c.l.Errorf("获取视频评论客户端[comment 服务]失败")
		return nil
	}

	return c.commentService
}

func (c *InteractionServiceClient) MessageService() message.ServiceClient {
	if c.messageService == nil {
		c.l.Errorf("获取消息客户端[message 服务]失败")
		return nil
	}

	return c.messageService
}
