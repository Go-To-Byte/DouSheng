// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/video/apps/video/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/video/apps/video/impl/dal/query"

	"go.uber.org/zap"
)

func (c *VideoServiceImpl) Add(video model.Video) (err error) {
	q := query.Use(c.db)
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

func (c *VideoServiceImpl) Delete(video model.Video) (err error) {
	q := query.Use(c.db)
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

func (c *VideoServiceImpl) VideoFindByUserID(userID int64) []*model.Video {
	q := query.Use(c.db)
	v := q.Video

	r, err := v.WithContext(context.Background()).Where(v.AuthID.Eq(userID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find AuthID: %v", userID)
	}
	return r
}

func (c *VideoServiceImpl) VideoFindByVideoID(videoID int64) []*model.Video {
	q := query.Use(c.db)
	v := q.Video

	r, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find videoID: %v", videoID)
	}
	return r
}

// VideoFindByTimeStamp 因为雪花id自带时间戳，直接根据id排序查找
func (c *VideoServiceImpl) VideoFindByTimeStamp(timeStamp int64) []*model.Video {
	q := query.Use(c.db)
	v := q.Video

	r, err := v.WithContext(context.Background()).
		Where(v.ID.Lte(timeStamp)).
		Order(v.ID.Desc()).
		Find()

	if err != nil {
		zap.S().Panicf("Failed find video time stamp: %v", timeStamp)
	}
	return r
}
