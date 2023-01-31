// Author: BeYoung
// Date: 2023/2/1 0:19
// Software: GoLand

package inits

import "go.uber.org/zap"

func initZap() {
	logger, _ := zap.NewDevelopment() // set logger as production
	zap.ReplaceGlobals(logger)
}
