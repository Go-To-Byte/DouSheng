// Author: BeYoung
// Date: 2023/1/28 1:41
// Software: GoLand

package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	V      *viper.Viper
	DB     *gorm.DB
	Node   *snowflake.Node
	Router *gin.Engine
	Config = ConfigYAML{}
)
var Port int
