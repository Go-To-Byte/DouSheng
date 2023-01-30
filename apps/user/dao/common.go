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

func Add(user model.User) {
	q := query.Use(models.DB)
	err := q.User.Create(&user)
	if err != nil {
		zap.S().Panicf("Failed create user: %v", err)
	}
}

func FindById(user model.User) []*model.User {
	q := query.Use(models.DB)
	u := q.User
	result, err := u.WithContext(context.Background()).
		Where(u.ID.Eq(user.ID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", user)
		return nil
	}
	return result
}

func FindByName(user model.User) []*model.User {
	q := query.Use(models.DB)
	u := q.User
	r, err := u.WithContext(context.Background()).
		Where(u.Username.Eq(user.Username)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", user)
		return nil
	}
	return r
}
