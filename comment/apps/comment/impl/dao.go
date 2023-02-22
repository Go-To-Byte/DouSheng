// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment/impl/dal/query"
	"go.uber.org/zap"
)

func (c *CommentServiceImpl) Add(comment model.Comment) (err error) {
	q := query.Use(c.db)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	if err = tx.Comment.Create(&comment); err != nil {
		zap.S().Panicf("Failed add favorite: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func (c *CommentServiceImpl) Delete(comment model.Comment) (err error) {
	q := query.Use(c.db)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	if _, err = tx.Comment.Delete(&comment); err != nil {
		zap.S().Panicf("Failed add favorite: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func (c *CommentServiceImpl) CommentFindByVideoID(videoID int64) []*model.Comment {
	q := query.Use(c.db)
	cmt := q.Comment

	r, err := cmt.WithContext(context.Background()).Where(cmt.VideoID.Eq(videoID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find follows: %v", videoID)
	}
	return r
}
