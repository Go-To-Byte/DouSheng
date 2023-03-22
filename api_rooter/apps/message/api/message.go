// @Author: Hexiaoming 2023/2/18
package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
)

type chatMessageListResp struct {
	*custom.CodeMsg
	*message.ChatMessageListResponse
}

type chatMessageActionResp struct {
	*custom.CodeMsg
	*message.ChatMessageActionResponse
}

func (h *Handler) chatMessageList(c context.Context, ctx *app.RequestContext) error {
	req := message.NewChatMessageListRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.ChatMessageList(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, chatMessageListResp{
		CodeMsg:                 custom.Ok(constant.ACQUIRE_OK),
		ChatMessageListResponse: resp,
	})
	return nil
}

func (h *Handler) chatMessageAction(c context.Context, ctx *app.RequestContext) error {
	req := message.NewChatMessageActionRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.ChatMessageAction(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, chatMessageActionResp{
		CodeMsg:                   custom.Ok(constant.ACQUIRE_OK),
		ChatMessageActionResponse: resp,
	})
	return nil
}
