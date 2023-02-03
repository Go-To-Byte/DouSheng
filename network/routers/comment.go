// Author: BeYoung
// Date: 2023/2/3 13:12
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func Comment() {
	r := models.Router.Group("/douyin/Comment")
	{
		r.POST("/action", services.Comment)
		r.GET("/list", services.CommentList)
	}
}
