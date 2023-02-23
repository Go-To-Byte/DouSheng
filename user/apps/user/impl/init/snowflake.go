// Author: BeYoung
// Date: 2023/1/29 14:21
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment/impl/models"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func initNode() *snowflake.Node {
	if models.Node != nil { // global snowflake node
		return models.Node
	}
	var err error
	models.Node, err = snowflake.NewNode(1) // todo: random node
	if err != nil {
		zap.S().Panicf("New snowflake node error: %v", err)
	}
	return models.Node
}
