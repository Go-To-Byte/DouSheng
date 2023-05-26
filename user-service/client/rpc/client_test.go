// @Author: Ciusyan 2023/2/9
package rpc_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
	"github.com/Go-To-Byte/DouSheng/user-service/client/rpc"
)

var (
	userService *rpc.UserServiceClient
)

func TestUserService(t *testing.T) {
	should := assert.New(t)

	req := user.NewUserInfoRequest()
	// 这里主要是为了获取 用户ID
	validatedToken, err := userService.UserService().UserInfo(context.Background(), req)
	if should.NoError(err) {
		t.Log(validatedToken)
	}
}

// TODO：编写关系服务SDK测试代码
func TestRelationService(t *testing.T) {

}

func init() {
	// 需要先加载配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 获取用户中心的客户端[从环境变量中获取配置]
	// 获取的配置去执行 kit 库中的 client.NewConfig(consulCfg)
	service, err := rpc.NewUserCenterClientFromEnv()
	if err != nil {
		panic(err)
	}

	userService = service
}
