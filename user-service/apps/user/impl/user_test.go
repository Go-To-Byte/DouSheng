// @Author: Ciusyan 2023/1/23
package impl_test

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
	// 驱动加载所有需要放入IOC的实例
	_ "github.com/Go-To-Byte/DouSheng/user-service/common/all"
)

var (
	service user.ServiceServer
)

func TestRegister(t *testing.T) {
	should := assert.New(t)
	newUser := user.NewLoginAndRegisterRequest()
	newUser.Username = "test001"
	newUser.Password = "222222"
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
	newUser.Password = "222222"
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
	req.Token = "xVXqrDdHbVG2uOKVE0BOnLj8"
	req.UserId = 21
	info, err := service.UserInfo(context.Background(), req)

	if should.NoError(err) {
		fmt.Println(info)
		fmt.Println(info.User.IsFollow)
	}
}

func BenchmarkUserServiceImpl_UserInfo(b *testing.B) {
	req := user.NewUserInfoRequest()
	req.Token = "kHdNO8b6zobfML4DF5WPuW7T"
	req.UserId = 16
	for i := 0; i < b.N; i++ {
		_, _ = service.UserInfo(context.Background(), req)
	}
}

func TestUserMap(t *testing.T) {
	should := assert.New(t)
	req := user.NewUserMapRequest()
	req.Token = "kHdNO8b6zobfML4DF5WPuW7T"
	req.UserIds = []int64{1, 2, 4, 16, 17, 18}
	info, err := service.UserMap(context.Background(), req)

	if should.NoError(err) {
		fmt.Println(info)
	}
}

func BenchmarkUserServiceImpl_UserMap(b *testing.B) {
	req := user.NewUserMapRequest()
	req.Token = "kHdNO8b6zobfML4DF5WPuW7T"
	req.UserIds = []int64{1, 2, 4, 16, 17, 18}
	for i := 0; i < b.N; i++ {
		_, _ = service.UserMap(context.Background(), req)
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
