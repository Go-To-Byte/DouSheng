// @Author: Ciusyan 2023/1/23
package impl_test

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service user.ServiceServer
)

func TestCreateUser(t *testing.T) {
	should := assert.New(t)
	newUser := user.NewLoginAndRegisterRequest()
	newUser.Username = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	newUser.Password = "xxxx"
	token, err := service.Register(context.Background(), newUser)

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
	zap.DevelopmentSetup()

	// 从IOC中获取接口实现
	service = ioc.GetInternalDependency(user.AppName).(user.ServiceServer)
}
