// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/chat/apps/chat/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/chat/apps/chat/impl/dal/query"

	"go.uber.org/zap"
)

func (c *ChatServiceImpl) Add(message model.Message) (err error) {
	q := query.Use(c.db)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	if err = tx.Message.Create(&message); err != nil {
		zap.S().Panicf("Failed add message: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func (c *ChatServiceImpl) MessageFindByUserIDWithToUserID(message model.Message) []*model.Message {
	q := query.Use(c.db)
	m := q.Message

	r, err := m.WithContext(context.Background()).
		Where(m.UserID.Eq(message.UserID), m.ToUserID.Eq(message.ToUserID)).
		Find()
	if err != nil {
		zap.S().Panicf("Failed find message: %+v", message)
	}
	return r
}
