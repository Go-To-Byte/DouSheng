// Created by yczbest at 2023/02/23 10:33

package api

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/interaction_service/client/rpc"
	"github.com/gin-gonic/gin"
)

// 用于注入IOC中
var handler = &Handler{}

// Handler 通过一个实体类，把内部接口用HTTP暴露出去【控制层Controller】
type Handler struct {
	favoriteService favorite.ServiceClient
	// 提供一个空结构体，用于默认实现方法
	ioc.GinDefault
}

func (h *Handler) RegistryWithMiddle(r gin.IRoutes) {
	r.POST("/action/", exception.GinErrWrapper(h.FavoriteAction))
	r.GET("/list/", exception.GinErrWrapper(h.GetFavoriteList))
}

// Init 初始化Handler对象
func (h *Handler) Init() error {
	// 从user_center拿到它对外提供的client，用这个Client去GRPC的调用用户中心的SDK
	client, err := rpc.NewInteractionServiceClientFromConfig()

	if err != nil {
		return err
	}
	h.favoriteService = client.FavoriteService()
	return nil
}

func (h *Handler) Name() string {
	return favorite.AppName
}

func init() {
	// 将此Gin服务注入IOC中
	ioc.GinDI(handler)
}
