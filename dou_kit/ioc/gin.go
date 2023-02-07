// @Author: Ciusyan 2023/1/28
package ioc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

// ========
// 内部依赖：GinDependency
// ========

var (
	// Gin依赖的 IoC 容器
	ginContainer = map[string]GinDependency{}
)

// GinDependency Gin HTTP的服务实例想要注入此容器，必须实现该接口
type GinDependency interface {
	InternalDependency
	// Registry 该模块所需要注册的路由
	Registry(r gin.IRouter)
	// Version 该服务API的版本
	Version() string
}

// GinDI :将依赖注入此容器，Gin DI（Gin Dependency Inject）
func GinDI(dependency GinDependency) {
	dependencyName := dependency.Name()

	// 1、检查服务是否已经被注册
	if _, ok := ginContainer[dependencyName]; ok {
		panic(fmt.Sprintf("[ %s ]服务的依赖已在容器中，请勿重复注入", dependencyName))
	}

	// 2、未注入，放入容器
	ginContainer[dependencyName] = dependency
}

// ExistingGinDependencies 返回Gin HTTP服务依赖的容器中已存在的依赖名称
func ExistingGinDependencies() (apps []string) {
	for k, _ := range ginContainer {
		apps = append(apps, k)
	}
	return
}

// GetGinDependency 根据模块名称 获取内部服务模块的依赖，外部使用时需自己断言，如：
//
//	userGin = ioc.GetGinDependency("user").(user.UserGin)
func GetGinDependency(name string) GinDependency {
	if v, ok := ginContainer[name]; ok {
		return v
	} else {
		panic(fmt.Sprintf("容器中没有此依赖[ %s ]", name))
	}
}

// RegistryGin 注册所有的Gin Http 服务
func RegistryGin(prefix string, r gin.IRouter) {
	// 防止API前缀没有 "/"
	if prefix != "" && !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	for _, v := range ginContainer {
		v.Registry(r.Group(path.Join(prefix, v.Version(), v.Name())))
	}
}
