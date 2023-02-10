// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

// 视频服务 rpc 的 SDK

var (
	discoverName = "video_service"
)

type VideoServiceClient struct {
	videoService video.ServiceClient

	l logger.Logger
}

// NewVideoServiceClientFromCfg 从配置文件读取注册中心配置
func NewVideoServiceClientFromCfg() (*VideoServiceClient, error) {
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

// NewVideoServiceClientFromEnv 从环境变量读取注册中心配置
func NewVideoServiceClientFromEnv() (*VideoServiceClient, error) {
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

func newDefault(clientSet *client.ClientSet) *VideoServiceClient {
	conn := clientSet.Conn()
	return &VideoServiceClient{
		l: zap.L().Named("Video_Service_RPC"),

		// Video 服务
		videoService: video.NewServiceClient(conn),
	}
}

func (c *VideoServiceClient) TokenService() video.ServiceClient {
	if c.videoService == nil {
		c.l.Errorf("获取视频客户端[Token 服务]失败")
		return nil
	}
	return c.videoService
}

