// @Author: Ciusyan 2023/1/24
package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

// GetUserByName 根据用户名称获取用户
func (s *userServiceImpl) GetUserByName(ctx context.Context, name string) (*user.UserPo, error) {

	po := user.NewDefaultUserPo()
	res := s.db.WithContext(ctx).Where("username = ?", name).Find(po)

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}

	if res.Error != nil {
		return nil, res.Error
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
