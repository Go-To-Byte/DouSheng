// Package api @Author: Hexiaoming 2023/2/18
package api

import (
	"github.com/cloudwego/hertz/pkg/route"
	"go.uber.org/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
	"github.com/Go-To-Byte/DouSheng/message_service/client/rpc"
)

type Handler struct {
	service message.ServiceClient
	log     *zap.SugaredLogger

	// 提供一个空结构体，用于默认实现方法
	ioc.HertzDefault
}

func (h *Handler) Init() error {
	h.log = zap.S().Named("MESSAGE HTTP")

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

func (h *Handler) RegistryWithMiddle(r route.IRoutes) {
	r.GET("/chat/", exception.HertzErrWrapper(h.chatMessageList))
	r.POST("/action/", exception.HertzErrWrapper(h.chatMessageAction))
}

func init() {
	ioc.HertzDI(&Handler{})
}
