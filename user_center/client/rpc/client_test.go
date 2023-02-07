// @Author: Ciusyan 2023/2/5
package rpc_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc"
	"github.com/Go-To-Byte/DouSheng/user_center/conf"
)

// user_center 客户端
// 需要配置注册中心的[地址、服务名称]
// 利用注册中心 获取user_center的客户端
func TestToken(t *testing.T) {
	should := assert.New(t)

	// 配置Consul[地址、服务名称]
	consulCfg := conf.NewDefaultConsul()
	consulCfg.Host = os.Getenv("CONSUL_HOST")
	consulCfg.Port, _ = strconv.Atoi(os.Getenv("CONSUL_PORT"))
	consulCfg.RegistryName = os.Getenv("CONSUL_NAME")

	// 去发现 user_center 服务
	rpcCfg := rpc.NewConfig(consulCfg, "user_center")
	// 根据注册中心的配置，获取用户中心的客户端
	client, err := rpc.NewClientSet(rpcCfg)

	// 下面就可以使用user_center提供的SDK了
	if should.NoError(err) {
		req := user.NewLoginAndRegisterRequest()
		req.Username = "ciusyan"
		req.Password = "111"

		serviceClient := client.Token()

		request := token.NewValidateTokenRequest("jYG8ai8OmXxvoUu2cSBfSmwe")

		resp, err := serviceClient.ValidateToken(context.Background(), request)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp.GetUserId())
	}
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
}
