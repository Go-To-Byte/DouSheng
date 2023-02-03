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

func Add(favorite model.Favorite) (err error) {
	q := query.Use(models.DB)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	if err = tx.Favorite.Create(&favorite); err != nil {
		zap.S().Errorf("Failed add favorite: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Errorf("Failed commit: %v", err)
		return err
	}
	return nil
}

func Delete(favorite model.Favorite) (err error) {
	q := query.Use(models.DB)
	f := q.Favorite
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	if _, err = tx.Favorite.
		Where(f.UserID.Eq(favorite.UserID), f.VideoID.Eq(favorite.VideoID)).
		Update(f.Flag, 0); err != nil {
		zap.S().Errorf("Failed delete favorite: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Errorf("Failed delete: %v", err)
		return err
	}
	return nil
}

func FavoriteFindByUserID(favorite model.Favorite) []*model.Favorite {
	q := query.Use(models.DB)
	f := q.Favorite

	r, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(favorite.UserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find follows: %+v", favorite)
	}
	return r
}

func FavoriteFindByVideoID(favorite model.Favorite) []*model.Favorite {
	q := query.Use(models.DB)
	f := q.Favorite

	r, err := f.WithContext(context.Background()).
		Where(f.VideoID.Eq(favorite.VideoID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find follows: %+v", favorite)
	}
	return r
}

func FavoriteFindByUserIDWithVideoID(favorite model.Favorite) []*model.Favorite {
	q := query.Use(models.DB)
	f := q.Favorite

	r, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(favorite.UserID), f.VideoID.Eq(favorite.VideoID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find favorite: %+v", favorite)
	}
	return r
}
