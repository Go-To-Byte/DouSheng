// Author: BeYoung
// Date: 2023/1/29 14:16
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/models"
	"github.com/Go-To-Byte/DouSheng/apps/user/service"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	models.Router = gin.Default()
	user := models.Router.Group("/douyin/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", service.Register)
	}
}
