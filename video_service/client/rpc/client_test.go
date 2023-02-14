// @Author: Ciusyan 2023/2/9
package rpc_test

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"testing"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/client/rpc"
)

var (
	videoService *rpc.VideoServiceClient
)

func TestPublishVideo(t *testing.T) {
	request := video.NewPublishVideoRequest()
	request.Title = "xxxx"
	request.Token = "xxxx"
	videoService.VideoService().PublishVideo(context.Background(), request)
}

func init() {
	// 需要先加载配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 获取视频服务的客户端[从环境变量中获取配置]
	// 获取的配置去执行 kit 库中的 client.NewConfig(consulCfg, discoverName)
	center, err := rpc.NewVideoServiceClientFromEnv()
	if err != nil {
		panic(err)
	}
	videoService = center
}
