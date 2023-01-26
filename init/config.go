// Author: BeYoung
// Date: 2023/1/26 3:23
// Software: GoLand

package service

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// initialize service config
func initConfig() {
	viper.SetConfigType("yaml") // set config type

	viper.AddConfigPath("/config/debug.yml")
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Panicf("Error reading config file: %v", err)
	}

	config.ID = viper.GetInt64("id")

	err = viper.Unmarshal(&config)
	if err != nil {
		zap.S().Panicf("Failed to unmarshal sqlconfig: %v", err)
	}
}
