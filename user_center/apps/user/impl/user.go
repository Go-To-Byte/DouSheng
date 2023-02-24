// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"

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

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据 Id 查询此用户
	userReq := NewGetUserReq()
	userReq.UserIds = append(userReq.UserIds, req.UserId)
	userPoRes, err := s.GetUser(ctx, userReq)

	if err != nil {
		switch e := err.(type) {
		case *custom.Exception:
			return nil, status.Error(codes.NotFound, e.Error())
		default:
			return nil, status.Error(codes.Unknown, e.Error())
		}
	}

	response := user.NewUserInfoResponse()
	// userPoRes[0]：因为前面只查询了一个，所以来到这里，直接取出就行
	response.User = userPoRes[0].Po2vo()

	// TODO：组合其他参数[如：关注数、粉丝数]

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
