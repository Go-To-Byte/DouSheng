// @Author: Ciusyan 2023/1/23
package impl_test

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/conf"
	// 驱动加载所有需要放入IOC的实例
	_ "github.com/Go-To-Byte/DouSheng/user_center/apps/all"
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
	// newUser.Password = "222"
	token, err := service.Login(context.Background(), newUser)

	if should.NoError(err) {
		fmt.Println(token)
		fmt.Println(token.UserId)
		fmt.Println(token.Token)
	}
}

func TestUserInfo(t *testing.T) {
	should := assert.New(t)
	req := user.NewUserInfoRequest()
	req.Token = "TgFCArASZIB5uXpJEgtvC2nB"
	req.UserId = 5
	info, err := service.UserInfo(context.Background(), req)

	if should.NoError(err) {
		fmt.Println(info)
		fmt.Println(info.User)
	}
}

func init() {

	// 加载配置文件
	if err := conf.LoadConfigFromToml("../../../etc/config.toml"); err != nil {
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
