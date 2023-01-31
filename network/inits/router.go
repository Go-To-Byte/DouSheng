// Author: BeYoung
// Date: 2023/2/1 0:13
// Software: GoLand

package inits

import (
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	models.Router = gin.Default()
	models.Router.Use(milddlewares.Cors())
}
