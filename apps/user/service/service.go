// Author: BeYoung
// Date: 2023/1/26 2:47
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/dal/model"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getID() int64 {
	node := initSnowflakeNode()
	id := node.Generate()
	return id.Int64()
}

func getSqlDB() *gorm.DB {
	return initSqlServer()
}

func getSnowflake() *snowflake.Node {
	return initSnowflakeNode()
}

func register(context *gin.Context) {
	user := model.UserInfo{
		ID:       getID(),
		Username: context.Query("username"),
		Passwd:   context.Query("password"),
	}
	Add(user)
}
