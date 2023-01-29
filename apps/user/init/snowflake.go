// Author: BeYoung
// Date: 2023/1/29 14:21
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/mod"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func initNode() *snowflake.Node {
	if mod.Node != nil { // global snowflake node
		return mod.Node
	}
	var err error
	mod.Node, err = snowflake.NewNode(mod.Config.ID)
	if err != nil {
		zap.S().Panicf("New snowflake node error: %v", err)
	}
	return mod.Node
}
