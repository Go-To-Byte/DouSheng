// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

func (s *userServiceImpl) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据 Username 查询此用户是否已经注册
	userReq := newGetUserReq()
	userReq.Username = req.Username
	po, err := s.getUser(ctx, userReq)
	if err != nil {
		return nil, status.Error(codes.Unavailable,
			constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	if po.Id > 0 {
		// 用户已存在
		return nil, status.Error(codes.AlreadyExists,
			constant.Code2Msg(constant.WRONG_EXIST_USERS))
	}

	// 3、未注册-创建用户，注册-返回提示
	po = user.NewUserPo(req.Hash())
	insertRes, err := s.insert(ctx, po)

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
	userReq := newGetUserReq()
	userReq.Username = req.Username
	po, err := s.getUser(ctx, userReq)
	if err != nil {
		return nil, status.Error(codes.Unavailable,
			constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	// 若用户名或密码有误，不返回具体的用户名或者密码错误
	if po == nil || !po.CheckHash(req.Password) {
		return nil, status.Error(codes.PermissionDenied,
			constant.Code2Msg(constant.BAD_NAME_PASSWORD))
	}

	// 3、颁发Token 并返回
	response := user.NewTokenResponse(po.Id, s.token(ctx, po))

	return response, nil
}

func (s *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	// 请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("user UserInfo：参数校验失败，%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	response := user.NewUserInfoResponse()
	response.User = user.NewDefaultUser()
	// get user info, user += userInfo

	userReq := newGetUserReq()
	userReq.UserId = req.UserId
	po, err := s.getUser(ctx, userReq)
	if err != nil {
		return nil, status.Error(codes.Unavailable,
			constant.Code2Msg(constant.ERROR_ACQUIRE))
	}
	response.User = po.Po2vo()

	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)
	err = s.composeInfo(tkCtx, response.User)

	return response, err
}

func (s *userServiceImpl) UserMap(ctx context.Context, req *user.UserMapRequest) (*user.UserMapResponse, error) {

	// 请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("user UserInfo：参数校验失败，%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 1、获取用户列表 []User
	userPoRes, err := s.userList(ctx, req.UserIds)

	// 这里为什么不把错误合并在一起返回，因为有可能这里已经报错了。就没必要往后面操作了
	if err != nil {
		switch e := err.(type) {
		case *custom.Exception:
			return nil, status.Error(codes.NotFound, e.Error())
		default:
			return nil, status.Error(codes.Unknown, e.Error())
		}
	}

	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	// 2、转换为 Map[UserId] = User
	UserMap := make(map[int64]*user.User)
	for _, po := range userPoRes {
		vo := po.Po2vo()
		err = s.composeInfo(tkCtx, vo)
		if err != nil {
			return nil, err
		}
		UserMap[vo.Id] = vo
	}

	return &user.UserMapResponse{UserMap: UserMap}, nil
}

func (s *userServiceImpl) composeInfo(ctx context.Context, uResp *user.User) error {

	var (
		wait = sync.WaitGroup{}
		errs = make([]error, 0)
	)

	wait.Add(3)

	// 组合 followListCount、followerListCount、isFollow
	go func() {
		defer wait.Done()
		relationErrs := s.composeRelation(ctx, uResp)
		errs = append(errs, relationErrs...)
	}()

	// 组合 publishCount
	go func() {
		defer wait.Done()
		videoErrs := s.composeVideo(ctx, uResp)
		errs = append(errs, videoErrs...)
	}()

	// 组合 favoriteCount
	go func() {
		defer wait.Done()
		favoriteErrs := s.composeFavorite(ctx, uResp)
		errs = append(errs, favoriteErrs...)
	}()
	wait.Wait()

	// 查看后台调用时，是否有错误产生
	for _, err := range errs {
		if err != nil {
			switch e := err.(type) {
			case *custom.Exception:
				return status.Error(codes.NotFound, e.Error())
			default:
				return status.Error(codes.Unknown, e.Error())
			}
		}
	}

	return nil
}
