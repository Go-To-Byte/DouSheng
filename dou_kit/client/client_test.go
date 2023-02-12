// @Author: Ciusyan 2023/2/5
package client_test

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/client"
)

// rpc服务通用客户端
// 需要配置注册中心的[地址、服务名称]
func TestClient(t *testing.T) {
	should := assert.New(t)

	// 配置Consul[地址、服务名称]
	cfg := conf.NewDefaultDiscover()
	cfg.SetAddr(os.Getenv("DISCOVER_ADDRESS"))
	cfg.SetDiscoverName("DISCOVER_NAME")

	// 比如这里去发现 user_center 服务
	// 根据注册中心的配置，获取用户中心的客户端
	client, err := client.NewClientSet(cfg)

	// 下面就可以使用user_center提供的SDK了
	if should.NoError(err) {
		t.Log(client)
	}
}

func init() {
	conf.LoadConfigFromEnv()
}
