// Author: BeYoung
// Date: 2023/1/26 3:23
// Software: GoLand

package service

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

var (
	serviceID     int64
	sqlconfig     SqlConfig
	sqlDB         *gorm.DB
	snowflakeNode *snowflake.Node
)

type SqlConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

// initialize service config
func initConfig() {
	viper.SetConfigType("yaml") // set config type

	// TODO: 改为相对路径
	f, err := os.Open("C:\\Users\\21941\\OneDrive\\Code\\DouSheng\\apps\\user\\service\\config.yml") // read config file
	if err != nil {
		zap.S().Panicf("Failed to open config: %v", err)
	}

	config, err := io.ReadAll(f)
	if err != nil {
		zap.S().Panicf("Failed to read config: %v", err)
	}

	err = viper.ReadConfig(bytes.NewBuffer(config)) // read config
	if err != nil {
		zap.S().Panicf("Failed to read config: %v", err)
	}

	serviceID = viper.GetInt64("snowflakeID")
	zap.S().Debugf("service_id: %d", serviceID)

	err = viper.Unmarshal(&sqlconfig)
	if err != nil {
		zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
	}
}

// set zap logger as singleton
func initLogger() {
	logger, _ := zap.NewDevelopment() // set logger as production
	zap.ReplaceGlobals(logger)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// get gorm DB
func initSqlServer() *gorm.DB {
	if sqlDB != nil { // global database
		return sqlDB
	}

	cnd := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",
		sqlconfig.User, sqlconfig.Password, sqlconfig.Host,
		sqlconfig.Port, sqlconfig.Database, sqlconfig.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	sqlDB, err := gorm.Open(driver)
	if err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
	return sqlDB
}

func initSnowflakeNode() *snowflake.Node {
	if snowflakeNode != nil { // global snowflake node
		return snowflakeNode
	}
	var err error
	snowflakeNode, err = snowflake.NewNode(serviceID)
	if err != nil {
		zap.S().Panicf("New snowflake node error: %v", err)
	}
	return snowflakeNode
}

func init() {
	initLogger()
	initConfig()
	initRouter()
	initSqlServer()
	initSnowflakeNode()
}
