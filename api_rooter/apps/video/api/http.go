// @Author: Ciusyan 2023/2/7
package api

import (
	"github.com/cloudwego/hertz/pkg/route"
	"go.uber.org/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/client/rpc"
)

type Handler struct {
	service video.ServiceClient
	log     *zap.SugaredLogger

	// 提供一个空结构体，用于默认实现方法
	ioc.HertzDefault
}

func (h *Handler) Init() error {
	h.log = zap.S().Named("VIDEO HTTP")

	client, err := rpc.NewVideoServiceClientFromCfg()
	if err != nil {
		return err
	}

	h.service = client.VideoService()
	return nil
}

func (h *Handler) Name() string {
	return video.AppName
}

func (h *Handler) Version() string {
	return ""
}

func (h *Handler) Registry(r route.IRoutes) {
	r.GET("/feed/", exception.HertzErrWrapper(h.feed))
}

func (h *Handler) RegistryWithMiddle(r route.IRoutes) {
	r.POST("/publish/action/", exception.HertzErrWrapper(h.publishAction))
	r.GET("/publish/list/", exception.HertzErrWrapper(h.publishList))
}

func init() {
	ioc.HertzDI(&Handler{})
}
