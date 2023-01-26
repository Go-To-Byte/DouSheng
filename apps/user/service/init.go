// Author: BeYoung
// Date: 2023/1/26 3:23
// Software: GoLand

package service

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	serviceID     int64
	sqlconfig     sqlConfig
	sqlDB         *gorm.DB
	snowflakeNode *snowflake.Node
)

type sqlConfig struct {
	driverName string `mapstructure:"driver"`
	host       string `mapstructure:"host"`
	port       string `mapstructure:"port"`
	user       string `mapstructure:"username"`
	password   string `mapstructure:"password"`
	database   string `mapstructure:"database"`
	charset    string `mapstructure:"charset"`
}

// initialize service config
func initConfig() {
	viper.SetConfigType("yaml")                  // set config type
	configFile, err := os.ReadFile("config.yml") // read config file
	if err != nil {                              // error
		zap.S().Panicf("Failed to read config: %v", err)
	}

	err = viper.ReadConfig(bytes.NewReader(configFile)) // read config
	if err != nil {
		zap.S().Panicf("Failed to read config: %v", err)
	}

	serviceID = viper.GetInt64("service_id")

	err = viper.Unmarshal(&sqlconfig)
	if err != nil {
		zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
	}
}

// set zap logger as singleton
func initLogger() {
	logger, _ := zap.NewProduction() // set logger as production
	zap.ReplaceGlobals(logger)
}

// get gorm DB
func initSqlServer() *gorm.DB {
	if sqlDB != nil { // global database
		return sqlDB
	}
	driver := mysql.Open(fmt.Sprintf("mysql://%s:%s@%s:%s/%s?charset=%s",
		sqlconfig.user, sqlconfig.password, sqlconfig.host,
		sqlconfig.port, sqlconfig.database, sqlconfig.charset))
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
