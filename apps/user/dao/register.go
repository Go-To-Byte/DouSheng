// Author: BeYoung
// Date: 2023/1/26 3:29
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/dal/query"
	model2 "github.com/Go-To-Byte/DouSheng/apps/user/models"
	"go.uber.org/zap"
)

func Add(info model.UserInfo) {
	q := query.Use(model2.DB)
	err := q.UserInfo.Create(&info)
	if err != nil {
		zap.S().Panicf("Failed create user: %v", err)
	}
}

func FindById(info model.UserInfo) []*model.UserInfo {
	q := query.Use(model2.DB)
	userInfo := q.UserInfo
	result, err := userInfo.WithContext(context.Background()).Where(userInfo.ID.Eq(info.ID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return result
}

func FindByName(info model.UserInfo) []*model.UserInfo {
	q := query.Use(model2.DB)
	userInfo := q.UserInfo
	result, err := userInfo.WithContext(context.Background()).
		Where(userInfo.Username.Eq(info.Username)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return result
}
