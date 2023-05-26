// @Author: Hexiaoming 2023/2/18
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/message-service/apps/message"
)

type chatMessageListResp struct {
	*custom.CodeMsg
	*message.ChatMessageListResponse
}

type chatMessageActionResp struct {
	*custom.CodeMsg
	*message.ChatMessageActionResponse
}

func (h *Handler) chatMessageList(ctx *gin.Context) error {
	req := message.NewChatMessageListRequest()
	// 1、接收参数
	if err := ctx.ShouldBindQuery(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.ChatMessageList(ctx, req)
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

func (h *Handler) chatMessageAction(ctx *gin.Context) error {
	req := message.NewChatMessageActionRequest()
	// 1、接收参数
	if err := ctx.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.ChatMessageAction(ctx, req)
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
