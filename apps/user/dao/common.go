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

func Add(user model.User) (err error) {
	q := query.Use(models.DB)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = tx.User.Create(&user); err != nil {
		zap.S().Panicf("Failed create user(%+v) ==> err: %v", user, err)
		return err
	}

	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit user(%+v) ==> err: %v", user, err)
		return err
	}

	return nil
}

func UserFindById(user model.User) []*model.User {
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

func UserFindByName(user model.User) []*model.User {
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
