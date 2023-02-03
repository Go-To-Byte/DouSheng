// Author: BeYoung
// Date: 2023/2/3 18:10
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func Message() {
	m := models.Router.Group("/douyin/message")
	{
		m.Use(milddlewares.JWTAuth())
		m.POST("/action", services.Message)
		m.GET("/chat", services.MessageList)
	}
}
