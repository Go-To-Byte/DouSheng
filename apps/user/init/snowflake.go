// Author: BeYoung
// Date: 2023/1/29 14:21
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/model"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func initNode() *snowflake.Node {
	if model.Node != nil { // global snowflake node
		return model.Node
	}
	var err error
	model.Node, err = snowflake.NewNode(model.Config.ID)
	if err != nil {
		zap.S().Panicf("New snowflake node error: %v", err)
	}
	return model.Node
}
