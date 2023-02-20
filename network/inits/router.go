// Author: BeYoung
// Date: 2023/2/1 0:13
// Software: GoLand

package inits

import (
	"github.com/Go-To-Byte/DouSheng/network/milddles"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter() {
	models.Router = gin.Default()
	models.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	models.Router.Use(milddles.Cors())
	routers.Init()
}
