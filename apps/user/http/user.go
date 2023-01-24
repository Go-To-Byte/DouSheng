// @Author: Ciusyan 2023/1/24
package http

import (
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

func (h *Handler) RegisterUser(c *gin.Context) {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	// 2、进行接口调用git
	token, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, token)
}
