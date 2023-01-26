// Author: BeYoung
// Date: 2023/1/27 5:42
// Software: GoLand

package service

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func initNode() *snowflake.Node {
	if node != nil { // global snowflake node
		return node
	}
	var err error
	node, err = snowflake.NewNode(config.ID)
	if err != nil {
		zap.S().Panicf("New snowflake node error: %v", err)
	}
	return node
}
