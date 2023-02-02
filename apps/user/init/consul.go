// Author: BeYoung
// Date: 2023/1/30 0:09
// Software: GoLand

package init

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user/models"
	"github.com/hashicorp/consul/api"
)

// 注册服务
func initConsul() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%v:%v", models.Config.Consul.Host, models.Config.Consul.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%v:%v", models.Config.Localhost.Host, models.Config.Localhost.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = models.Config.Consul.Name
	registration.ID = models.Node.Generate().String()
	registration.Port = models.Config.Localhost.Port
	registration.Tags = []string{models.Config.Consul.Tags}
	registration.Address = models.Config.Localhost.Host
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	// client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}
	return
}
