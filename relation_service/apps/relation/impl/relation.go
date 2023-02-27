package impl

import (
	"context"
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
	return s.composeFollowListResp(ctx, pos)
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
	return s.composeFollowerListResp(ctx, pos)
}

func (s *relationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListRequest) (
	*relation.FriendListResponse, error) {

	s.l.Errorf("relation: FriendList ", req)

	// 1、校验参数[防止GRPC调用时参数异常]
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FriendList 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据用户ID获取粉丝列表
	pos, err := s.getFollowerListByUserId(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 3、组合用户关注列表信息
	return s.composeFriendListResp(ctx, pos, req.Token)
}

func (s *relationServiceImpl) FollowAction(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.FollowActionResponse, error) {

	s.l.Errorf("relation: Token ：%s", req.Token)

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("relation: FollowAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	if req.ActionType == constant.FOLLOW_ACTION {
		s.l.Errorf("relation: FollowAction 关注", req)
		_, err := s.insert(ctx, req)
		if err != nil {
			return relation.NewFollowActionResponse(), err
		}
	} else if req.ActionType == constant.UNFOLLOW_ACTION {
		s.l.Errorf("relation: FollowAction 取消关注", req)
		_, err := s.update(ctx, req)
		if err != nil {
			return relation.NewFollowActionResponse(), err
		}
	} else {
		s.l.Errorf("relation: FollowAction 未知的动作类型：", req.ActionType)
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 这里不需要返回数据，若需要，可以包装在 Mate 中返回
	return relation.NewFollowActionResponse(), nil

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

func (s *relationServiceImpl) composeFriendListResp(ctx context.Context, pos []*relation.UserFollowerPo, userToken string) (
	*relation.FriendListResponse, error) {

	set := relation.NewFriendListResponse()
	if pos == nil || len(pos) <= 0 {
		// 可能存在用户粉丝列表为空, 不应该抛出异常而范围空值
		return set, nil
	}

	// 转换 pos -> vos
	vos, err := s.friendPos2Vos(ctx, pos, userToken)
	if err != nil {
		return set, err
	}
	set.FriendList = vos

	return set, nil
}

func (s *relationServiceImpl) followPos2Vos(ctx context.Context, pos []*relation.UserFollowPo) (
	[]*user.User, error) {

	set := make([]*user.User, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	errCount := 0
	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.followPo2Vo(ctx, po)
		if err != nil {

			errCount++
			if errCount > 1 {
				return nil, err
			}

			s.l.Errorf("relation: composeFollowListResp 组合关注用户信息异常：%s", err.Error())
			continue
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

	errCount := 0
	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.followerPo2Vo(ctx, po)
		if err != nil {
			errCount++
			if errCount > 1 {
				return nil, err
			}

			s.l.Errorf("relation: composeFollowListResp 组合关注用户信息异常：%s", err.Error())
			continue
		}
		set[i] = vo
	}

	return set, nil
}

func (s *relationServiceImpl) friendPos2Vos(ctx context.Context, pos []*relation.UserFollowerPo, userToken string) (
	[]*relation.UserFriend, error) {

	set := make([]*relation.UserFriend, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		return set, nil
	}

	errCount := 0
	for i, po := range pos {
		// 将 po -> vo
		vo, err := s.followerPo2FriendVo(ctx, po, userToken)
		if err != nil {
			errCount++
			if errCount > 1 {
				return nil, err
			}

			s.l.Errorf("relation: composeFriendListResp 组合关注用户信息异常：%s", err.Error())
			continue
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

func (s *relationServiceImpl) followerPo2FriendVo(ctx context.Context, po *relation.UserFollowerPo, userToken string) (
	*relation.UserFriend, error) {

	// 走GRPC调用，获取用户粉丝信息
	req := user.NewUserInfoRequest()
	req.UserId = po.FollowerId

	resp, err := s.userServer.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	toUser := resp.User
	// 根据粉丝id与用户id去查找最新的聊天信息
	// TODO 这样做效率不高, 待优化
	msgReq := message.NewChatMessageListRequest()
	msgReq.ToUserId = po.FollowerId
	msgReq.Token = userToken

	// 走GRPC调用, 获取最新聊天信息
	msgResp, err := s.messageService.ChatMessageList(ctx, msgReq)
	if err != nil {
		s.l.Errorf(err.Error())
		return nil, exception.GrpcErrWrapper(err)
	}

	msgList := msgResp.MessageList
	content := ""
	if len(msgList) > 0 {
		content = msgList[0].Content
	}

	// userInfo.User
	// user.is_follow = true
	// po -> vo
	return &relation.UserFriend{
		Id:            toUser.Id,
		Name:          toUser.Name,
		FollowCount:   toUser.FollowCount,
		FollowerCount: toUser.FollowerCount,
		IsFollow:      toUser.IsFollow,
		Message:       content,
		MsgType:       1,
	}, nil
}
