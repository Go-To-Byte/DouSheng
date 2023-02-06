// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/common/constant"
	"github.com/Go-To-Byte/DouSheng/user_center/common/exception"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register TODO：完成注册逻辑
func (s *userServiceImpl) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, exception.WithCodeMsg(constant.BAD_ARGS_VALIDATE)
	}

	// 2、根据 Username 查询此用户是否已经注册
	userRes, err := s.GetUserByName(ctx, req.Username)

	// 用户已存在
	if userRes != nil {
		return nil, exception.WithCodeMsg(constant.WARNING_USER_EXIST)
	}

	// 3、未注册-创建用户，注册-返回提示
	po := user.NewUserPo(req.Hash())
	insertRes, err := s.Insert(ctx, po)

	if err != nil {
		return nil, exception.WithCodeMsg(constant.BAD_SAVE)
	}

	token := xid.New().String()
	response := user.NewTokenResponse(insertRes.Id, token)

	// TODO：存储token[redis or mongo] 加上验证中间件
	return response, nil
}

// Login TODO：完成登录逻辑
func (s *userServiceImpl) Login(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 2、根据用户名查询用户信息
	userRes, _ := s.GetUserByName(ctx, req.Username)
	// 3、不返回具体的用户名或者密码错误
	if userRes == nil || !userRes.CheckHash(req.Password) {
		return nil, exception.WithCodeMsg(constant.BAD_NAME_PASSWORD)
	}

	// 4、TODO:将Token放入缓存
	token := xid.New().String()

	response := user.NewTokenResponse(userRes.Id, token)
	return response, nil
}
func (u *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
