// Author: BeYoung
// Date: 2023/1/29 14:16
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/comment/models"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// initialize service config
func initConfig() {
	models.V = viper.New()
	models.V.SetConfigType("yaml") // set config type

	models.V.SetConfigFile("config/config.yml")
	if err := models.V.ReadInConfig(); err != nil {
		zap.S().Panicf("Error reading config file: %v", err)
	}

	if err := models.V.Unmarshal(&models.Config); err != nil {
		zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
	}

	models.V.WatchConfig()
	models.V.OnConfigChange(func(e fsnotify.Event) {
		if err := models.V.ReadInConfig(); err != nil {
			zap.S().Panicf("Error reading config file: %v", err)
		}

		if err := models.V.Unmarshal(&models.Config); err != nil {
			zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
		}

		zap.S().Infof("Config changed:%v", e.String())
	})
}
