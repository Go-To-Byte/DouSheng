// Author: BeYoung
// Date: 2023/2/1 1:18
// Software: GoLand

package inits

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initGrpc() {
	// target := fmt.Sprintf("%v:%v", models.Config.GrpcConfig.Host, models.Config.GrpcConfig.Port)
	target:=fmt.Sprintf("consul://%s:%d/%s?wait=14s", "192.168.75.141", 8500, global.ServerConfig.GoodsSrvInfo.Name),
	if dial, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		zap.S().Panicf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", target)
		models.GrpcConn = dial
	}

}
