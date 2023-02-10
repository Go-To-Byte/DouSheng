// @Author: Ciusyan 2023/1/24
package api

import (
	"github.com/gin-gonic/gin"

	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

// 用于注入IOC中
var handler = &Handler{}

func NewUserHttpHandler() *Handler {
	return &Handler{}
}

// Handler 通过一个实体类，把内部接口用HTTP暴露出去【控制层Controller】
type Handler struct {
	service user.ServiceServer

	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

// Registry 用于注册Handler所需要暴露的路由
func (h *Handler) Registry(r gin.IRoutes) {
	r.POST("/register/", h.Register)
	r.POST("/login/", h.Login)
	r.GET("/", h.GetUserInfo)
}

// Init 初始化Handler对象
func (h *Handler) Init() error {
	// 从IOC中获取UserServiceImpl实例
	h.service = ioc.GetGrpcDependency(user.AppName).(user.ServiceServer)
	return nil
}

func (h *Handler) Name() string {
	return user.AppName
}

func init() {
	// 将此Gin服务注入IOC中
	ioc.GinDI(handler)
}
