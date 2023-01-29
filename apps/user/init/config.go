// Author: BeYoung
// Date: 2023/1/29 14:16
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/mod"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// initialize service config
func initConfig() {
	V := viper.New()
	V.SetConfigType("yaml") // set config type

	V.SetConfigFile("apps/user/service/config.yml")
	if err := V.ReadInConfig(); err != nil {
		zap.S().Panicf("Error reading config file: %v", err)
	}

	if err := V.Unmarshal(&mod.Config); err != nil {
		zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
	}

	V.WatchConfig()
	V.OnConfigChange(func(e fsnotify.Event) {
		if err := V.ReadInConfig(); err != nil {
			zap.S().Panicf("Error reading config file: %v", err)
		}

		if err := V.Unmarshal(&mod.Config); err != nil {
			zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
		}

		zap.S().Infof("Config changed:%v", e.String())
	})
}
