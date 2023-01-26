// Author: BeYoung
// Date: 2023/1/27 5:23
// Software: GoLand

package global

import (
	"github.com/Go-To-Byte/DouSheng/model"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Node   *snowflake.Node
	Router *gin.Engine
	Config *model.Config
)
