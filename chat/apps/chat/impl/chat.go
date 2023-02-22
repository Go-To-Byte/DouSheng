// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/chat/apps/chat"
	"github.com/Go-To-Byte/DouSheng/chat/apps/chat/impl/dal/model"
	_ "github.com/Go-To-Byte/DouSheng/chat/apps/chat/impl/init"
	"github.com/Go-To-Byte/DouSheng/chat/apps/chat/impl/models"

	"strconv"

	"go.uber.org/zap"
)

func (c *ChatServiceImpl) MessageAction(ctx context.Context, req *chat.MessageRequest) (*chat.MessageResponse, error) {
	message := model.Message{
		ID:       models.Node.Generate().Int64(),
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Content:  req.Content,
	}

	if err := c.Add(message); err != nil {
		zap.S().Infof("failed to add message: %+v", message)
		return &chat.MessageResponse{
			StatusCode: 1,
			StatusMsg:  "failed to add message",
		}, err
	}

	zap.S().Debugf("success to add message: %+v", message)
	return &chat.MessageResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
	}, nil
}

func (c *ChatServiceImpl) MessageHistory(ctx context.Context, req *chat.MessageListRequest) (*chat.MessageListResponse, error) {
	// 聊天记录 == 我发送的消息 + 对方发送的消息
	r1 := c.MessageFindByUserIDWithToUserID(model.Message{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Content:  "",
	})
	r2 := c.MessageFindByUserIDWithToUserID(model.Message{
		ID:       0,
		UserID:   req.ToUserId,
		ToUserID: req.UserId,
		Content:  "",
	})

	messages := make([]*chat.Message, 0)

	// 合并所有聊天记录
	for i := 0; i < len(r1); i++ {
		message := &chat.Message{
			Id:         r1[i].ID,
			UserId:     r1[i].UserID,
			ToUserId:   r1[i].ToUserID,
			Content:    r1[i].Content,
			CreateTime: strconv.FormatInt(r1[i].ID>>22, 10),
		}
		messages = append(messages, message)
	}
	for i := 0; i < len(r2); i++ {
		message := &chat.Message{
			Id:         r2[i].ID,
			UserId:     r2[i].UserID,
			ToUserId:   r2[i].ToUserID,
			Content:    r2[i].Content,
			CreateTime: strconv.FormatInt(r2[i].ID>>22, 10),
		}
		messages = append(messages, message)
	}

	zap.S().Debugf("ID(%v) ==> len(message): %v", req.UserId, len(messages))
	return &chat.MessageListResponse{
		StatusCode:  0,
		StatusMsg:   "ok",
		MessageList: messages,
	}, nil
}
