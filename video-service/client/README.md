# 视频服务[video_service]SDK

使用此SDK的方式如下测试代码：

```go
// @Author: Ciusyan 2023/2/9
var (
	videoService *rpc.VideoServiceClient
)

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

```

详细使用方式请看 client_test.go 文件