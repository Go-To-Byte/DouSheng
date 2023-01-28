// @Author: Ciusyan 2023/1/28
package ioc

import (
	"fmt"
	"google.golang.org/grpc"
)

// ========
// 内部依赖：GrpcDependency
// ========

var (
	// Grpc依赖的 IoC 容器
	grpcContainer = map[string]GrpcDependency{}
)

// GrpcDependency GRPC 服务的实例想要注入此容器，必须实现该接口
type GrpcDependency interface {
	InternalDependency
	// Registry 该模块所需要注册的GRPC服务
	Registry(s *grpc.Server)
}

// GrpcDI :将依赖注入此容器，Grpc DI（Grpc Dependency Inject）
func GrpcDI(dependency GrpcDependency) {
	dependencyName := dependency.Name()

	// 1、检查服务是否已经被注册
	if _, ok := grpcContainer[dependencyName]; ok {
		panic(fmt.Sprintf("[ %s ]服务的依赖已在容器中，请勿重复注入", dependencyName))
	}

	// 2、未注入，放入容器
	grpcContainer[dependencyName] = dependency
}

// ExistingGrpcDependencies 返回GRPC服务依赖的容器中已存在的依赖名称
func ExistingGrpcDependencies() (apps []string) {
	for k, _ := range grpcContainer {
		apps = append(apps, k)
	}
	return
}

// GetGrpcDependency 根据模块名称 获取GRPC服务模块的依赖，外部使用时需自己断言，如：
//
//	userGrpc = ioc.GetGrpcDependency("user").(user.UserGrpc)
func GetGrpcDependency(name string) GrpcDependency {
	if v, ok := grpcContainer[name]; ok {
		return v
	} else {
		panic(fmt.Sprintf("容器中没有此依赖[ %s ]", name))
	}
}

// RegistryGrpc 注册所有的Grpc服务
func RegistryGrpc(s *grpc.Server) {
	for _, v := range grpcContainer {
		v.Registry(s)
	}
}
