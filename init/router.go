// Author: BeYoung
// Date: 2023/1/27 5:43
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/service"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	router = gin.Default()
	initUser(router)
}

func initUser(r *gin.Engine) {
	user := r.Group("/douyin/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", service.Register)
	}
}
