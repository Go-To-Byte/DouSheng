// Author: BeYoung
// Date: 2023/1/29 14:16
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/middle"
	"github.com/Go-To-Byte/DouSheng/apps/user/models"
	"github.com/Go-To-Byte/DouSheng/apps/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter() {
	models.Router = gin.Default()

	// 健康检查
	models.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 添加处理跨域中间件
	models.Router.Use(middle.Cors())

	// 添加 api(http) 路由
	user := models.Router.Group("/douyin/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", service.Register)
	}
}
