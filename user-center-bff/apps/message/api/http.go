// Package api @Author: Hexiaoming 2023/2/18
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/message"
	"github.com/Go-To-Byte/DouSheng/interaction-service/client/rpc"
)

type Handler struct {
	service message.ServiceClient
	l       logger.Logger

	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

func (h *Handler) Init() error {
	h.l = zap.L().Named(message.AppName)

	client, err := rpc.NewInteractionServiceClientFromConfig()
	if err != nil {
		return err
	}

	h.service = client.MessageService()
	return nil
}

func (h *Handler) Name() string {
	return message.AppName
}

func (h *Handler) Version() string {
	return ""
}

func (h *Handler) RegistryWithMiddle(r gin.IRoutes) {
	r.GET("/chat/", exception.GinErrWrapper(h.chatMessageList))
	r.POST("/action/", exception.GinErrWrapper(h.chatMessageAction))
}

func init() {
	ioc.GinDI(&Handler{})
}
