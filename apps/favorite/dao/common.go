// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/models"
	"go.uber.org/zap"
)

func Add(favorite model.Favorite) error {
	q := query.Use(models.DB)
	tx := q.Begin()
	if err := tx.Favorite.Create(&favorite); err != nil {
		zap.S().Panicf("Failed add favorite: %v", err)
		return err
	}
	if err := tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func FavoriteFindByUserID(userID int64) []*model.Favorite {
	q := query.Use(models.DB)
	f := q.Favorite

	r, err := f.WithContext(context.Background()).Where(f.UserID.Eq(userID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find follows: %v", userID)
	}
	return r
}
