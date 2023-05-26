package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
)

// TODO：完善评论接口

// CommentAction 实现评论操作功能
func (c *commentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (
	*comment.CommentActionResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}

// GetCommentList 获取评论列表，成功返回GetCommentListResponse对象地址，失败返回错误信息
func (c *commentServiceImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (
	*comment.GetCommentListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
