// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	kitUtils "github.com/Go-To-Byte/DouSheng/dou_kit/utils"
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

func (s *videoServiceImpl) GetVideo(ctx context.Context, req *video.GetVideoRequest) (*video.Video, error) {

	// 1、参数校验
	if req.VideoId == 0 {
		s.l.Errorf("video: GetVideo 参数校验失败：video_id 不能为空")
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、查询
	po := video.NewVideoPo()
	s.db.WithContext(ctx).Where("id = ?", req.VideoId).Find(&po)
	// 走GRPC调用，获取视频对应的用户信息
	userMap, err := s.GetUser(ctx, po)
	if err != nil {
		return nil, err
	}

	// 3、po -> vo
	return po.Po2vo(userMap), nil
}

// 获取视频流的列表
func (s *videoServiceImpl) composeFeedSetResp(ctx context.Context, pos []*video.VideoPo) (
	*video.FeedSetResponse, error) {

	set := video.NewFeedSet()
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	// 1、取出视频列表的 userIds
	userIds := kitUtils.NewSet()
	for _, po := range pos {
		userIds.Add(po.AuthorId)
	}

	// 2、获取UserMap列表
	userMapReq := user.NewUserMapRequest()
	userMapReq.UserIds = userIds.Items()

	// GRPC调用
	userMap, err := s.userServer.UserMap(ctx, userMapReq)
	if err != nil {
		return nil, err
	}

	// 3、转换并且组合用户信息
	set.VideoList = s.pos2vos(pos, userMap.GetUserMap())
	// 获取此处的最后一条的视频 创建时间， 作为下次调用的请求开始时间
	set.NextTime = utils.V2P(pos[len(pos)-1].CreatedAt)

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

	// 走GRPC调用，获取视频对应的用户信息
	userMap, err := s.GetUser(ctx, pos[0])
	if err != nil {
		return nil, err
	}

	// 转换 pos -> vos
	set.VideoList = s.pos2vos(pos, userMap)

	return set, nil
}

// 将 []videoPo -> []video，并且会组合用户信息
// pos：数据库查询到的视频列表
// userMap：用户信息 map[userId] = User
func (s *videoServiceImpl) pos2vos(pos []*video.VideoPo, userMap map[int64]*user.User) []*video.Video {

	// 判空
	set := make([]*video.Video, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set
	}

	// 再次遍历，po -> vo并且组合用户信息
	for i, po := range pos {
		// 将 po -> vo
		vo := po.Po2vo(userMap)
		set[i] = vo
	}

	return set
}

// GetUser GRPC调用，去获取用户信息
func (s *videoServiceImpl) GetUser(ctx context.Context, po *video.VideoPo) (map[int64]*user.User, error) {
	req := user.NewUserInfoRequest()
	req.UserId = po.AuthorId
	info, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return map[int64]*user.User{info.User.Id: info.User}, nil
}
