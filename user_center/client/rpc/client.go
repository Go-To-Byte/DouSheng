// @Author: Ciusyan 2023/2/5
package rpc

import (
	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientSet(consulConf *conf.Consul) (*ClientSet, error) {

	conn, err := grpc.Dial(
		consulConf.GrpcDailUrl(),
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

func (c *ClientSet) Token() token.ServiceClient {
	return token.NewServiceClient(c.conn)
}
