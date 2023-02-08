// @Author: Ciusyan 2023/2/7
package rpc

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Go-To-Byte/DouSheng/video_service/conf"
)

// Config 客户端配置对象
type Config struct {
	// Consul 的配置通过配置文件or环境变量获取
	conf.Consul
	// 服务发现的名称手动传入，因为只有使用方才知道需要去发现谁
	DiscoverName string
}

func NewClientSet(cfg *Config) (*ClientSet, error) {

	conn, err := grpc.Dial(
		cfg.GrpcDailUrl(cfg.DiscoverName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	if err != nil {
		return nil, err
	}

	// 初始化client全局日志对象
	zap.DevelopmentSetup()
	return &ClientSet{
		conn: conn,
		l:    zap.L(),
	}, nil
}

// ClientSet  客户端
type ClientSet struct {
	conn *grpc.ClientConn
	l    logger.Logger
}
