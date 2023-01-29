// Author: BeYoung
// Date: 2023/1/29 14:17
// Software: GoLand

package init

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// get gorm DB
func initDB() {
	cnd := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		model.Config.DBConfig.User, model.Config.DBConfig.Password, model.Config.DBConfig.Host,
		model.Config.DBConfig.Port, model.Config.DBConfig.Database, model.Config.DBConfig.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	var err error
	if model.DB, err = gorm.Open(driver); err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
}
