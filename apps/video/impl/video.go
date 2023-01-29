// @Author: Ciusyan 2023/1/29
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/video"
	"github.com/Go-To-Byte/DouSheng/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (v *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (*video.FeedSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVideos not implemented")
}
func (v *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (*common.CodeAndMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (v *videoServiceImpl) PublishList(ctx context.Context, req *common.UserIDAndTokenRequest) (*video.PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
