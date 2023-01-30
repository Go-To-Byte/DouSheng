// Author: BeYoung
// Date: 2023/1/29 14:17
// Software: GoLand

package init

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/info/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// get gorm DB
func initDB() {
	cnd := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		models.Config.DBConfig.User, models.Config.DBConfig.Password, models.Config.DBConfig.Host,
		models.Config.DBConfig.Port, models.Config.DBConfig.Database, models.Config.DBConfig.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	var err error
	if models.DB, err = gorm.Open(driver); err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
}
