// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (*video.FeedSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVideos not implemented")
}

func (s *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (*video.PublishVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
