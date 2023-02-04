// @Author: Ciusyan 2023/1/24
package api

import (
	"Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		// TODO：字符串转字符串指针的小工具
		msg := err.Error()
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}

	// 2、进行接口调用
	resp, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		msg := err.Error()
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}
	resp.StatusCode = 0

	c.JSON(http.StatusBadRequest, resp)
}

func (h *Handler) Login(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		// TODO：字符串转字符串指针的小工具
		msg := err.Error()
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		msg := err.Error()
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
	}
	resp.StatusCode = 0

	c.JSON(http.StatusBadRequest, resp)
}

func (h *Handler) GetUserInfo(c *gin.Context) {

}
