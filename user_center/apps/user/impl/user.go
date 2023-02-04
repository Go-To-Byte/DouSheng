// @Author: Ciusyan 2023/1/23
package impl

import (
	"Go-To-Byte/DouSheng/user_center/apps/user"
	"context"
	"fmt"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register TODO：完成注册逻辑
func (u *userServiceImpl) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// ====
	// 1、请求参数校验
	// ====
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 生成UUID
	token := xid.New().String()

	// ====
	// 2、根据 Username 查询此用户是否已经注册
	// ====

	// ====
	// 3、未注册-创建用户，注册-返回提示
	// ====

	response := user.NewTokenResponse()
	response.UserId = 1111
	response.Token = token

	return response, nil
}

// Login TODO：完成登录逻辑
func (u *userServiceImpl) Login(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {
	// ====
	// 1、请求参数校验
	// ====
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// ====
	// 2、根据用户名查询用户信息
	// ====
	if "ciusyan" != req.Username {
		return nil, fmt.Errorf("没有此用户")
	}

	// ====
	// 3、比对用户密码
	// 需要加密后对比
	// ====

	if "111" != req.Password {
		return nil, fmt.Errorf("密码错误")
	}

	// ====
	// 4、将Token放入缓存
	// ====
	// 生成UUID
	token := xid.New().String()
	response := user.NewTokenResponse()
	response.UserId = 1111
	response.Token = token

	return response, nil
}
func (u *userServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
