package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"

	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
)

func (s *relationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListRequest) (
	*relation.FollowListResponse, error) {

	// 1、校验参数[防止GRPC调用时参数异常]
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FollowList 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据用户ID获取关注列表
	pos, err := s.getFollowListByUserId(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、组合用户关注列表信息
	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	return s.composeFollowListResp(tkCtx, pos)
}

func (s *relationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (
	*relation.FollowerListResponse, error) {

	// 1、校验参数[防止GRPC调用时参数异常]
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FollowerList 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据用户ID获取粉丝列表
	pos, err := s.getFollowerListByUserId(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、组合用户粉丝列表信息
	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	return s.composeFollowerListResp(tkCtx, pos)
}

func (s *relationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListRequest) (
	*relation.FriendListResponse, error) {

	// 1、校验参数[防止GRPC调用时参数异常]
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FriendList 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据用户ID获取粉丝列表
	pos, err := s.getFriendListByUId(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、组合用户关注列表信息
	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	return s.composeFriendListResp(tkCtx, pos)
}

func (s *relationServiceImpl) FollowAction(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.FollowActionResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FollowAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 这里不需要返回数据，若需要，可以包装在 Mate 中返回 [主要是 grpc 调用，不能返回 nil，会序列化失败]
	return relation.NewFollowActionResponse(), s.followAction(ctx, req)
}

func (s *relationServiceImpl) ListCount(ctx context.Context, req *relation.ListCountRequest) (
	*relation.ListCountResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FollowAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、获取数目
	resp, err := s.getFavoriteCount(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	return resp, nil
}

func (s *relationServiceImpl) IsFollower(ctx context.Context, req *relation.UserFollowPo) (
	*relation.IsFollowerResponse, error) {

	isFollow, err := s.exist(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	resp := relation.NewIsFollowerResponse()
	resp.MyFollower = isFollow

	return resp, nil
}

func (s *relationServiceImpl) followAction(ctx context.Context, req *relation.FollowActionRequest) error {

	// 1、获取登录用户的ID
	tokenReq := token.NewValidateTokenRequest(req.Token)
	loginUid, err := s.tokenService.GetUIDFromTk(ctx, tokenReq)
	if err != nil {
		s.l.Errorf(err.Error())
		return err
	}

	// 2、获取 UserFollow po 对象
	userFollowPo := relation.NewUserFollowPo()
	userFollowPo.UserId = loginUid.UserId
	userFollowPo.FollowId = req.ToUserId
	userFollowPo.FollowFlag = req.ActionType

	// 3、获取 UserFollower po 对象
	userFollowerPo := relation.NewUserFollowerPo()
	userFollowerPo.UserId = req.ToUserId
	userFollowerPo.FollowerId = loginUid.UserId
	userFollowerPo.FollowerFlag = req.ActionType

	po := newSavePo(userFollowPo, userFollowerPo)
	po.action = req.ActionType

	// 如果是关注操作，需要检查是否是 再次关注的
	if req.ActionType == relation.ActionType_FOLLOW_ACTION {
		exist, err := s.exist(ctx, userFollowPo)
		if err != nil {
			return err
		}

		// 如果存在用户，，将 操作类型 调整为 ActionType_UN_FOLLOW_ACTION (再次关注)
		if exist {
			po.action = relation.ActionType_AGAIN_FOLLOW
		}

	}

	// 4、保存操作
	return s.save(ctx, po)
}

func (s *relationServiceImpl) composeFollowListResp(ctx context.Context, pos []*relation.UserFollowPo) (
	*relation.FollowListResponse, error) {

	set := relation.NewFollowListResponse()
	if pos == nil || len(pos) <= 0 {
		// 可能存在用户关注列表为空, 不应该抛出异常而范围空值
		return set, nil
	}

	// 转换 pos -> vos
	vos, err := s.followPos2Vos(ctx, pos)
	if err != nil {
		return set, err
	}
	set.UserList = vos

	return set, nil
}

func (s *relationServiceImpl) composeFollowerListResp(ctx context.Context, pos []*relation.UserFollowerPo) (
	*relation.FollowerListResponse, error) {

	set := relation.NewFollowerListResponse()
	if pos == nil || len(pos) <= 0 {
		// 可能存在用户粉丝列表为空, 不应该抛出异常而范围空值
		return set, nil
	}

	// 转换 pos -> vos
	vos, err := s.followerPos2Vos(ctx, pos)
	if err != nil {
		return set, err
	}

	set.UserList = vos

	return set, nil
}

func (s *relationServiceImpl) composeFriendListResp(ctx context.Context, pos []*relation.UserFollowerPo) (
	*relation.FriendListResponse, error) {

	set := relation.NewFriendListResponse()
	if pos == nil || len(pos) <= 0 {
		// 可能存在用户粉丝列表为空, 不应该抛出异常而范围空值
		return set, nil
	}

	// 转换 pos -> vos
	vos, err := s.friendPos2Vos(ctx, pos)
	if err != nil {
		return set, err
	}

	set.UserList = vos

	return set, nil
}

func (s *relationServiceImpl) followPos2Vos(ctx context.Context, pos []*relation.UserFollowPo) (
	[]*user.User, error) {

	set := make([]*user.User, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.followPo2Vo(ctx, po)
		if err != nil {
			s.l.Errorf("relation: composeFollowListResp 组合关注用户信息异常：%s", err.Error())
			return set, err
		}

		set[i] = vo
	}

	return set, nil
}

func (s *relationServiceImpl) followerPos2Vos(ctx context.Context, pos []*relation.UserFollowerPo) (
	[]*user.User, error) {

	set := make([]*user.User, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.followerPo2Vo(ctx, po)
		if err != nil {
			s.l.Errorf("relation: composeFollowListResp 组合关注用户信息异常：%s", err.Error())
			return nil, err
		}
		set[i] = vo
	}

	return set, nil
}

func (s *relationServiceImpl) friendPos2Vos(ctx context.Context, pos []*relation.UserFollowerPo) (
	[]*relation.UserFriend, error) {

	set := make([]*relation.UserFriend, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.followerPo2FriendVo(ctx, po)
		if err != nil {
			s.l.Errorf("relation: composeFriendListResp 组合关注用户信息异常：%s", err.Error())
			return set, err
		}

		set[i] = vo
	}

	return set, nil
}

func (s *relationServiceImpl) followPo2Vo(ctx context.Context, po *relation.UserFollowPo) (
	*user.User, error) {

	// 走GRPC调用，获取用户信息
	req := user.NewUserInfoRequest()
	req.UserId = po.FollowId
	req.Token = utils.TokenStrFromCtx(ctx)
	userInfo, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO 修改VO对象关注状态
	// userInfo.User
	// user.is_follow = true
	// po -> vo
	return userInfo.User, nil
}

func (s *relationServiceImpl) followerPo2Vo(ctx context.Context, po *relation.UserFollowerPo) (
	*user.User, error) {

	// 走GRPC调用，获取用户信息
	req := user.NewUserInfoRequest()
	req.UserId = po.FollowerId
	req.Token = utils.TokenStrFromCtx(ctx)
	userInfo, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	// TODO 修改VO对象关注状态
	// userInfo.User
	// user.is_follow = true
	// po -> vo
	return userInfo.User, nil
}

func (s *relationServiceImpl) followerPo2FriendVo(ctx context.Context, po *relation.UserFollowerPo) (
	*relation.UserFriend, error) {

	// 走GRPC调用，获取用户粉丝信息
	req := user.NewUserInfoRequest()
	req.UserId = po.UserId
	tk := utils.TokenStrFromCtx(ctx)
	req.Token = tk
	resp, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	friend := relation.NewUserFriend(resp.User)

	// 根据粉丝id与用户id去查找最新的聊天信息
	// TODO 这样做效率不高, 待优化 （可用携程优化一下）
	msgReq := message.NewChatMessageListRequest()
	msgReq.ToUserId = po.UserId
	msgReq.Token = tk
	// 走GRPC调用, 获取最新聊天信息
	msgResp, err := s.messageService.ChatMessageList(ctx, msgReq)
	if err != nil {
		s.l.Errorf(err.Error())
		return nil, exception.GrpcErrWrapper(err)
	}

	msgList := msgResp.MessageList
	if len(msgList) > 0 {
		friend.Message = msgList[len(msgList)-1].Content
	}

	// userInfo.User
	// user.is_follow = true
	// po -> vo

	return friend, nil
}
