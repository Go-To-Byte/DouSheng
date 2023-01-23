// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user"
)

func (i *UserServiceImpl) CreateUser(ctx context.Context, request *user.LoginAndRegisterRequest) (*user.Token, error) {
	i.l.Debug("创建用户")
	i.l.Debugf("创建用户：%s", request.Username)
	return nil, nil
}

func (i *UserServiceImpl) Login(ctx context.Context, request *user.LoginAndRegisterRequest) (*user.Token, error) {
	return nil, nil
}
