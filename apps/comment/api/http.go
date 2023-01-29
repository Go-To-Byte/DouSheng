// @Author: Ciusyan 2023/1/29
package api

import (
	"github.com/Go-To-Byte/DouSheng/apps/comment"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service comment.ServiceServer
}

func (h *Handler) Init() error {
	h.service = ioc.GetGrpcDependency(comment.AppName).(comment.ServiceServer)
	return nil
}

func (h *Handler) Name() string {
	return comment.AppName
}

func (h *Handler) Version() string {
	return ""
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/action/", h.CommentAction)
	r.GET("/list/", h.CommentList)
}

func init() {
	ioc.GinDI(&Handler{})
}
