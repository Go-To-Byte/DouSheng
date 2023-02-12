// @Author: Ciusyan 2023/2/5
package client

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// ClientSet  客户端
type ClientSet struct {
	conn *grpc.ClientConn
	l    logger.Logger
}

func NewClientSet(cfg *conf.Discover) (*ClientSet, error) {

	conn, err := grpc.Dial(
		cfg.GrpcDailUrl(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	// 初始化client全局日志对象
	zap.DevelopmentSetup()
	return &ClientSet{
		conn: conn,
		l:    zap.L(),
	}, nil
}

func (c *ClientSet) Conn() *grpc.ClientConn {
	return c.conn
}
