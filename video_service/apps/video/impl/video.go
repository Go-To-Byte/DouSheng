// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
)

func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (
	*video.FeedSetResponse, error) {

	// 1、查询视频列表，放入集合中 map [video_id] = video
	pos, err := s.query(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 2、根据返回的视频列表，组装用户信息
	return s.composeResp(ctx, pos)
}

func (s *videoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (
	*video.PublishVideoResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("video: PublishVideo 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	_, err := s.Insert(ctx, req)

	// 这里不需要返回数据，若需要，可以包装在 Mate 中返回
	return video.NewPublishVideoResponse(), err
}

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}

// TODO：可优化为批量查询。
func (s *videoServiceImpl) po2vo(ctx context.Context, po *video.VideoPo) (*video.Video, error) {

	// 走GRPC调用，获取用户信息
	req := user.NewUserInfoRequest()
	req.UserId = po.AuthorId
	info, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	// po -> vo
	return &video.Video{
		Id:       po.Id,
		Author:   info.User,
		PlayUrl:  utils.URLPrefix(po.PlayUrl),
		CoverUrl: utils.URLPrefix(po.CoverUrl),
		Title:    po.Title,
	}, nil
}

func (s *videoServiceImpl) composeResp(ctx context.Context, pos []*video.VideoPo) (
	*video.FeedSetResponse, error) {

	set := video.NewFeedSet()
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	// 获取最新一条的视频 创建时间， 作为下次调用的请求开始时间
	set.NextTime = utils.V2P(pos[1].CreatedAt)

	errCount := 0
	for _, po := range pos {
		// 将 po -> vo
		vo, err := s.po2vo(ctx, po)
		if err != nil {

			// 有异常，先别着急返回，给两个冗余错误，
			// 因为可能是自己业务有问题，毕竟是测试数据
			// TODO：上线删除
			errCount++
			if errCount > 1 {
				return nil, err
			}

			s.l.Errorf("video: composeResp 组合用户信息异常：%s", err.Error())
			continue
		}
		set.VideoList = append(set.VideoList, vo)
	}

	return set, nil
}
