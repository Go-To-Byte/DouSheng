// Author: BeYoung
// Date: 2023/1/25 23:50
// Software: GoLand

package main

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/proto"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"net"
	"time"

	_ "github.com/Go-To-Byte/DouSheng/apps/message/init"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		zap.S().Infof("failed to listen: %v", err)
	}

	s := grpc.NewServer()             // 注册 grpc server
	healthcheck := health.NewServer() // 注册健康检查
	healthgrpc.RegisterHealthServer(s, healthcheck)
	proto.RegisterChatServer(s, &service.Message{})

	// 异步检查依赖关系并根据需要切换服务状态
	go func() {
		next := healthpb.HealthCheckResponse_SERVING

		for {
			healthcheck.SetServingStatus("user", next)

			if next == healthpb.HealthCheckResponse_SERVING {
				next = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				next = healthpb.HealthCheckResponse_SERVING
			}

			time.Sleep(1 * time.Second)
		}
	}()

	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Infof("failed to serve: %v", err)
	}
}
