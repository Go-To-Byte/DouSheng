// @Author: Ciusyan 2023/1/25
package user

import (
	"github.com/Go-To-Byte/DouSheng/user_center/common/utils"
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

func (r *UserInfoRequest) Validate() error {
	return validate.Struct(r)
}

func NewTokenResponse(id int64, token string) *TokenResponse {
	return &TokenResponse{
		UserId: id,
		Token:  token,
	}
}

func NewDefaultUser() *User {
	return &User{}
}

func NewDefaultUserPo() *UserPo {
	return &UserPo{}
}

func NewUserPo(req *LoginAndRegisterRequest) *UserPo {
	return &UserPo{
		Username: req.Username,
		Password: req.Password,
	}
}

// Hash 将敏感信息做Hash
func (r *LoginAndRegisterRequest) Hash() *LoginAndRegisterRequest {
	r.Password = utils.BcryptHash(r.Password)
	return r
}

// CheckHash 比对Hash
func (u *UserPo) CheckHash(data any) bool {
	return utils.VerifyBcryptHash(data, u.Password)
}

// TableName 指明表名 -> gorm 参数映射
func (*UserPo) TableName() string {
	return AppName
}

func (po *UserPo) Po2vo() *User {
	return &User{
		Id:   po.Id,
		Name: po.Username,
	}
}

func NewUserInfoResponse() *UserInfoResponse {
	return &UserInfoResponse{}
}

func NewUserInfoRequest() *UserInfoRequest {
	return &UserInfoRequest{}
}

func NewUserMapRequest() *UserMapRequest {
	return &UserMapRequest{}
}
