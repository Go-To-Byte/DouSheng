// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	userconstant "github.com/Go-To-Byte/DouSheng/user_center/common/constant"
)

func (s *userServiceImpl) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, exception.WithCodeMsg(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、根据 Username 查询此用户是否已经注册
	userPo := user.NewDefaultUserPo()
	userPo.Username = req.Username
	userRes, err := s.GetUser(ctx, userPo)

	if userRes != nil {
		return nil, exception.WithCodeMsg(constant.WRONG_EXIST_USERS)
	}

	// 3、未注册-创建用户，注册-返回提示
	po := user.NewUserPo(req.Hash())
	insertRes, err := s.Insert(ctx, po)

	if err != nil {
		return nil, exception.WithCodeMsg(constant.ERROR_SAVE)
	}

	// 4、颁发Token并返回
	response := user.NewTokenResponse(insertRes.Id, s.token(ctx, insertRes))

	return response, nil
}

func (s *userServiceImpl) Login(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, exception.WithCodeMsg(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、根据用户名查询用户信息
	userPo := user.NewDefaultUserPo()
	userPo.Username = req.Username
	userRes, _ := s.GetUser(ctx, userPo)

	// 若用户名或密码有误，不返回具体的用户名或者密码错误
	if userRes == nil || !userRes.CheckHash(req.Password) {
		return nil, exception.WithCodeMsg(userconstant.BAD_NAME_PASSWORD)
	}

	// 3、颁发Token 并返回
	response := user.NewTokenResponse(userRes.Id, s.token(ctx, userRes))
	return response, nil
}

func (s *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, exception.WithCodeMsg(constant.ERROR_ARGS_VALIDATE)
	}

	// 同一服务直接走内部服务调用，不用GRPC调用
	_, err := s.tokenService.ValidateToken(ctx, token.NewValidateTokenRequest(req.Token))
	if err != nil {
		return nil, exception.WithMsg("校验Token失败：%s", err.Error())
	}

	// 1、根据 Id 查询此用户
	userPo := user.NewDefaultUserPo()
	userPo.Id = req.UserId
	userPoRes, err := s.GetUser(ctx, userPo)

	if err != nil {
		return nil, exception.WithMsg(err.Error())
	}

	response := user.NewUserInfoResponse()
	response.User = user.NewUserWithPo(userPoRes)

	// TODO：组合其他参数[如：关注数、粉丝数]

	return response, nil
}
