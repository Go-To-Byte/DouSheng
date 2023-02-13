// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (*video.FeedSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVideos not implemented")
}

func (s *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (
	*video.PublishVideoResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	_, err := s.Insert(ctx, req)

	// 这里不需要返回数据，若需要，可以包装在 Mate 中返回
	return nil, err
}

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
