// Author: BeYoung
// Date: 2023/2/2 19:42
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/models"
	"github.com/Go-To-Byte/DouSheng/apps/user/service"
)

func initPort() {
	port, err := service.GetFreePort()
	if err == nil {
		models.Port = port
	}
}
