// Author: BeYoung
// Date: 2023/1/27 5:45
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/global"
	"github.com/Go-To-Byte/DouSheng/model"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	config model.Config
	db     *gorm.DB
	node   *snowflake.Node
	router *gin.Engine
)

func init() {
	initLogger()
	initConfig()
	initRouter()
	initDB()
	initNode()

	global.DB = db
	global.Node = node
	global.Config = &config
}
