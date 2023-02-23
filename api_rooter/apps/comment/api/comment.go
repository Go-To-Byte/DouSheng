// @Author: Ciusyan 2023/1/24
package api

import (
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

type CommentResponse struct {
	*custom.CodeMsg
	Comment Comment `json:"comment"`
}

type Comment struct {
	*user.UserInfoResponse
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
	ID         int64  `json:"id"`
}

type CommentListResponse struct {
	*custom.CodeMsg
	Comments []Comment `json:"comment_list"`
}

func (h *Handler) Comment(c *gin.Context) error {
	req := comment.CommentRequest{}

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.service.Comment(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	cmt := Comment{}
	cmt.ID = resp.Comment.User
	cmt.Content = resp.Comment.Content
	cmt.UserInfoResponse

	c.JSON(http.StatusOK,
		CommentResponse{
			CodeMsg: custom.Ok(constant.OK_REGISTER),
			Comment: Comment{},
		})
	return nil
}

func (h *Handler) Login(c *gin.Context) error {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       custom.NewWithCode(constant.OPERATE_OK),
			TokenResponse: resp,
		})
	return nil
}

func (h *Handler) GetUserInfo(c *gin.Context) error {
	req := user.NewUserInfoRequest()

	// 1、接收参数
	if err := c.ShouldBindQuery(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	info, err := h.service.UserInfo(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		userInfoResp{
			CodeMsg:          custom.Ok(constant.ACQUIRE_OK),
			UserInfoResponse: info,
		})
	return nil
}
