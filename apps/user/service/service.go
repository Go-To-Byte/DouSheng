// Author: BeYoung
// Date: 2023/1/26 2:47
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/dal/model"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

func getID() int64 {
	node := getSnowflake()
	id := node.Generate()
	return id.Int64()
}

func getSqlDB() *gorm.DB {
	return initSqlServer()
}

func getSnowflake() *snowflake.Node {
	return initSnowflakeNode()
}

func GetRouter() *gin.Engine {
	return initRouter()
}

func register(context *gin.Context) {
	zap.S().Debugf("register")
	user := model.UserInfo{
		ID:       getID(),
		Username: context.Query("username"),
		Passwd:   context.Query("password"),
	}

	passwd := make([]byte, 128)
	copy(passwd, user.Passwd)

	result := FindByName(user)
	if result == nil || len(result) > 0 {
		context.String(200, "register failed")
		return
	}

	zap.S().Debugf("add user: %+v", user)
	Add(user)
	context.JSON(http.StatusOK, gin.H{"id": user.ID})
}
