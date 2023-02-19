// Author: BeYoung
// Date: 2023/1/30 21:38
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/message/proto"
	"go.uber.org/zap"
	"strconv"
)

func (m *Chat) MessageHistory(ctx context.Context, req *proto.MessageListRequest) (*proto.MessageListResponse, error) {
	// 聊天记录 == 我发送的消息 + 对方发送的消息
	r1 := dao.MessageFindByUserIDWithToUserID(model.Message{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Content:  "",
	})
	r2 := dao.MessageFindByUserIDWithToUserID(model.Message{
		ID:       0,
		UserID:   req.ToUserId,
		ToUserID: req.UserId,
		Content:  "",
	})

	messages := make([]*proto.Message, 0)

	// 合并所有聊天记录
	for i := 0; i < len(r1); i++ {
		message := &proto.Message{
			Id:         r1[i].ID,
			UserId:     r1[i].UserID,
			ToUserId:   r1[i].ToUserID,
			Content:    r1[i].Content,
			CreateTime: strconv.FormatInt(r1[i].ID>>22, 10),
		}
		messages = append(messages, message)
	}
	for i := 0; i < len(r2); i++ {
		message := &proto.Message{
			Id:         r2[i].ID,
			UserId:     r2[i].UserID,
			ToUserId:   r2[i].ToUserID,
			Content:    r2[i].Content,
			CreateTime: strconv.FormatInt(r2[i].ID>>22, 10),
		}
		messages = append(messages, message)
	}

	zap.S().Debugf("ID(%v) ==> len(message): %v", req.UserId, len(messages))
	return &proto.MessageListResponse{
		StatusCode:  0,
		StatusMsg:   "ok",
		MessageList: messages,
	}, nil
}
