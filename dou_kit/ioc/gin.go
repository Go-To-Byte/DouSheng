// @Author: Ciusyan 2023/1/28
package ioc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	// Registry 该模块所需要注册的路由：添加传递的中间件
	// [因为gin不支持路由装饰，只能使用路由分组的方式区分特殊的路由(如：添加中间件)]
	Registry(r gin.IRoutes)
	// RegistryWithMiddle 该用于注册特殊的路由：不添加传递的中间件
	RegistryWithMiddle(r gin.IRoutes)
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
func RegistryGin(opts *GinOptions) {

	if opts.Router == nil {
		log.Fatal("请传递路由对象")
		return
	}

	prefix := opts.Prefix
	// 防止API前缀没有 "/"
	if prefix != "" && !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}

	for _, v := range ginContainer {
		vers := v.Version()
		name := v.Name()

		// 不传就是需要 Version
		if opts.NotVersion {
			vers = ""
		}
		// 不传就是需要 Name
		if opts.NotName {
			name = ""
		}

		group := opts.Router.Group(path.Join(prefix, vers, name))
		v.Registry(group)
		// 使用中间件[如：大部分路由都需要开启认证]
		v.RegistryWithMiddle(group.Use(opts.Middleware...))

	}
}

// GinOptions 注册IOC中Gin服务的路由时，可传入配置
type GinOptions struct {
	// 路由
	Router gin.IRouter `json:"router"`
	// API 前缀
	Prefix string `json:"prefix"`
	// API 中间件
	Middleware []gin.HandlerFunc `json:"middleware"`
	// API 是否需要添加版本
	NotVersion bool `json:"not_version"`
	// API 是否需要添加服务名称
	NotName bool `json:"not_name"`

	// ...
}

// NewGinOption ：路由对象、前缀、中间件
func NewGinOption(r gin.IRouter, prefix string, middle ...gin.HandlerFunc) *GinOptions {
	return &GinOptions{
		Router:     r,
		Prefix:     prefix,
		Middleware: middle,
	}
}

// ---------------分割线-----------------
// GinDefault 用于默认实现GinIoc 的所有方法

var (
	// 用于验证是否满足 GinDependency 接口
	_ GinDependency = (*GinDefault)(nil)
)

// GinDefault 类似与GRPC的做法，提供一个默认实现的结构体，方便外界使用
type GinDefault struct{}

func (g GinDefault) Init() error {
	return fmt.Errorf("请实现方法")
}

func (g GinDefault) Name() string {
	log.Fatal("请实现方法[Name()]")
	return ""
}

func (g GinDefault) Registry(r gin.IRoutes) {

}

func (g GinDefault) RegistryWithMiddle(r gin.IRoutes) {

}

func (g GinDefault) Version() string {
	return ""
}
