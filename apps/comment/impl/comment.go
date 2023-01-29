// @Author: Ciusyan 2023/1/29
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/comment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *commentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (c *commentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
