// Author: BeYoung
// Date: 2023/2/3 15:21
// Software: GoLand

package services

import (
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Favorite(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Favorite: %v", userID)
	c := proto.NewFavoriteClient(models.Dials["favorite"])

}

func FavoriteList(ctx *gin.Context) {

}
