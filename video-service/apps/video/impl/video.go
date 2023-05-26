// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO：5、实现视频服务

func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (
	*video.FeedSetResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}

func (s *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (
	*video.PublishVideoResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (
	*video.PublishListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}

func (s *videoServiceImpl) GetVideo(ctx context.Context, req *video.GetVideoRequest) (*video.Video, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
