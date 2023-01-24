// @Author: Ciusyan 2023/1/25
package apps

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user"
)

// IOC容器：管理所有服务的实例

var (
	UserService user.Service
	// 维护IOC中的所有服务
	services = map[string]Service{}
)

// DependenceInject DI：将依赖注入此IOC容器中
func DependenceInject(service Service) {
	serviceName := service.Name()
	// 1、检查是否已经注入依赖
	if _, ok := services[serviceName]; ok {
		panic(fmt.Sprintf("Service %s has Injected", serviceName))
	}

	// 2、注入依赖
	services[serviceName] = service

	// 3、需要在IOC中注册具体的Service
	if v, ok := service.(user.Service); ok {
		UserService = v
	}
}

// Init 初始化所有已经注入IOC的服务
func Init() {
	for _, v := range services {
		v.Config()
	}
}

// Service 想要自动注入IOC的服务，必须实现的接口
type Service interface {
	Config()
	Name() string
}
