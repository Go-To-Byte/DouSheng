// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

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

// NewVideoServiceClientFromCfg 从配置读取注册中心配置
func NewVideoServiceClientFromCfg() (*VideoServiceClient, error) {
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

// NewVideoServiceClientFromEnv 从环境变量读取注册中心配置
func NewVideoServiceClientFromEnv() (*VideoServiceClient, error) {
	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("CONSUL_ADDR"))
	cfg.SetDiscoverName(os.Getenv("CONSUL_DISCOVER_NAME"))

	// 根据注册中心的配置，获取视频服务的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", discoverName, err.Error())
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

func (c *VideoServiceClient) VideoService() video.ServiceClient {
	if c.videoService == nil {
		c.l.Errorf("获取视频客户端[video 服务]失败")
		return nil
	}
	return c.videoService
}
