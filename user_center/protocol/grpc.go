// @Author: Ciusyan 2023/1/25
package protocol

import (
	"Go-To-Byte/DouSheng/user_center/conf"
	"Go-To-Byte/DouSheng/user_center/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"net"
)

// =====
// 使用GRPC对外暴露服务
// =====

func NewGRPCService() *GRPCService {
	return &GRPCService{
		server: grpc.NewServer(),
		l:      zap.L().Named("GRPC Service"),
		c:      conf.C(),
	}
}

type GRPCService struct {
	server *grpc.Server
	l      logger.Logger
	c      *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// =====
	// 1、注册IOC中所有的GRPC服务
	// =====
	ioc.RegistryGrpc(s.server)

	// =====
	// 2、启动HTTP服务
	// =====
	addr := s.c.App.GRPC.Addr()
	listener, err := net.Listen("tcp", addr)
	s.l.Infof("[GRPC] 服务监听地址：%s", addr)
	if err != nil {
		s.l.Errorf("启动HTTP服务错误：%s", err.Error())
		return
	}

	// =====
	// 3、监听GRPC服务
	// =====
	if err = s.server.Serve(listener); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Infof("[GRPC] 服务关闭成功")
		}
		s.l.Errorf("开启 [HTTP] 服务异常：%s", err.Error())
		return
	}

}

// Stop 优雅关闭GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Infof("优雅关闭 GRPC 服务")
	s.server.GracefulStop()
	return nil
}
