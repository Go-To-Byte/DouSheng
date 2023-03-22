// @Author: Ciusyan 2023/3/14

package ioc

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
)

// ========
// 内部依赖：HertzDependency
// ========

var (
	// Hertz依赖的 IoC 容器
	hertzContainer = map[string]HertzDependency{}
)

// HertzDependency Hertz HTTP的服务实例想要注入此容器，必须实现该接口
type HertzDependency interface {
	InternalDependency
	// Registry 该模块所需要注册的路由：添加传递的中间件
	// [因为 hertz不支持路由装饰，只能使用路由分组的方式区分特殊的路由(如：添加中间件)]
	Registry(r route.IRoutes)
	// RegistryWithMiddle 该用于注册特殊的路由：不添加传递的中间件
	RegistryWithMiddle(r route.IRoutes)
	// Version 该服务API的版本
	Version() string
}

// HertzDI :将依赖注入此容器，Hertz DI（Hertz Dependency Inject）
func HertzDI(dependency HertzDependency) {
	dependencyName := dependency.Name()

	// 1、检查服务是否已经被注册
	if _, ok := hertzContainer[dependencyName]; ok {
		panic(fmt.Sprintf("[ %s ]服务的依赖已在容器中，请勿重复注入", dependencyName))
	}

	// 2、未注入，放入容器
	hertzContainer[dependencyName] = dependency
}

// ExistingHertzDependencies 返回Hertz HTTP服务依赖的容器中已存在的依赖名称
func ExistingHertzDependencies() (apps []string) {
	for k := range hertzContainer {
		apps = append(apps, k)
	}
	return
}

// GetHertzDependency 根据模块名称 获取内部服务模块的依赖，外部使用时需自己断言，如：
//
//	userHertz = ioc.GetHertzDependency("user").(user.UserHertz)
func GetHertzDependency(name string) HertzDependency {
	if v, ok := hertzContainer[name]; ok {
		return v
	} else {
		panic(fmt.Sprintf("容器中没有此依赖[ %s ]", name))
	}
}

// RegistryHertz 注册所有的Hertz Http 服务
func RegistryHertz(opts *HertzOptions) {

	if opts.Router == nil {
		log.Fatal("请传递路由对象")
		return
	}

	prefix := opts.Prefix
	// 防止API前缀没有 "/"
	if prefix != "" && !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}

	for _, v := range hertzContainer {
		vers := v.Version()
		name := v.Name()

		// 处理特殊路由
		switch AppName(name) {
		case USER_API:
		case VIDEO_API:
			vers = ""
			name = ""
		case RELATION_API:
		default:
		}

		group := opts.Router.Group(path.Join(prefix, vers, name))

		v.Registry(group)
		// 使用中间件[如：大部分路由都需要开启认证]
		v.RegistryWithMiddle(group.Use(opts.Middleware...))
	}
}

// HertzOptions 注册IOC中Hertz服务的路由时，可传入配置
type HertzOptions struct {
	ApiOptions
	// 路由
	Router *server.Hertz `json:"router"`
	// API 中间件
	Middleware []app.HandlerFunc `json:"middleware"`

	// ...
}

// NewHertzOption ：路由对象、前缀、中间件
func NewHertzOption(r *server.Hertz, prefix string, middle ...app.HandlerFunc) *HertzOptions {
	return &HertzOptions{
		Router:     r,
		Middleware: middle,
		ApiOptions: ApiOptions{
			Prefix: prefix,
		},
	}
}

// ---------------分割线-----------------
// HertzDefault 用于默认实现HertzIoc 的所有方法

var (
	// 用于验证是否满足 HertzDependency 接口
	_ HertzDependency = (*HertzDefault)(nil)
)

// HertzDefault 类似与GRPC的做法，提供一个默认实现的结构体，方便外界使用
type HertzDefault struct{}

func (g HertzDefault) Init() error {
	return fmt.Errorf("请实现方法")
}

func (g HertzDefault) Name() string {
	log.Fatal("请实现方法[Name()]")
	return ""
}

func (g HertzDefault) Registry(r route.IRoutes) {

}

func (g HertzDefault) RegistryWithMiddle(r route.IRoutes) {

}

func (g HertzDefault) Version() string {
	return ""
}
