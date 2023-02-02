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
	targetUser := fmt.Sprintf("consul://%v:%v/%v?wait=14s", models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.User),
	if dial, err := grpc.Dial(targetUser, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		zap.S().Panicf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", target)
		models.GrpcConn = dial
	}

}
