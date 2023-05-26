// @Author: Hexiaoming 2023/2/19
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/message-service/apps/message"
)

// 消息 rpc 的 SDK

var (
	discoverName = "message-service"
)

type MessageServiceClient struct {
	messageService message.ServiceClient

	l logger.Logger
}

// 从配置文件读取注册中心配置
func NewMessageServiceClientFromCfg() (*MessageServiceClient, error) {
	// 注册中心配置 [从配置文件中读取]
	cfg := conf.C().Consul.Discovers[discoverName]

	// 去发现 message-service 服务
	// 根据注册中心的配置，获取消息服务的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// 从环境变量读取注册中心配置
func NewMessageServiceClientFromEnv() (*MessageServiceClient, error) {
	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 根据注册中心的配置，获取消息服务的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *MessageServiceClient {
	conn := clientSet.Conn()
	return &MessageServiceClient{
		l: zap.L().Named("Message_Service_RPC"),

		// Message 服务
		messageService: message.NewServiceClient(conn),
	}
}

func (c *MessageServiceClient) MessageService() message.ServiceClient {
	if c.messageService == nil {
		c.l.Errorf("获取消息客户端[message 服务]失败")
		return nil
	}
	return c.messageService
}
