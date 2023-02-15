// @Author: Ciusyan 2023/1/24
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

// loginAndRegisterResp 登录和注册的响应对象
type loginAndRegisterResp struct {
	*custom.CodeMsg
	*user.TokenResponse
}

// userInfoResp 用户信息的响应对象
type userInfoResp struct {
	*custom.CodeMsg
	*user.UserInfoResponse
}

func (h *Handler) Register(c *gin.Context) error {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       custom.Ok(constant.OK_REGISTER),
			TokenResponse: resp,
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
