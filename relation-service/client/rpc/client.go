// @Author: Hexiaoming 2023/2/18
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/relation-service/apps/relation"
)

// 关系服务 rpc 的 SDK

var (
	discoverName = "relation-service"
)

type RelationServiceClient struct {
	relationService relation.ServiceClient

	l logger.Logger
}

// 从配置文件读取注册中心配置
func NewRelationServiceClientFromCfg() (*RelationServiceClient, error) {
	// 注册中心配置 [从配置文件中读取]
	cfg := conf.C().Consul.Discovers[discoverName]

	// 去发现 relation-service 服务
	// 根据注册中心的配置，获取关系服务的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

// 从环境变量读取注册中心配置
func NewRelationServiceClientFromEnv() (*RelationServiceClient, error) {
	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 根据注册中心的配置，获取关系服务的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", discoverName, err.Error())
	}
	return newDefault(clientSet), nil
}

func newDefault(clientSet *client.ClientSet) *RelationServiceClient {
	conn := clientSet.Conn()
	return &RelationServiceClient{
		l: zap.L().Named("Relation_Service_RPC"),

		// Relation 服务
		relationService: relation.NewServiceClient(conn),
	}
}

func (c *RelationServiceClient) RelationService() relation.ServiceClient {
	if c.relationService == nil {
		c.l.Errorf("获取关系服务客户端[relation 服务]失败")
		return nil
	}
	return c.relationService
}
