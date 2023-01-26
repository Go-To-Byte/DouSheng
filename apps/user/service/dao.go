// Author: BeYoung
// Date: 2023/1/26 3:29
// Software: GoLand

package service

import (
	"context"
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

func FindById(info model.UserInfo) []*model.UserInfo {
	q := query.Use(getSqlDB())
	userInfo := q.UserInfo
	result, err := userInfo.WithContext(context.Background()).Where(userInfo.ID.Eq(info.ID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return result
}

func FindByName(info model.UserInfo) []*model.UserInfo {
	q := query.Use(getSqlDB())
	userInfo := q.UserInfo
	result, err := userInfo.WithContext(context.Background()).
		Where(userInfo.Username.Eq(info.Username)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return result
}
