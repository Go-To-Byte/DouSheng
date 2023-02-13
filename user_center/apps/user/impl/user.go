// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
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
	userPo := user.NewDefaultUserPo()
	userPo.Username = req.Username
	userRes, err := s.GetUser(ctx, userPo)

	if userRes != nil {
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
	userPo := user.NewDefaultUserPo()
	userPo.Username = req.Username
	userRes, _ := s.GetUser(ctx, userPo)

	// 若用户名或密码有误，不返回具体的用户名或者密码错误
	if userRes == nil || !userRes.CheckHash(req.Password) {
		return nil, status.Error(codes.PermissionDenied,
			constant.Code2Msg(constant.BAD_NAME_PASSWORD))
	}

	// 3、颁发Token 并返回
	response := user.NewTokenResponse(userRes.Id, s.token(ctx, userRes))
	return response, nil
}

func (s *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// GRPC调用
	_, err := s.tokenService.ValidateToken(ctx, token.NewValidateTokenRequest(req.Token))
	if err != nil {
		// 因为走GRPC调用，肯定会返回 Status类型的错误
		return nil, err
	}

	// 2、根据 Id 查询此用户
	userPo := user.NewDefaultUserPo()
	userPo.Id = req.UserId
	userPoRes, err := s.GetUser(ctx, userPo)

	if err != nil {
		switch e := err.(type) {
		case *custom.Exception:
			return nil, status.Error(codes.NotFound, e.Error())
		default:
			return nil, status.Error(codes.Unknown, e.Error())
		}
	}

	response := user.NewUserInfoResponse()
	response.User = user.NewUserWithPo(userPoRes)

	// TODO：组合其他参数[如：关注数、粉丝数]

	return response, nil
}
