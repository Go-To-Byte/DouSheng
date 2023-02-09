// @Author: Ciusyan 2023/2/9
package rpc_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc"
)

var (
	userCenter *rpc.UserCenterClient
)

func TestUserCenter(t *testing.T) {
	should := assert.New(t)

	tokenReq := token.NewValidateTokenRequest("xxx")
	// 这里主要是为了获取 用户ID
	validatedToken, err := userCenter.TokenService().ValidateToken(context.Background(), tokenReq)
	if should.NoError(err) {
		t.Log(validatedToken)
	}

}

func TestUserCenter_GinAuthHandlerFunc(t *testing.T) {
	r := gin.New()
	// 使用 auth 中间件
	group := r.Group("/v1", userCenter.GinAuthHandlerFunc())

	group.GET("/", func(c *gin.Context) {
		c.String(200, "Get")
	})
	r.Run()
}

func init() {
	// 需要先加载配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 获取用户中心的客户端[从环境变量中获取配置]
	// 获取的配置去执行 kit 库中的 client.NewConfig(consulCfg, discoverName)
	center, err := rpc.NewUserCenterClientFromEnv()
	if err != nil {
		panic(err)
	}
	userCenter = center
}
