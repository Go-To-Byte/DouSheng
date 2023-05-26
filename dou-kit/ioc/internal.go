// @Author: Ciusyan 2023/1/28
package ioc

import "fmt"

// ========
// 内部依赖：InternalDependency
// ========

var (
	// 内部依赖的 IoC 容器
	internalContainer = map[string]InternalDependency{}
)

// InternalDependency 内部服务的实例想要注入此容器，必须实现该接口
type InternalDependency interface {
	// Init 如何初始化注入此 IoC 的实例
	Init() error
	// Name 注入服务模块的名称
	Name() string
}

// InternalDI :将依赖注入此容器，Internal DI（Internal Dependency Inject）
func InternalDI(dependency InternalDependency) {
	dependencyName := dependency.Name()

	// 1、检查服务是否已经被注册
	if _, ok := internalContainer[dependencyName]; ok {
		panic(fmt.Sprintf("[ %s ]服务的依赖已在容器中，请勿重复注入", dependencyName))
	}

	// 2、未注入，放入容器
	internalContainer[dependencyName] = dependency
}

// ExistingInternalDependencies 返回内部服务依赖的容器中已存在的依赖名称
func ExistingInternalDependencies() (apps []string) {
	for k, _ := range internalContainer {
		apps = append(apps, k)
	}
	return
}

// GetInternalDependency 根据模块名称 获取Gin HTTP服务模块的依赖，外部使用时需自己断言，如：
//
//	userService = ioc.GetInternalDependency("user").(user.Service)
func GetInternalDependency(name string) InternalDependency {
	if v, ok := internalContainer[name]; ok {
		return v
	} else {
		panic(fmt.Sprintf("容器中没有此依赖[ %s ]", name))
	}
}
