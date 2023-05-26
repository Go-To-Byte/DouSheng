package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// "github.com/Go-To-Byte/DouSheng/user-service/apps/user"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/message"
)

func (s *messageServiceImpl) ChatMessageList(ctx context.Context, req *message.ChatMessageListRequest) (
	*message.ChatMessageListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}

func (s *messageServiceImpl) ChatMessageAction(ctx context.Context, req *message.ChatMessageActionRequest) (
	*message.ChatMessageActionResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
