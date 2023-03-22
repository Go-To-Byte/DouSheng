// Package api @Author: Hexiaoming 2023/2/17
package api

import (
	"github.com/cloudwego/hertz/pkg/route"
	"go.uber.org/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
	"github.com/Go-To-Byte/DouSheng/relation_service/client/rpc"
)

type Handler struct {
	service relation.ServiceClient
	log     *zap.SugaredLogger

	// 提供一个空结构体，用于默认实现方法
	ioc.HertzDefault
}

func (h *Handler) Init() error {
	h.log = zap.S().Named("RELATION HTTP")

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

func (h *Handler) Registry(r route.IRoutes) {

}

func (h *Handler) RegistryWithMiddle(r route.IRoutes) {
	r.GET("/follow/list/", exception.HertzErrWrapper(h.followList))
	r.GET("/follower/list/", exception.HertzErrWrapper(h.followerList))
	r.POST("/action/", exception.HertzErrWrapper(h.followAction))
	r.GET("/friend/list/", exception.HertzErrWrapper(h.friendList))
}

func init() {
	ioc.HertzDI(&Handler{})
}
