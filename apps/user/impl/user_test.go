// @Author: Ciusyan 2023/1/23
package impl_test

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/Go-To-Byte/DouSheng/apps/user/impl"
	"github.com/infraboard/mcube/logger/zap"
	"testing"
)

var (
	service user.Service
)

func TestCreateUser(t *testing.T) {
	newUser := user.NewLoginAndRegisterRequest()
	newUser.Username = "ciusyan"
	service.CreateUser(context.Background(), newUser)
}

func init() {
	// 初始化全局Logger
	zap.DevelopmentSetup()
	// 接口实现
	service = impl.NewUserServiceImpl()
}
