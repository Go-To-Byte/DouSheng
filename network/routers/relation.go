// Author: BeYoung
// Date: 2023/2/2 0:17
// Software: GoLand

package routers

import (
	"github.com/Go-To-Byte/DouSheng/network/milddles"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/Go-To-Byte/DouSheng/network/services"
)

func Relation() {
	r := models.Router.Group("/douyin/relation")
	{
		r.Use(milddles.JWTAuth())
		r.POST("/action", services.Follow)
		r.GET("/follow/list", services.FollowList)
		r.GET("/follower/list", services.FollowerList)
		r.GET("/friend/list", services.FriendList)
	}
}
