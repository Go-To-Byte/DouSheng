// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/user/models"
	"go.uber.org/zap"
)

func Add(info model.User) {
	q := query.Use(models.DB)
	err := q.User.Create(&info)
	if err != nil {
		zap.S().Panicf("Failed create user: %v", err)
	}
}

func FindById(info model.User) []*model.User {
	q := query.Use(models.DB)
	u := q.User
	result, err := u.WithContext(context.Background()).Where(u.ID.Eq(info.ID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return result
}

func FindByName(info model.User) []*model.User {
	q := query.Use(models.DB)
	u := q.User
	r, err := u.WithContext(context.Background()).
		Where(u.Username.Eq(info.Username)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return r
}
