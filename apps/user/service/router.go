// Author: BeYoung
// Date: 2023/1/25 23:49
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/run"
)

func registerRouter() {
	r := run.GetRouter()
	user := r.Group("/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", register)
	}
}
