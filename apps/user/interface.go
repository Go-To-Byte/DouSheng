// @Author: Ciusyan 2023/1/24
package user

import "context"

// Service 定义User Service 的接口
type Service interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, request *LoginAndRegisterRequest) (*Token, error)
	Login(ctx context.Context, request *LoginAndRegisterRequest) (*Token, error)
}
