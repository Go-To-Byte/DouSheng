// @Author: Ciusyan 2023/2/9
package rpc_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
)

var (
	apiRooter *rpc.ApiRooterClient
)

func TestUserCenter(t *testing.T) {
	should := assert.New(t)

	tokenReq := token.NewValidateTokenRequest("xxx")
	// 这里主要是为了获取 用户ID
	validatedToken, err := apiRooter.TokenService().ValidateToken(context.Background(), tokenReq)
	if should.NoError(err) {
		t.Log(validatedToken)
	}

}

func init() {
	// 需要先加载配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 获取用户中心的客户端[从环境变量中获取配置]
	// 获取的配置去执行 kit 库中的 client.NewConfig(consulCfg)
	center, err := rpc.NewApiRooterClientFromEnv()
	if err != nil {
		panic(err)
	}
	apiRooter = center
}
