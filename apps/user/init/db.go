// Author: BeYoung
// Date: 2023/1/29 14:17
// Software: GoLand

package init

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user/mod"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// get gorm DB
func initDB() {
	cnd := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		mod.Config.DBConfig.User, mod.Config.DBConfig.Password, mod.Config.DBConfig.Host,
		mod.Config.DBConfig.Port, mod.Config.DBConfig.Database, mod.Config.DBConfig.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	var err error
	if mod.DB, err = gorm.Open(driver); err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
}
