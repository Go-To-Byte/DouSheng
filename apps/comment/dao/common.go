// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/comment/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/comment/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/comment/models"
	"go.uber.org/zap"
)

func Add(comment model.Comment) error {
	q := query.Use(models.DB)
	tx := q.Begin()
	if err := tx.Comment.Create(&comment); err != nil {
		zap.S().Panicf("Failed add favorite: %v", err)
		return err
	}
	if err := tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func Delete(comment model.Comment) error {
	q := query.Use(models.DB)
	tx := q.Begin()
	if _, err := tx.Comment.Delete(&comment); err != nil {
		zap.S().Panicf("Failed add favorite: %v", err)
		return err
	}
	if err := tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func CommentFindByVideoID(videoID int64) []*model.Comment {
	q := query.Use(models.DB)
	c := q.Comment

	r, err := c.WithContext(context.Background()).Where(c.VideoID.Eq(videoID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find follows: %v", videoID)
	}
	return r
}
