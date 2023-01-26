// Author: BeYoung
// Date: 2023/1/27 5:42
// Software: GoLand

package service

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// get gorm DB
func initDB() *gorm.DB {
	if db != nil { // global database
		return db
	}

	cnd := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.DBConfig.User, config.DBConfig.Password, config.DBConfig.Host,
		config.DBConfig.Port, config.DBConfig.Database, config.DBConfig.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	db, err := gorm.Open(driver)
	if err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
	return db
}
