// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/info/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/info/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/info/models"
	"go.uber.org/zap"
)

func Add(info model.Info) {
	q := query.Use(models.DB)
	err := q.Info.Create(&info)
	if err != nil {
		zap.S().Panicf("Failed create user: %v", err)
	}
}

func FindById(info model.Info) []*model.Info {
	q := query.Use(models.DB)
	u := q.Info
	r, err := u.WithContext(context.Background()).
		Where(u.ID.Eq(info.ID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return r
}

func FindByName(info model.Info) []*model.Info {
	q := query.Use(models.DB)
	u := q.Info
	r, err := u.WithContext(context.Background()).
		Where(u.Name.Eq(info.Name)).Find()
	if err != nil {
		zap.S().Panicf("Failed find user_info: %+v", info)
		return nil
	}
	return r
}
