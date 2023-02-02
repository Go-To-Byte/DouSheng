// Author: BeYoung
// Date: 2023/1/30 0:09
// Software: GoLand

package init

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/message/models"
	"github.com/hashicorp/consul/api"
)

// 注册服务
func initConsul() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.75.141:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.10.5:%v", models.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = "user"
	registration.ID = models.Node.Generate().String()
	registration.Port = models.Port
	registration.Tags = []string{"user"}
	registration.Address = "192.168.10.5"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	// client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}
	return
}
