// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception/custom"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
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

	// 4、Token不在这儿颁发
	response := user.NewTokenResponse(insertRes.Id)

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

	// 3、Token不在这里颁发
	response := user.NewTokenResponse(po.Id)

	return response, nil
}

func (s *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	// 请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("user UserInfo：参数校验失败，%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	userReq := newGetUserReq()
	userReq.UserId = req.UserId
	po, err := s.getUser(ctx, userReq)

	if err != nil {
		return nil, status.Error(codes.Unavailable,
			constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	response := user.NewUserInfoResponse()
	response.User = po.Po2vo()

	return response, nil
}

func (s *userServiceImpl) UserMap(ctx context.Context, req *user.UserMapRequest) (*user.UserMapResponse, error) {

	// 1、获取用户列表 []User
	userPoRes, err := s.userList(ctx, req.UserIds)

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
