// Author: BeYoung
// Date: 2023/2/1 0:12
// Software: GoLand

package models

import (
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	V              *viper.Viper
	Router         *gin.Engine
	Config         = ConfigYAML{}
	UserClient     proto.UserClient
	VideoClient    proto.VideoClient
	CommentClient  proto.CommentClient
	MessageClient  proto.ChatClient
	RelationClient proto.RelationClient
	FavoriteClient proto.FavoriteClient
)
