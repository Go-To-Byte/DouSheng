// Author: BeYoung
// Date: 2023/1/26 2:47
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/dal/model"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func createID() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		zap.S().Panicf("createID failed: %v", err)
	}
	id := node.Generate()
	return id.Int64()
}

func register(context *gin.Context) {
	user := model.UserInfo{
		ID:       createID(),
		Username: context.Query("username"),
		Passwd:   context.Query("password"),
	}

}
