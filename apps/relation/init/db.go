// Author: BeYoung
// Date: 2023/1/29 14:17
// Software: GoLand

package init

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/relation/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// get gorm DB
func initDB() {
	cnd := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=%s&parseTime=True",
		models.Config.DB.User, models.Config.DB.Password, models.Config.DB.Host,
		models.Config.DB.Port, models.Config.DB.Database, models.Config.DB.Charset)
	driver := mysql.Open(cnd)

	zap.S().Debugf("connecting to %s", cnd)

	var err error
	if models.DB, err = gorm.Open(driver); err != nil {
		zap.S().Panicf("Failed to connect database: %v", err)
	}
}
