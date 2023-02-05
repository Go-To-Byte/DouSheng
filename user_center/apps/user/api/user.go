// @Author: Ciusyan 2023/1/24
package api

import (
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {

		msg := err.Error()
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
		return
	}

	// 2、进行接口调用
	resp, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		msg := err.Error()
		c.JSON(http.StatusBadRequest, user.TokenResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
		return
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserInfo(c *gin.Context) {

}
