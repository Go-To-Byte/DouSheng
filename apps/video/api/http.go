// @Author: Ciusyan 2023/1/29
package api

import (
	"github.com/Go-To-Byte/DouSheng/apps/video"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/gin-gonic/gin"
)

var (
	handler = &Handler{}
)

type Handler struct {
	service video.ServiceServer
}

func (h *Handler) Init() error {
	h.service = ioc.GetGrpcDependency(video.AppName).(video.ServiceServer)
	return nil
}

func (h *Handler) Name() string {
	return video.AppName
}

func (h *Handler) Version() string {
	return ""
}

// Registry 此模块提供的HTTP路由
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("/feed/", h.Feed)
	r.POST("/publish/action/", h.PublishAction)
	r.GET("/publish/list/", h.PublishList)
}

func init() {
	// 注入IOC容器
	ioc.GinDI(handler)
}
