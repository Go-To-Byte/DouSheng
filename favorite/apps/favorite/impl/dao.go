// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite/impl/dal/query"
	"go.uber.org/zap"
)

func (f *FavoriteServiceImpl) Add(favorite model.Favorite) (err error) {
	q := query.Use(f.db)

	r := f.FavoriteFindByUserIDAndVideoID(favorite)
	if len(r) != 0 && r[0].Flag == 1 {
		return nil
	}

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

func (f *FavoriteServiceImpl) Delete(favorite model.Favorite) (err error) {
	q := query.Use(f.db)
	fav := q.Favorite
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	if _, err = tx.Favorite.
		Where(fav.UserID.Eq(favorite.UserID), fav.VideoID.Eq(favorite.VideoID)).
		Update(fav.Flag, 0); err != nil {
		zap.S().Errorf("Failed delete favorite: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Errorf("Failed delete: %v", err)
		return err
	}
	return nil
}

func (f *FavoriteServiceImpl) FavoriteFindByUserID(favorite model.Favorite) []*model.Favorite {
	q := query.Use(f.db)
	fav := q.Favorite

	r, err := fav.WithContext(context.Background()).
		Where(fav.UserID.Eq(favorite.UserID), fav.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find follows: %+v", favorite)
	}
	return r
}

func (f *FavoriteServiceImpl) FavoriteFindByVideoID(favorite model.Favorite) []*model.Favorite {
	q := query.Use(f.db)
	fav := q.Favorite

	r, err := fav.WithContext(context.Background()).
		Where(fav.VideoID.Eq(favorite.VideoID), fav.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find follows: %+v", favorite)
	}
	return r
}

func (f *FavoriteServiceImpl) FavoriteFindByUserIDAndVideoID(favorite model.Favorite) []*model.Favorite {
	q := query.Use(f.db)
	fav := q.Favorite

	r, err := fav.WithContext(context.Background()).
		Where(fav.UserID.Eq(favorite.UserID), fav.VideoID.Eq(favorite.VideoID), fav.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find favorite: %+v", favorite)
	}
	return r
}
