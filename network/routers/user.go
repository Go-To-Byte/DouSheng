// Author: BeYoung
// Date: 2023/2/1 0:26
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/milddles"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func User() {
	u := models.Router.Group("/douyin/user")
	{
		u.GET("/", milddles.JWTAuth(), services.Info)
		u.POST("/login/", services.Login)
		u.POST("/register/", services.Register)
		u.GET("/login/", services.Login)
		u.GET("/register/", services.Register)
	}
}
