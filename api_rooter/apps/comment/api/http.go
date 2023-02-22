// Package api Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand
package api

import (
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment"
	"github.com/gin-gonic/gin"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc"
)

// 用于注入IOC中
var handler = &Handler{}

// Handler 通过一个实体类，把内部接口用HTTP暴露出去【控制层Controller】
type Handler struct {
	service comment.CommentClient

	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

// Registry 用于注册Handler所需要暴露的路由
func (h *Handler) Registry(r gin.IRoutes) {
	r.POST("/register/", exception.GinErrWrapper(h.Register))
	r.POST("/login/", exception.GinErrWrapper(h.Login))
}

func (h *Handler) RegistryWithMiddle(r gin.IRoutes) {
	r.GET("/", exception.GinErrWrapper(h.GetUserInfo))
}

// Init 初始化Handler对象
func (h *Handler) Init() error {
	// 从user_center拿到它对外提供的client，用这个Client去GRPC的调用用户中心的SDK
	client, err := rpc.NewUserCenterClientFromCfg()

	if err != nil {
		return err
	}
	h.service = client.UserService()
	return nil
}

func (h *Handler) Name() string {
	return user.AppName
}

func init() {
	// 将此Gin服务注入IOC中
	ioc.GinDI(handler)
}
