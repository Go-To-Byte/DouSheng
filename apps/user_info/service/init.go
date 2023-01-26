// Author: BeYoung
// Date: 2023/1/26 3:23
// Software: GoLand

package service

import "go.uber.org/zap"

type sqlConfig struct {
	driverName string
	host       string
	port       string
	user       string
	password   string
	database   string
}

func initConfig() {
}

func setLogger() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}

func setSqlServer() {

}
