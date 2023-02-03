// Author: BeYoung
// Date: 2023/2/3 20:52
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func Video() {
	v := models.Router.Group("/douyin")
	{
		v.GET("/feed", services.Feed) // 可不进行身份认证
		v.POST("/publish/action", milddlewares.JWTAuth(), services.Publish)
		v.GET("/publish/list", milddlewares.JWTAuth(), services.PublishList)
	}
}
