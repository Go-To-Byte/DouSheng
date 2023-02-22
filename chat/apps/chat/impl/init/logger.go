// Author: BeYoung
// Date: 2023/1/29 14:16
// Software: GoLand

package init

import "go.uber.org/zap"

// set zap logger as singleton
func initLogger() {
	logger, _ := zap.NewDevelopment() // set logger as production
	zap.ReplaceGlobals(logger)
}
