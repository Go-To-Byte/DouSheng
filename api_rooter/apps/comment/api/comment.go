// Created by yczbest at 2023/02/23 10:33

package api

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) CommentAction(c *gin.Context) error {

	req := comment.NewCommentActionRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.commentService.CommentAction(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		commentActionResponse{
			CodeMsg:               custom.Ok(constant.OK_REGISTER),
			CommentActionResponse: resp,
		})
	return nil
}

func (h *Handler) GetCommentList(c *gin.Context) error {

	req := comment.NewDefaultGetCommentListRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.commentService.GetCommentList(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		getCommentListResponse{
			CodeMsg:                custom.NewWithCode(constant.OPERATE_OK),
			GetCommentListResponse: resp,
		})
	return nil
}
