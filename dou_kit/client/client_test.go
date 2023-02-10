// @Author: Ciusyan 2023/2/5
package client_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
)

// rpc服务通用客户端
// 需要配置注册中心的[地址、服务名称]
func TestClient(t *testing.T) {
	should := assert.New(t)

	// 配置Consul[地址、服务名称]
	consulCfg := conf.NewDefaultConsul()
	consulCfg.Host = os.Getenv("CONSUL_HOST")
	consulCfg.Port, _ = strconv.Atoi(os.Getenv("CONSUL_PORT"))
	consulCfg.RegistryName = os.Getenv("CONSUL_NAME")

	// 比如这里去发现 user_center 服务
	rpcCfg := client.NewConfig(consulCfg, "user_center")
	// 根据注册中心的配置，获取用户中心的客户端
	client, err := client.NewClientSet(rpcCfg)

	// 下面就可以使用user_center提供的SDK了
	if should.NoError(err) {
		t.Log(client)
	}
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
}
