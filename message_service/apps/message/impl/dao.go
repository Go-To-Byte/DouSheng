// @Author: Hexiaoming 2023/2/18
package impl

import (
	"context"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
)

// 根据接收者userid和当前用户token获取聊天消息列表(DAO层)
func (s *messageServiceImpl) getChatMessageListByUserId(ctx context.Context, toUserId int64, userToken string) (
	[]*message.MessagePo, error) {
	
	tokenReq := token.NewValidateTokenRequest(userToken)

	// 获取用户ID
	validatedToken, err := s.tokenService.ValidateToken(ctx, tokenReq)

	if err != nil {
		s.l.Errorf("message: 验证用户Token失败: %s", err.Error())
		return nil, err
	}

	// 发送方userId & 接收方userId
	var fromUserId = validatedToken.GetUserId()

	db := s.db.WithContext(ctx)
	set := make([]*message.MessagePo, 50)

	// 查询
	s.db.Where("to_user_id=?", toUserId).Where("from_user_id=?", fromUserId).Order("created_at desc").Find(&set)
	if db.Error != nil {
		s.l.Errorf("message: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}


func (s *messageServiceImpl) insert(ctx context.Context, req *message.ChatMessageActionRequest) (
	*message.MessagePo, error) {
	
	// 写入消息记录
	messagePo, err := s.getMessagePo(ctx, req)
	if err != nil {
		return nil, err
	}

	tx := s.db.WithContext(ctx).Create(messagePo)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}

	return nil, err
}

func (s *messageServiceImpl) getMessagePo(ctx context.Context, req *message.ChatMessageActionRequest) (
	*message.MessagePo, error) {

	tokenReq := token.NewValidateTokenRequest(req.Token)

	// 这里主要是为了获取 用户ID
	validatedToken, err := s.tokenService.ValidateToken(ctx, tokenReq)

	if err != nil {
		s.l.Errorf(err.Error())
		// GRPC 调用，不需要继续包装了
		return nil, err
	}

	MessagePo := message.NewMessagePo(req)
	MessagePo.FromUserId = validatedToken.GetUserId()
	return MessagePo, nil
}
