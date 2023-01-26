// Author: BeYoung
// Date: 2023/1/27 5:41
// Software: GoLand

package service

import "go.uber.org/zap"

// set zap logger as singleton
func initLogger() {
	logger, _ := zap.NewDevelopment() // set logger as production
	zap.ReplaceGlobals(logger)
}
