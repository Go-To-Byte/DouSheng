// Author: BeYoung
// Date: 2023/1/25 23:49
// Software: GoLand

package service

func RegisterRouter() {
	r := GetRouter()
	user := r.Group("/douyin/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", register)
	}

	go func() {
		err := r.Run("127.0.0.1:8080")
		if err != nil {
			return
		}
	}()
}