// Author: BeYoung
// Date: 2023/2/3 15:19
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func Favorite() {
	f := models.Router.Group("/douyin/favorite")
	{
		f.POST("/action", services.Favorite)
		f.GET("/list", services.FavoriteList)
	}
	f.Use(milddlewares.JWTAuth())
}