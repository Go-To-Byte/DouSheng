// @Author: Ciusyan 2023/1/24
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	userconst "github.com/Go-To-Byte/DouSheng/user_center/common/constant"
)

// loginAndRegisterResp 登录和注册的响应对象
type loginAndRegisterResp struct {
	constant.CodeMsg
	user.TokenResponse
}

// userInfoResp 用户信息的响应对象
type userInfoResp struct {
	constant.CodeMsg
	user.UserInfoResponse
}

func (h *Handler) Register(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, constant.ERROR_ARGS_VALIDATE)
		return
	}

	// 2、进行接口调用
	resp, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		e := err.(*exception.Exception)
		c.JSON(http.StatusBadRequest, e.GetCodeMsg())
		return
	}

	c.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       *userconst.OK_REGISTER,
			TokenResponse: *resp.Clone(),
		})
}

func (h *Handler) Login(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, constant.ERROR_ARGS_VALIDATE)
		return
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		e := err.(*exception.Exception)
		c.JSON(http.StatusBadRequest, e.GetCodeMsg())
		return
	}

	c.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       *constant.OPERATE_OK,
			TokenResponse: *resp.Clone(),
		})
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	req := user.NewUserInfoRequest()

	// 1、接收参数
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, constant.ERROR_ARGS_VALIDATE)
		return
	}

	info, err := h.service.UserInfo(c.Request.Context(), req)
	if err != nil {
		e := err.(*exception.Exception)
		c.JSON(http.StatusBadRequest, e.GetCodeMsg())
		return
	}

	c.JSON(http.StatusOK,
		userInfoResp{
			CodeMsg:          *constant.OPERATE_OK,
			UserInfoResponse: *info.Clone(),
		})
}
