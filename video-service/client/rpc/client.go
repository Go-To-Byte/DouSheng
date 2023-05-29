// @Author: Ciusyan 2023/2/8
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou-kit/client"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/video-service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
)

// 视频服务 rpc 的 SDK

const (
	discoverName = "video-service"
)

// VideoServiceClient 视频服务SDK
type VideoServiceClient struct {
	// 视频服务
	videoService video.ServiceClient
	// 点赞服务
	favoriteService favorite.ServiceClient

	l logger.Logger
}

// NewVideoServiceClientFromCfg 从配置文件读取注册中心配置
func NewVideoServiceClientFromCfg() (*VideoServiceClient, error) {
	// 注册中心配置 [从配置文件中读取]
	cfg := conf.C().Consul.Discovers[discoverName]

	// 去发现 video-service 服务
	// 根据注册中心的配置，获取用户中心的客户端
	clientSet, err := client.NewClientSet(cfg)

	if err != nil {
		return nil,
			exception.WithStatusMsgf("获取服务[%s]失败：%s", cfg.DiscoverName, err.Error())
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
		// 视频服务
		favoriteService: favorite.NewServiceClient(conn),
	}
}

func (c *VideoServiceClient) VideoService() video.ServiceClient {
	if c.videoService == nil {
		c.l.Errorf("获取视频客户端[video 服务]失败")
		return nil
	}
	return c.videoService
}

// FavoriteService 视频点赞RPC客户端实例
func (c *VideoServiceClient) FavoriteService() favorite.ServiceClient {
	// 构建客户端失败，返回错误信息
	if c.favoriteService == nil {
		c.l.Errorf("获取视频点赞客户端[favorite 服务]失败")
		return nil
	}

	return c.favoriteService
}
