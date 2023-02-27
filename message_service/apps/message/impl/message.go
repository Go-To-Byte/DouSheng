// @Author: Hexiaoming 2023/2/18
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	// "github.com/Go-To-Byte/DouSheng/user_center/apps/user"

	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
)

func (s *messageServiceImpl) ChatMessageList(ctx context.Context, req *message.ChatMessageListRequest) (
	*message.ChatMessageListResponse, error) {

	// 1、校验参数[防止GRPC调用时参数异常]
	if err := req.Validate(); err != nil {
		s.l.Errorf("message: ChatMessageList 参数校验失败：%s", req)
		s.l.Errorf("message: ChatMessageList 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据接收用户ID和当前用户token获取消息列表
	pos, err := s.getChatMessageListByUserId(ctx, req.ToUserId, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、组合用户关注列表信息
	return s.composeChatMessageListResp(ctx, pos)
}

func (s *messageServiceImpl) ChatMessageAction(ctx context.Context, req *message.ChatMessageActionRequest) (
	*message.ChatMessageActionResponse, error) {
	
	s.l.Errorf("message: Token ：%s", req.Token)

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("message: ChatMessageAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	_, err := s.insert(ctx, req)

	// 这里不需要返回数据，若需要，可以包装在 Mate 中返回
	return message.NewChatMessageActionResponse(), err
}

func (s *messageServiceImpl) composeChatMessageListResp(ctx context.Context, pos []*message.MessagePo) (
	*message.ChatMessageListResponse, error) {
	
	set := message.NewChatMessageListResponse()
	if pos == nil || len(pos) <= 0 {
		// 可能存在聊天记录为空, 不应该抛出异常而返回空值
		return set, nil
	}

	// 转换 pos -> vos
	vos, err := s.messagePos2Vos(ctx, pos)
	if err != nil {
		return set, err
	}
	set.MessageList = vos

	return set, nil
}

// 这里Message的PO与VO属性一致, 为了后续拓展方便还是实现了以下方法
func (s *messageServiceImpl) messagePos2Vos(ctx context.Context, pos []*message.MessagePo) (
	[]*message.Message, error) {
	set := make([]*message.Message, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	errCount := 0
	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.messagePo2Vo(ctx, po)
		if err != nil {
			errCount++
			if errCount > 1 {
				return nil, err
			}
			s.l.Errorf("message: composeMessageListResp 组合聊天消息异常：%s", err.Error())
			continue
		}
		set[i] = vo
	}
	return set, nil
}

func (s *messageServiceImpl) messagePo2Vo(ctx context.Context, po *message.MessagePo) (
	*message.Message, error) {
	
	return &message.Message{
		Id: 		po.Id,
		ToUserId: 	po.ToUserId,
		FromUserId: po.FromUserId,
		Content: 	po.Content,
		CreatedAt: 	po.CreatedAt,
	}, nil
}

