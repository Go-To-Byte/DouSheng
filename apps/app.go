// @Author: Ciusyan 2023/1/25
package apps

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/gin-gonic/gin"
)

// IOC容器：管理所有服务的实例

var (
	UserService user.Service
	// 维护IOC中的所有服务
	implApps = map[string]ImplService{}
	// 维护IOC中的所有Gin服务
	ginApps = map[string]GinService{}
)

// DIServiceImpl Dependence Inject Service Impl：将依赖注入此IOC容器中
func DIServiceImpl(service ImplService) {
	serviceName := service.Name()
	// 1、检查是否已经注入依赖
	if _, ok := implApps[serviceName]; ok {
		panic(fmt.Sprintf("ImplService %s has Injected", serviceName))
	}

	// 2、注入依赖
	implApps[serviceName] = service

	// 3、需要在IOC中注册具体的Service
	if v, ok := service.(user.Service); ok {
		UserService = v
	}
}

// InitService 初始化所有已经注入IOC的服务
func InitService() {
	for _, v := range implApps {
		v.Config()
	}
}

// ImplService 想要自动注入IOC的服务，必须实现的接口
type ImplService interface {
	Config()
	Name() string
}

// GetServiceImpl 从IOC中获取ServiceImpl实例，外界使用时需要自己断言
func GetServiceImpl(name string) interface{} {
	if v, ok := implApps[name]; ok {
		return v
	}
	return nil
}

func DIGinService(service GinService) {
	serviceName := service.Name()

	// 1、判断是否已经注入过依赖了
	if _, ok := ginApps[serviceName]; ok {
		panic(fmt.Sprintf("Gin Service %s has Injected", serviceName))
	}

	// 2、将其注入IOC
	ginApps[serviceName] = service
}

// InitGin 初始化Gin服务的配置以及路由
func InitGin(r gin.IRouter) {
	// 1、初始化所有Gin服务的对象
	for _, v := range ginApps {
		v.Config()
	}

	// 2、初始化所有对象的路由
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
