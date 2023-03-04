// @Author: Hexiaoming 2023/2/18
package impl

import (
	"context"
	"time"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
)

// 根据接收者userid和当前用户token获取聊天消息列表(DAO层)
func (s *messageServiceImpl) getChatMessageListByUserId(ctx context.Context, req *message.ChatMessageListRequest) (
	[]*message.MessagePo, error) {

	// 获取用户ID
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := s.tokenService.GetUIDFromTk(ctx, tokenReq)

	if err != nil {
		s.l.Errorf("message: 验证用户Token失败: %s", err.Error())
		return nil, err
	}

	// 发送方userId & 接收方userId
	fromUserId := validatedToken.GetUserId()

	db := s.db.WithContext(ctx)
	set := make([]*message.MessagePo, 50)

	// 查询双方聊天的所有消息

	// 这里是为了适配客户端传入的时间戳不统一，做的额外操作
	nowS := time.Now().Unix()
	if req.PreMsgTime < nowS {
		req.PreMsgTime = time.Unix(req.PreMsgTime, 0).UnixMilli()
	}
	// 毫秒
	lastTime := time.UnixMilli(req.PreMsgTime).Add(2 * time.Second).Unix()

	// 构建sql并且查询
	db.Where("created_at > ? AND created_at < ?", lastTime, nowS).
		Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
			req.ToUserId, fromUserId, fromUserId, req.ToUserId).
		Order("created_at asc").Find(&set)

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

	// 获取用户ID
	tokenReq := token.NewValidateTokenRequest(req.Token)
	tk, err := s.tokenService.GetUIDFromTk(ctx, tokenReq)

	if err != nil {
		s.l.Errorf(err.Error())
		// GRPC 调用，不需要继续包装了
		return nil, err
	}

	MessagePo := message.NewMessagePo(req)
	MessagePo.FromUserId = tk.GetUserId()
	return MessagePo, nil
}
