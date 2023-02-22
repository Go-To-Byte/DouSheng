// Author: BeYoung
// Date: 2023/2/2 21:08
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/relation/models"
	"github.com/Go-To-Byte/DouSheng/apps/relation/service"
)

func initPort() {
	port, err := service.GetFreePort()
	if err == nil {
		models.Config.Localhost.Port = port
	}
}
