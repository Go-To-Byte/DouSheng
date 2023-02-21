// @Author: Ciusyan 2023/1/24
package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

// GetUser 根据用户名称获取用户
func (s *userServiceImpl) GetUser(ctx context.Context, po *user.UserPo) (*user.UserPo, error) {
	db := s.db.WithContext(ctx)
	if po.Username != "" {
		db = db.Where("username = ?", po.Username)
	}

	if po.Id != 0 {
		db = db.Where("id = ?", po.Id)
	}

	// 查询
	db = db.Find(po)

	if db.RowsAffected == 0 {
		return nil, exception.WithStatusCode(constant.WRONG_USER_NOT_EXIST)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return po, nil
}

// Insert 创建用户
func (s *userServiceImpl) Insert(ctx context.Context, user *user.UserPo) (*user.UserPo, error) {

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
