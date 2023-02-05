// @Author: Ciusyan 2023/1/23
package impl_test

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/conf"
	"github.com/Go-To-Byte/DouSheng/user_center/ioc"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service user.ServiceServer
)

func TestRegister(t *testing.T) {
	should := assert.New(t)
	newUser := user.NewLoginAndRegisterRequest()
	newUser.Username = "ciusyan"
	newUser.Password = "222"
	token, err := service.Register(context.Background(), newUser)

	if should.NoError(err) {
		fmt.Println(token)
		fmt.Println(token.UserId)
		fmt.Println(token.Token)
	}
}

func TestLogin(t *testing.T) {
	should := assert.New(t)
	newUser := user.NewLoginAndRegisterRequest()
	newUser.Username = "ciusyan"
	newUser.Password = "222"
	token, err := service.Login(context.Background(), newUser)

	if should.NoError(err) {
		fmt.Println(token)
		fmt.Println(token.UserId)
		fmt.Println(token.Token)
	}
}

func init() {

	// 加载配置文件
	if err := conf.LoadConfigFromToml("../../../etc/dousheng.toml"); err != nil {
		panic(err)
	}

	// 初始化全局Logger
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	// 初始化IOC容器
	if err := ioc.InitAllDependencies(); err != nil {
		panic(err)
	}

	// 从IOC中获取接口实现
	service = ioc.GetGrpcDependency(user.AppName).(user.ServiceServer)
}
