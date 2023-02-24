// @Author: Hexiaoming 2023/2/17
package api

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/relation_service/client/rpc"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
)

type Handler struct {
	service relation.ServiceClient
	l       logger.Logger

	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

func (h *Handler) Init() error {
	h.l = zap.L().Named(relation.AppName)

	client, err := rpc.NewRelationServiceClientFromCfg()
	if err != nil {
		return err
	}

	h.service = client.RelationService()
	return nil
}

func (h *Handler) Name() string {
	return relation.AppName
}

func (h *Handler) Version() string {
	return ""
}

func (h *Handler) Registry(r gin.IRoutes) {

}

func (h *Handler) RegistryWithMiddle(r gin.IRoutes) {
	r.GET("/follow/list/", exception.GinErrWrapper(h.followList))
	r.GET("/follower/list/", exception.GinErrWrapper(h.followerList))
	// TODO 待处理
	// r.POST("/action/", exception.GinErrWrapper(h.followAction))
	r.POST("/action/", exception.GinErrWrapper(h.followAction))
	r.GET("/friend/list/", exception.GinErrWrapper(h.friendList))
}

func init() {
	ioc.GinDI(&Handler{})
}
