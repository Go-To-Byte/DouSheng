// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception/custom"
	kitUtils "github.com/Go-To-Byte/DouSheng/dou-kit/utils"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"

	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
)

func (s *videoServiceImpl) FeedVideos(ctx context.Context, req *video.FeedVideosRequest) (
	*video.FeedSetResponse, error) {

	// 1、查询视频列表，放入集合中 map [video_id] = video
	pos, err := s.query(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 2、根据返回的视频列表，组装用户信息
	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	return s.composeFeedSetResp(tkCtx, pos)
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
	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	return s.composeUserListResp(tkCtx, pos)
}

func (s *videoServiceImpl) GetVideo(ctx context.Context, req *video.GetVideoRequest) (*video.Video, error) {

	// 1、参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("video GetVideo: 参数校验失败，%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、查询
	po := video.NewVideoPo()
	s.db.WithContext(ctx).Where("id = ?", req.VideoId).Find(&po)

	// 1、获取Tk + loginID
	tkReq := token.NewValidateTokenRequest(req.Token)
	uidFromTk, err := s.tokenService.GetUIDFromTk(ctx, tkReq)
	if err != nil {
		s.l.Errorf("video composeUserListResp：从tk获取UID失败，%s", err.Error())
		return nil, err
	}

	// 2、转换并且组合用户信息
	vos, err := s.composeVideoResp(ctx, newComposePosInfo([]*video.VideoPo{po}, req.Token, uidFromTk.UserId))
	if err != nil {
		return nil, err
	}

	if len(vos) != 1 {
		return nil, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、po -> vo
	return vos[0], nil
}

func (s *videoServiceImpl) ComposeVideoCount(ctx context.Context, req *video.PublishListCountRequest) (
	*video.PublishListCountResponse, error) {

	// 1、校验参数[防止GRPC调用时参数异常]
	if req.UserId <= 0 {
		s.l.Errorf("video: PublishListCount ：请正确的传递用户ID，%d", req.UserId)
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 查询用户视频数量
	resp, err := s.getTotalCount(ctx, req.UserId)

	if err != nil {
		// 查询失败
		s.l.Errorf("video ComposeVideoCount：", err.Error())
		return resp, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	return resp, nil
}

type composePosInfo struct {
	// 视频的 pos
	pos []*video.VideoPo
	// 登录用户的 tk
	tk string
	// 登录用户的 ID
	loginUid int64
}

func newComposePosInfo(pos []*video.VideoPo, tk string, loginUid int64) *composePosInfo {
	return &composePosInfo{
		pos:      pos,
		tk:       tk,
		loginUid: loginUid,
	}
}

func (s *videoServiceImpl) composeVideoResp(ctx context.Context, info *composePosInfo) (
	[]*video.Video, error) {

	set := make([]*video.Video, 10)

	// 1、校验参数
	if info.pos == nil || len(info.pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	var (
		wait             = sync.WaitGroup{}
		errs             = make([]error, 0)
		userMap          = make(map[int64]*user.User, 10)
		favoriteCountMap = make(map[int64]*favorite.FavoriteMap, 10)
		commentCountMap  = make(map[int64]int64, 10)
	)

	// 3、取出视频列表的 videoIds
	videoIds := kitUtils.NewSet()
	for _, po := range info.pos {
		videoIds.Add(po.Id)
	}

	wait.Add(3)

	// (1)、组合用户信息
	go func() {
		defer wait.Done()

		um, err := s.composeUser(ctx, info.pos, info.tk)
		errs = append(errs, err)
		userMap = um
	}()

	// (2)、组合点赞信息
	go func() {
		defer wait.Done()

		fm, err := s.composeFavorite(ctx, videoIds.Items(), info.loginUid)
		errs = append(errs, err)
		favoriteCountMap = fm
	}()

	// (3)、组合评论信息
	go func() {
		defer wait.Done()

		cm, err := s.composeComment(ctx, videoIds.Items())
		errs = append(errs, err)
		commentCountMap = cm
	}()

	wait.Wait()

	// 查看后台调用时，是否有错误产生
	for _, err := range errs {
		if err != nil {
			switch e := err.(type) {
			case *custom.Exception:
				return set, status.Error(codes.NotFound, e.Error())
			default:
				return set, status.Error(codes.Unknown, e.Error())
			}
		}
	}

	// 拿到信息后，去组合信息 转换为 vos
	vos := s.pos2vos(info.pos, newComposeInfo(userMap, favoriteCountMap, commentCountMap))

	return vos, nil
}

// 获取视频流的列表
func (s *videoServiceImpl) composeFeedSetResp(ctx context.Context, pos []*video.VideoPo) (
	*video.FeedSetResponse, error) {

	set := video.NewFeedSet()
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	// 1、获取Tk + loginID
	var (
		tk       = ""
		logonUid = int64(0)
	)

	tkValue := ctx.Value(constant.REQUEST_TOKEN)
	if tkValue != nil && tkValue.(*string) != nil {

		tkReq := token.NewValidateTokenRequest(*tkValue.(*string))
		// 这里需要使用 ValidateToken，因为这里是登录后才有Token，防止原先过期的Token请求，却拿到数据了
		uidFromTk, err := s.tokenService.ValidateToken(ctx, tkReq)
		if err != nil {
			s.l.Errorf("video composeFeedSetResp：从tk获取UID失败，%s", err.Error())
			tk = ""
			logonUid = 0
		} else {
			// 验证正确才设置，要不然就默认值
			tk = *tkValue.(*string)
			logonUid = uidFromTk.GetUserId()
		}
	}

	// 2、转换并且组合用户信息
	vos, err := s.composeVideoResp(ctx, newComposePosInfo(pos, tk, logonUid))
	if err != nil {
		return set, err
	}

	set.VideoList = vos
	// 获取此处的最后一条的视频 创建时间， 作为下次调用的请求开始时间
	set.NextTime = kitUtils.V2P(pos[len(pos)-1].CreatedAt)

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

	// 1、获取Tk + loginID
	tk := kitUtils.TokenStrFromCtx(ctx)
	tkReq := token.NewValidateTokenRequest(tk)
	uidFromTk, err := s.tokenService.GetUIDFromTk(ctx, tkReq)
	if err != nil {
		s.l.Errorf("video composeUserListResp：从tk获取UID失败，%s", err.Error())
		return set, err
	}

	// 2、转换并且组合用户信息
	vos, err := s.composeVideoResp(ctx, newComposePosInfo(pos, tk, uidFromTk.UserId))
	if err != nil {
		return set, err
	}

	set.VideoList = vos

	return set, nil
}

type composeVoInfo struct {
	// [author_id] = User
	userMap map[int64]*user.User
	// [video_id] = FavoriteMap (favorite_count + isFavorite)
	favoriteCountMap map[int64]*favorite.FavoriteMap
	// [video_id] = comment_id
	commentCountMap map[int64]int64
}

func newComposeInfo(userMap map[int64]*user.User,
	favoriteCountMap map[int64]*favorite.FavoriteMap, commentCountMap map[int64]int64) *composeVoInfo {
	return &composeVoInfo{
		userMap:          userMap,
		favoriteCountMap: favoriteCountMap,
		commentCountMap:  commentCountMap,
	}
}

// 将 []videoPo -> []video，并且会组合用户信息、点赞、评论信息
// pos：数据库查询到的视频列表
func (s *videoServiceImpl) pos2vos(pos []*video.VideoPo, info *composeVoInfo) []*video.Video {

	// 判空
	set := make([]*video.Video, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set
	}

	// 再次遍历，po -> vo并且组合用户信息
	for i, po := range pos {
		// 将 po -> vo
		vo := po.Po2vo()
		vo.Author = info.userMap[po.AuthorId]                         // 用户信息
		vo.IsFavorite = info.favoriteCountMap[po.Id].IsFavorite       // 是否点赞
		vo.FavoriteCount = info.favoriteCountMap[po.Id].FavoriteCount // 视频赞数量
		vo.CommentCount = info.commentCountMap[po.Id]                 // 评论数量
		set[i] = vo
	}

	return set
}

// getUser GRPC调用，去获取用户信息
func (s *videoServiceImpl) getUser(ctx context.Context, po *video.VideoPo) (map[int64]*user.User, error) {

	req := user.NewUserInfoRequest()
	req.UserId = po.AuthorId
	req.Token = kitUtils.TokenStrFromCtx(ctx)
	info, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return map[int64]*user.User{info.User.Id: info.User}, nil
}
