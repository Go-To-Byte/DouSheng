// @Author: Ciusyan 2023/1/23
package impl_test

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/Go-To-Byte/DouSheng/apps/user/impl"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service user.Service
)

func TestCreateUser(t *testing.T) {
	should := assert.New(t)
	newUser := user.NewLoginAndRegisterRequest()
	newUser.Username = "ciusyan"
	newUser.Password = "xxxx"
	token, err := service.CreateUser(context.Background(), newUser)

	if should.NoError(err) {
		fmt.Println(token)
		fmt.Println(token.UserId)
		fmt.Println(token.Token)
	}

}

func init() {

	if err := conf.LoadConfigFromToml("../../../etc/dousheng.toml"); err != nil {
		panic(err)
	}

	// 初始化全局Logger
	zap.DevelopmentSetup()
	// 接口实现
	service = impl.NewUserServiceImpl()
}
