// Created by yczbest at 2023/02/23 11:27

package api

import (
	"github.com/cloudwego/hertz/pkg/route"
	"go.uber.org/zap"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"github.com/Go-To-Byte/DouSheng/interaction_service/client/rpc"
)

// 用于注入IOC中
var handler = &Handler{}

// Handler 通过一个实体类，把内部接口用HTTP暴露出去【控制层Controller】
type Handler struct {
	commentService comment.ServiceClient
	log            *zap.SugaredLogger

	// 提供一个空结构体，用于默认实现方法
	ioc.HertzDefault
}

// Registry 用于注册Handler所需要暴露的路由
func (h *Handler) Registry(r route.IRoutes) {
	r.POST("/action/", exception.HertzErrWrapper(h.CommentAction))
}

func (h *Handler) RegistryWithMiddle(r route.IRoutes) {
	r.GET("/list/", exception.HertzErrWrapper(h.GetCommentList))
}

// Init 初始化Handler对象
func (h *Handler) Init() error {
	h.log = zap.S().Named("COMMENT HTTP")

	// 从user_center拿到它对外提供的client，用这个Client去GRPC的调用用户中心的SDK
	client, err := rpc.NewInteractionServiceClientFromConfig()

	if err != nil {
		return err
	}
	h.commentService = client.CommentService()
	return nil
}

func (h *Handler) Name() string {
	return comment.AppName
}

func init() {
	// 将此Gin服务注入IOC中
	ioc.HertzDI(handler)
}
