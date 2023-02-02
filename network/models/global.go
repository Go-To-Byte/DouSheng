// Author: BeYoung
// Date: 2023/2/1 0:12
// Software: GoLand

package models

import (
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	V              *viper.Viper
	Router         *gin.Engine
	GrpcConn       *grpc.ClientConn
	Config         = ConfigYAML{}
	UserClient     proto.UserClient
	RelationClient proto.RelationClient
)
