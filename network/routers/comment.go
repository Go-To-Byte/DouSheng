// Author: BeYoung
// Date: 2023/2/3 13:12
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func Comment() {
	c := models.Router.Group("/douyin/Comment")
	{
		c.POST("/action", services.Comment)
		c.GET("/list", services.CommentList)
	}
	c.Use(milddlewares.JWTAuth())
}
