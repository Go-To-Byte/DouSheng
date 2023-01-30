// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/message/models"

	"go.uber.org/zap"
)

func Add(message model.Message) error {
	q := query.Use(models.DB)
	tx := q.Begin()
	if err := tx.Message.Create(&message); err != nil {
		zap.S().Panicf("Failed add message: %v", err)
		return err
	}
	if err := tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func MessageFindByUserID(userID int64) []*model.Message {
	q := query.Use(models.DB)
	m := q.Message

	r, err := m.WithContext(context.Background()).Where(m.UserID.Eq(userID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find follows: %v", userID)
	}
	return r
}
