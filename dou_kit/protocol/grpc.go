// @Author: Ciusyan 2023/2/10
package protocol

import (
	"github.com/hashicorp/consul/api"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"

	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
)

// =====
// 使用GRPC对外暴露服务
// =====

func NewGRPCService() *GRPCService {
	return &GRPCService{
		server: grpc.NewServer(),
		l:      zap.L().Named("GRPC Service"),
		cfg:    conf.C(),

		client: conf.ConsulClient(),
	}
}

type GRPCService struct {
	server *grpc.Server
	l      logger.Logger
	cfg    *conf.Config

	// consul的客户端
	client       *api.Client
	registration *api.AgentServiceRegistration
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// =====
	// 1、注册IOC中所有的GRPC服务
	// =====
	ioc.RegistryGrpc(s.server)

	// =====
	// 2、启动GRPC服务
	// =====
	addr := s.cfg.App.GRPC.Addr()
	listener, err := net.Listen("tcp", addr)
	s.l.Infof("[GRPC] 服务监听地址：%s", addr)
	if err != nil {
		s.l.Errorf("启动GRPC服务错误：%s", err.Error())
		return
	}

	// 理论上我们需要等待GRPC服务启动成功后，才注册此服务到注册中心
	// 但是 GRPC Server 并没有什么成功的回调通知

	// 所以我们只能假设GRPC Server 1秒后启动成功
	time.AfterFunc(1*time.Second, s.Register)

	// 注册服务开启健康检查
	grpc_health_v1.RegisterHealthServer(s.server, health.NewServer())

	// =====
	// 3、监听GRPC服务
	// =====
	if err = s.server.Serve(listener); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Infof("[GRPC] 服务关闭成功")
		}
		s.l.Errorf("开启 [GRPC] 服务异常：%s", err.Error())
		return
	}

}

// Stop 优雅关闭GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Infof("优雅关闭 [GRPC] 服务")
	// 注销在consul中的 GRPC 服务
	s.DeRegister()
	s.server.GracefulStop()
	return nil
}

// Register 注册GRPC服务到consul
func (s *GRPCService) Register() {

	// 生成对应grpc的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           s.cfg.App.GRPC.Addr(),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	consul := s.cfg.Consul

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = consul.RegistryName
	registration.ID = xid.New().String()
	registration.Port = s.cfg.App.GRPC.Port
	registration.Tags = consul.Tags
	registration.Address = s.cfg.App.GRPC.Host
	registration.Check = check

	err := s.client.Agent().ServiceRegister(registration)
	if err != nil {
		s.l.Errorf("注册GRPC服务到consul失败：%s", err)
		return
	}

	s.registration = registration
	s.l.Info("成功注册GRPC服务到consul")
}

// DeRegister 注销在Consul的GRPC服务
func (s *GRPCService) DeRegister() {
	if s.registration != nil {
		err := s.client.Agent().ServiceDeregister(s.registration.ID)
		if err != nil {
			s.l.Errorf("注销实例失败：%s", err)
		} else {
			s.l.Info("注销实例成功")
		}
	}
}
