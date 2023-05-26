// @Author: Ciusyan 2023/1/24
package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
)

func newGetUserReq() *getUserReq {
	return &getUserReq{}
}

// getUserReq 查询用户信息
type getUserReq struct {
	// 用户名
	Username string `json:"username"`
	// IDS
	UserId int64 `json:"user_id"`
}

// userList 根据用户名称获取用户
func (s *userServiceImpl) userList(ctx context.Context, userIds []int64) ([]*user.UserPo, error) {

	pos := make([]*user.UserPo, 0)
	if userIds == nil || len(userIds) <= 0 {
		s.l.Errorf("user userList：你的参数可能有问题哟~")
		return pos, nil
	}

	// 查询
	db := s.db.WithContext(ctx).Where("id IN ?", userIds).Find(&pos)

	if db.Error != nil {
		return nil, db.Error
	}

	if db.RowsAffected == 0 {
		return nil, exception.WithStatusCode(constant.WRONG_USER_NOT_EXIST)
	}

	return pos, nil
}

// 通过user_id 或 username 查找用户
func (s *userServiceImpl) getUser(ctx context.Context, req *getUserReq) (*user.UserPo, error) {

	db := s.db.WithContext(ctx)

	po := user.NewDefaultUserPo()
	if req.Username != "" && req.UserId <= 0 {
		db = db.Where("username = ?", req.Username)
	} else if req.UserId > 0 && req.Username == "" {
		db = db.Where("id = ?", req.UserId)
	} else {
		s.l.Errorf("user userList：你的参数可能有问题哟~")
		return po, nil
	}

	db = db.Find(po)

	if db.Error != nil {
		s.l.Errorf("user getUser：%s", db.Error.Error())
		return nil, db.Error
	}

	return po, nil
}

// insert 创建用户
func (s *userServiceImpl) insert(ctx context.Context, user *user.UserPo) (*user.UserPo, error) {

	res := s.db.WithContext(ctx).Create(user)

	// TODO：统一异常处理
	if res.Error != nil {
		return nil, fmt.Errorf("创建用户失败：%s", res.Error.Error())
	}

	return user, nil
}

func (s *userServiceImpl) token(ctx context.Context, po *user.UserPo) (accessToken string) {
	// 颁发Token
	tkReq := token.NewIssueTokenRequest(po)
	tk, err := s.tokenService.IssueToken(ctx, tkReq)

	// 若Token颁发失败，不要报错，打印日志即可
	if err != nil {
		accessToken = ""
		s.l.Errorf("Token颁发失败：%s", err.Error())
	} else {
		accessToken = tk.AccessToken
		s.l.Infof("Token颁发成功：%s", accessToken)
	}
	return
}
