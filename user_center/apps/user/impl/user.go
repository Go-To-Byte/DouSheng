// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

func (s *userServiceImpl) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据 Username 查询此用户是否已经注册
	userReq := NewGetUserReq()
	userReq.Username = req.Username
	userRes, err := s.GetUser(ctx, userReq)

	if userRes != nil && len(userRes) == 1 {
		return nil, status.Error(codes.AlreadyExists,
			constant.Code2Msg(constant.WRONG_EXIST_USERS))
	}

	// 3、未注册-创建用户，注册-返回提示
	po := user.NewUserPo(req.Hash())
	insertRes, err := s.Insert(ctx, po)

	if err != nil {
		return nil, status.Error(codes.Unknown,
			constant.Code2Msg(constant.ERROR_SAVE))
	}

	// 4、颁发Token并返回
	response := user.NewTokenResponse(insertRes.Id, s.token(ctx, insertRes))

	return response, nil
}

func (s *userServiceImpl) Login(ctx context.Context, req *user.LoginAndRegisterRequest) (
	*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据用户名查询用户信息
	userReq := NewGetUserReq()
	userReq.Username = req.Username
	userRes, _ := s.GetUser(ctx, userReq)

	// 若用户名或密码有误，不返回具体的用户名或者密码错误
	if userRes == nil || len(userRes) != 1 || !userRes[0].CheckHash(req.Password) {
		return nil, status.Error(codes.PermissionDenied,
			constant.Code2Msg(constant.BAD_NAME_PASSWORD))
	}

	// 3、颁发Token 并返回
	response := user.NewTokenResponse(userRes[0].Id, s.token(ctx, userRes[0]))
	return response, nil
}

func (s *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	var (
		errors   []error
		wait     = sync.WaitGroup{}
		response = user.NewUserInfoResponse()
	)

	// 请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// get user info, user += userInfo
	wait.Add(1)
	go func() {
		defer wait.Done()
		userReq := NewGetUserReq()
		userReq.UserIds = append(userReq.UserIds, req.UserId)
		userPoRes, err := s.GetUser(ctx, userReq)
		if len(userPoRes) == 0 || err != nil {
			errors = append(errors, err)
			return
		}
		response.User.Id = userPoRes[0].Id
		response.User.Name = userPoRes[0].Username
		var avatar = ""
		var signature = "hello world"
		var background = ""
		response.User.Avatar = &avatar              // TODO: database
		response.User.Signature = &signature        // TODO: database
		response.User.BackgroundImage = &background // TODO: database
	}()

	// get follow list count, user += followListCount, user += isFollow
	go func() {
		wait.Add(1)
		defer wait.Done()
		followListReq := relation.NewFollowListRequest()
		followListReq.Token = req.Token
		followListReq.UserId = req.UserId
		followList, err := s.relation.FollowList(ctx, followListReq)
		followCount := int64(len(followList.UserList))
		response.User.FollowCount = &followCount
		errors = append(errors, err)

		// 获取用户ID
		tokenReq := token.NewValidateTokenRequest(req.Token)
		t, err := s.tokenService.ValidateToken(ctx, tokenReq)
		errors = append(errors, err)

		for _, follow := range followList.UserList {
			if follow.Id == t.GetUserId() {
				response.User.IsFollow = true
			}
		}
	}()

	// get follower list, user += followerList
	go func() {
		wait.Add(1)
		defer wait.Done()
		followerListReq := relation.NewFollowerListRequest()
		followerListReq.Token = req.Token
		followerListReq.UserId = req.UserId
		followerList, err := s.relation.FollowerList(ctx, followerListReq)
		followerCount := int64(len(followerList.UserList))
		response.User.FollowerCount = &followerCount
		errors = append(errors, err)
	}()

	// get publish list, user += publishCount
	go func() {
		wait.Add(1)
		defer wait.Done()
		publishListReq := video.NewPublishListRequest()
		publishListReq.Token = req.Token
		publishListReq.UserId = req.UserId
		publishList, err := s.video.PublishList(ctx, publishListReq)
		publishCount := int64(len(publishList.VideoList))
		response.User.WorkCount = &publishCount
		errors = append(errors, err)
	}()

	// get favorite list, user += favoriteCount
	go func() {
		wait.Add(1)
		defer wait.Done()
		favoriteListReq := &favorite.GetFavoriteListRequest{ // TODO: favorite model's naming specification
			Token:  req.Token,
			UserId: req.UserId,
		}
		favoriteList, err := s.favorite.GetFavoriteList(ctx, favoriteListReq) // TODO: favorite model's naming specification
		favoriteCount := int64(len(favoriteList.VideoList))
		response.User.FavoriteCount = &favoriteCount
		errors = append(errors, err)
	}()

	wait.Wait()
	for _, err := range errors {
		if err != nil {
			switch e := err.(type) {
			case *custom.Exception:
				return nil, status.Error(codes.NotFound, e.Error())
			default:
				return nil, status.Error(codes.Unknown, e.Error())
			}
		}
	}

	return response, nil
}

func (s *userServiceImpl) UserMap(ctx context.Context, req *user.UserMapRequest) (*user.UserMapResponse, error) {
	// 1、获取用户列表 []User
	userReq := NewGetUserReq()
	userReq.UserIds = req.UserIds
	userPoRes, err := s.GetUser(ctx, userReq)
	if err != nil {
		switch e := err.(type) {
		case *custom.Exception:
			return nil, status.Error(codes.NotFound, e.Error())
		default:
			return nil, status.Error(codes.Unknown, e.Error())
		}
	}

	// 2、转换为 Map[UserId] = User
	UserMap := make(map[int64]*user.User)
	for _, po := range userPoRes {
		UserMap[po.Id] = po.Po2vo()
	}

	return &user.UserMapResponse{UserMap: UserMap}, nil
}
