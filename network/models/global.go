// Author: BeYoung
// Date: 2023/2/1 0:12
// Software: GoLand

package models

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	V        *viper.Viper
	Router   *gin.Engine
	GrpcConn *grpc.ClientConn
	Config   = ConfigYAML{}
)
