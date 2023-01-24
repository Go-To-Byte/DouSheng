// @Author: Ciusyan 2023/1/23
package user

import (
	"context"
	"github.com/go-playground/validator"
)

var (
	validate = validator.New()
)

// Service 定义User Service 的接口
type Service interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, request *LoginAndRegisterRequest) (*Token, error)
	Login(ctx context.Context, request *LoginAndRegisterRequest) (*Token, error)
}

func NewUser() *User {
	return &User{}
}

// User 模型的定义
type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func NewToken(userId int64, token string) *Token {
	return &Token{
		UserId: userId,
		Token:  token,
	}
}

// Token 用于登录和注册的模型
type Token struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func NewLoginAndRegisterRequest() *LoginAndRegisterRequest {
	return &LoginAndRegisterRequest{}
}

// LoginAndRegisterRequest 登录和注册的请求模型
type LoginAndRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Validate 校验结构体参数
func (r *LoginAndRegisterRequest) Validate() error {
	return validate.Struct(r)
}
