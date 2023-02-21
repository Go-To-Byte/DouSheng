// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"

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
	return s.composeFeedSetResp(ctx, pos)
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

func (s *videoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (
	*video.PublishListResponse, error) {
	// 1、校验参数[防止GRPC调用时参数异常]
	if err := req.Validate(); err != nil {
		s.l.Errorf("video: PublishList 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据用户ID获取视频列表
	pos, err := s.listFromUserId(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、组合视频的用户信息
	return s.composeUserListResp(ctx, pos)
}

// 获取视频流的列表
func (s *videoServiceImpl) composeFeedSetResp(ctx context.Context, pos []*video.VideoPo) (
	*video.FeedSetResponse, error) {

	set := video.NewFeedSet()
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	// 获取最新一条的视频 创建时间， 作为下次调用的请求开始时间
	set.NextTime = utils.V2P(pos[len(pos)-1].CreatedAt)

	vos, err := s.pos2vos(ctx, pos, nil)
	if err != nil {
		return set, err
	}
	set.VideoList = vos
	return set, nil
}

// 获取用户主页的视频列表
func (s *videoServiceImpl) composeUserListResp(ctx context.Context, pos []*video.VideoPo) (
	*video.PublishListResponse, error) {

	set := video.NewPublishListResponse()
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	// 走GRPC调用，获取用户信息
	// 因为这里查询的都是同一个用户的视频，所以可以先查出视频的用户信息
	req := user.NewUserInfoRequest()
	req.UserId = pos[0].AuthorId
	info, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	// 转换 pos -> vos
	vos, err := s.pos2vos(ctx, pos, info)
	if err != nil {
		return set, err
	}
	set.VideoList = vos

	return set, nil
}

// 将 []videoPo -> []video，并且会组合用户信息
// pos：数据库查询到的视频列表
// userInfo：用户信息
func (s *videoServiceImpl) pos2vos(ctx context.Context, pos []*video.VideoPo,
	userInfo *user.UserInfoResponse) ([]*video.Video, error) {

	set := make([]*video.Video, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	errCount := 0
	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.po2vo(ctx, po, userInfo)
		if err != nil {

			// 有异常，先别着急返回，给两个冗余错误，
			// 因为可能是自己业务有问题，毕竟是测试数据
			// TODO：上线删除
			errCount++
			if errCount > 1 {
				return nil, err
			}

			s.l.Errorf("video: composeFeedSetResp 组合用户信息异常：%s", err.Error())
			continue
		}
		set[i] = vo
	}

	return set, nil
}

// TODO：可优化为批量查询。
// 将 videoPo -> video，并且会组合用户信息
// po：单个videoPo对象
// userInfo：用户信息
func (s *videoServiceImpl) po2vo(ctx context.Context, po *video.VideoPo,
	userInfo *user.UserInfoResponse) (*video.Video, error) {

	// 也可以是单个查询
	if userInfo == nil {
		// 走GRPC调用，获取用户信息
		req := user.NewUserInfoRequest()
		req.UserId = po.AuthorId
		info, err := s.userServer.UserInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		userInfo = info
	}

	// po -> vo
	return &video.Video{
		Id:       po.Id,
		Author:   userInfo.User,
		PlayUrl:  utils.URLPrefix(po.PlayUrl),
		CoverUrl: utils.URLPrefix(po.CoverUrl),
		Title:    po.Title,
	}, nil
}
