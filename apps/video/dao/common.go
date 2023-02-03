// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/video/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/video/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/video/models"

	"go.uber.org/zap"
)

func Add(video model.Video) (err error) {
	q := query.Use(models.DB)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = tx.Video.Create(&video); err != nil {
		zap.S().Panicf("Failed add video: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit add: %v", err)
		return err
	}
	return nil
}

func Delete(video model.Video) (err error) {
	q := query.Use(models.DB)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.Video.Delete(&video); err != nil {
		zap.S().Panicf("Failed delete video: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit delete: %v", err)
		return err
	}
	return nil
}

func VideoFindByUserID(userID int64) []*model.Video {
	q := query.Use(models.DB)
	v := q.Video

	r, err := v.WithContext(context.Background()).Where(v.AuthID.Eq(userID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find AuthID: %v", userID)
	}
	return r
}

func VideoFindByVideoID(videoID int64) []*model.Video {
	q := query.Use(models.DB)
	v := q.Video

	r, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find videoID: %v", videoID)
	}
	return r
}
