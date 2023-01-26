// Author: BeYoung
// Date: 2023/1/26 3:29
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/dal/model"
	"github.com/Go-To-Byte/DouSheng/dal/query"
	"go.uber.org/zap"
)

func Add(info model.UserInfo) {
	q := query.Use(getSqlDB())
	err := q.UserInfo.Create(&info)
	if err != nil {
		zap.S().Panicf("Failed create user: %v", err)
	}
}
