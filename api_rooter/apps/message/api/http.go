// Package api @Author: Hexiaoming 2023/2/18
package api

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/message_service/client/rpc"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
)

type Handler struct {
	service message.ServiceClient
	l       logger.Logger

	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

func (h *Handler) Init() error {
	h.l = zap.L().Named(message.AppName)

	client, err := rpc.NewMessageServiceClientFromCfg()
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

func (h *Handler) Registry(r gin.IRoutes) {
	r.GET("/chat", exception.GinErrWrapper(h.chatMessageList))
}

func (h *Handler) RegistryWithMiddle(r gin.IRoutes) {
	r.GET("/action", exception.GinErrWrapper(h.chatMessageAction))
}

func init() {
	ioc.GinDI(&Handler{})
}
