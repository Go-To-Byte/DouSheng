// Created by yczbest at 2023/02/23 10:33

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
)

type commentActionResponse struct {
	*custom.CodeMsg
	*comment.CommentActionResponse
}

// userInfoResp 用户信息的响应对象
type getCommentListResponse struct {
	*custom.CodeMsg
	*comment.GetCommentListResponse
}

func (h *Handler) CommentAction(c context.Context, ctx *app.RequestContext) error {

	req := comment.NewCommentActionRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.commentService.CommentAction(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK,
		commentActionResponse{
			CodeMsg:               custom.NewWithCode(constant.OPERATE_OK),
			CommentActionResponse: resp,
		})
	return nil
}

func (h *Handler) GetCommentList(c context.Context, ctx *app.RequestContext) error {

	req := comment.NewDefaultGetCommentListRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}
	// 2、进行接口调用
	resp, err := h.commentService.GetCommentList(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK,
		getCommentListResponse{
			CodeMsg:                custom.NewWithCode(constant.OPERATE_OK),
			GetCommentListResponse: resp,
		})
	return nil
}
