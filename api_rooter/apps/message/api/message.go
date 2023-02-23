// @Author: Hexiaoming 2023/2/18
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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


func (h *Handler) chatMessageList(ctx *gin.Context) error {
	req := message.NewChatMessageListRequest()
	// 1、接收参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }
	
	// TODO 需替换为绑定参数
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	req.ToUserId = toUserId
	req.Token = ctx.Query("token")

	// 业务请求
	resp, err := h.service.ChatMessageList(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, chatMessageListResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		ChatMessageListResponse: resp,
	})
	return nil
}

func (h *Handler) chatMessageAction(ctx *gin.Context) error {
	req := message.NewChatMessageActionRequest()
	// 1、接收参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }
	
	// TODO 需替换为绑定参数
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 32)
	req.ToUserId = toUserId
	req.Token = ctx.Query("token")
	req.ActionType = int32(actionType)
	req.Content = ctx.Query("content")

	// req.Token = "Tjs37Alvfx8jobp9epWtwT5X"
	// req.ToUserId = 11
	// req.ActionType = 1
	// req.Content = "hello 123"

	// 业务请求
	resp, err := h.service.ChatMessageAction(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, chatMessageActionResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		ChatMessageActionResponse: resp,
	})
	return nil
}
