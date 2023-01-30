// Author: BeYoung
// Date: 2023/1/30 21:38
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/message/dao"
	"github.com/Go-To-Byte/DouSheng/apps/message/proto"
	"go.uber.org/zap"
	"strconv"
)

func (m *Message) MessageHistory(ctx context.Context, req *proto.MessageListRequest) (*proto.MessageListResponse, error) {
	// 聊天记录 == 我发送的消息 + 对方发送的消息
	r1 := dao.MessageFindByUserID(req.UserId)
	r2 := dao.MessageFindByUserID(req.ToUserId)

	messages := make([]*proto.Message, len(r1)+len(r2))

	// 合并所有聊天记录
	for i := 0; i < len(r1); i++ {
		message := &proto.Message{
			Id:         r1[i].ID,
			Content:    r1[i].Content,
			CreateTime: strconv.FormatInt(r1[i].ID>>22, 10),
		}
		messages = append(messages, message)
	}
	for i := 0; i < len(r2); i++ {
		message := &proto.Message{
			Id:         r2[i].ID,
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
