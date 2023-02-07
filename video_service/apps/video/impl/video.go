// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_common/constant"
	"github.com/Go-To-Byte/DouSheng/dou_common/exception"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (*video.FeedSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVideos not implemented")
}

func (s *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (
	*video.PublishVideoResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, exception.WithCodeMsg(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、保存到数据库

	return nil, nil
}

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
