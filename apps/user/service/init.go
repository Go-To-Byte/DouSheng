// Author: BeYoung
// Date: 2023/1/26 3:23
// Software: GoLand

package service

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	initLogger()
	initConfig()
	initRouter()
	initDB()
	initNode()
}

// initialize service config
func initConfig() {
	V := viper.New()
	V.SetConfigType("yaml") // set config type

	V.SetConfigFile("apps/user/service/config.yml")
	if err := V.ReadInConfig(); err != nil {
		zap.S().Panicf("Error reading config file: %v", err)
	}

	if err := V.Unmarshal(&Config); err != nil {
		zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
	}

	V.WatchConfig()
	V.OnConfigChange(func(e fsnotify.Event) {
		if err := V.ReadInConfig(); err != nil {
			zap.S().Panicf("Error reading config file: %v", err)
		}

		if err := V.Unmarshal(&Config); err != nil {
			zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
		}

		zap.S().Infof("Config changed:%v", e.String())
	})
}

// set zap logger as singleton
func initLogger() {
	logger, _ := zap.NewDevelopment() // set logger as production
	zap.ReplaceGlobals(logger)
}

func initRouter() {
	Router = gin.Default()
	user := Router.Group("/douyin/user")
	{
		user.GET("/")
		user.POST("/login")
		user.POST("/register", register)
	}
}

// get gorm DB
func initDB() {
	cnd := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		Config.DBConfig.User, Config.DBConfig.Password, Config.DBConfig.Host,
		Config.DBConfig.Port, Config.DBConfig.Database, Config.DBConfig.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	var err error
	if DB, err = gorm.Open(driver); err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
}

func initNode() *snowflake.Node {
	if Node != nil { // global snowflake node
		return Node
	}
	var err error
	Node, err = snowflake.NewNode(Config.ID)
	if err != nil {
		zap.S().Panicf("New snowflake node error: %v", err)
	}
	return Node
}
