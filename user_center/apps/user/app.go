// @Author: Ciusyan 2023/1/25
package user

import (
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "user"
)

var (
	validate = validator.New()
)

func NewLoginAndRegisterRequest() *LoginAndRegisterRequest {
	return &LoginAndRegisterRequest{}
}

// Validate 参数校验
func (r *LoginAndRegisterRequest) Validate() error {
	return validate.Struct(r)
}

func NewTokenResponse() *TokenResponse {
	return &TokenResponse{}
}
