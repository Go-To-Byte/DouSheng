// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register TODO：完成注册逻辑
func (s *userServiceImpl) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 2、根据 Username 查询此用户是否已经注册
	userRes, err := s.GetUserByName(ctx, req.Username)

	response := user.NewTokenResponse()
	var msg string

	// 用户已存在
	if userRes != nil {
		// TODO：1、小工具：取指针的值，2、封装响应[成功、失败]
		msg = "用户已存在"
		response.StatusMsg = &msg
		response.StatusCode = 60001
		return response, err
	}

	// 3、未注册-创建用户，注册-返回提示
	po := user.NewUserPo(req.Hash())
	insertRes, err := s.Insert(ctx, po)

	if err != nil {
		msg = "保存失败"
		response.StatusMsg = &msg
		response.StatusCode = 50001
		return nil, err
	}

	msg = "注册成功"
	token := xid.New().String()

	response.StatusMsg = &msg
	response.StatusCode = 0
	response.UserId = insertRes.Id
	response.Token = token

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
	response := user.NewTokenResponse()
	var msg string

	// 3、不返回具体的用户名或者密码错误
	if userRes == nil || !userRes.CheckHash(req.Password) {
		msg = "用户名或密码错误"
		response.StatusMsg = &msg
		response.StatusCode = 60002
		return response, fmt.Errorf(msg)
	}

	// 4、将Token放入缓存
	msg = "登录成功"
	token := xid.New().String()

	response.StatusMsg = &msg
	response.StatusCode = 0
	response.UserId = userRes.Id
	response.Token = token

	return response, nil
}
func (u *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
