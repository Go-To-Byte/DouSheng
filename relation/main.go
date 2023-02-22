// Author: BeYoung
// Date: 2023/1/25 23:50
// Software: GoLand

package main

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/cmd"
	_ "github.com/Go-To-Byte/DouSheng/relation/apps/relation/impl/init"

	// 驱动加载所有变量，主要是[IOC的实例]
	_ "github.com/Go-To-Byte/DouSheng/relation/common/all"
)

func main() {
	// 交由CLI启动
	cmd.Main()
}

//
// import (
// 	"fmt"
// 	"github.com/Go-To-Byte/DouSheng/apps/relation/apps/relation/impl/models"
// 	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
// 	"github.com/Go-To-Byte/DouSheng/apps/relation/service"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/health"
// 	"net"
// 	"time"
//
// 	_ "github.com/Go-To-Byte/DouSheng/apps/relation/init"
// 	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
// 	healthpb "google.golang.org/grpc/health/grpc_health_v1"
// )
//
// func main() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", models.Config.Localhost.Port))
// 	if err != nil {
// 		zap.S().Infof("failed to listen: %v", err)
// 	}
//
// 	s := grpc.NewServer()             // 注册 grpc server
// 	healthcheck := health.NewServer() // 注册健康检查
// 	healthgrpc.RegisterHealthServer(s, healthcheck)
// 	proto.RegisterRelationServer(s, &service.Relation{})
//
// 	// 异步检查依赖关系并根据需要切换服务状态
// 	go func() {
// 		next := healthpb.HealthCheckResponse_SERVING
//
// 		for {
// 			healthcheck.SetServingStatus("user", next)
//
// 			if next == healthpb.HealthCheckResponse_SERVING {
// 				next = healthpb.HealthCheckResponse_NOT_SERVING
// 			} else {
// 				next = healthpb.HealthCheckResponse_SERVING
// 			}
//
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
//
// 	zap.S().Infof("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		zap.S().Infof("failed to serve: %v", err)
// 	}
// }
