// Author: BeYoung
// Date: 2023/1/30 21:38
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/models"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/proto"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao/dal/model"
	"go.uber.org/zap"
)

func (m *Message) MessageAction(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	message := model.Message{
		ID:       models.Node.Generate().Int64(),
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Content:  req.Content,
	}

	if err := dao.Add(message); err != nil {
		zap.S().Infof("failed to add message: %+v", message)
		return &proto.MessageResponse{
			StatusCode: 1,
			StatusMsg:  "failed to add message",
		}, err
	}

	zap.S().Debugf("success to add message: %+v", message)
	return &proto.MessageResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
	}, nil
}
