// @Author: Ciusyan 2023/1/24
package http

import (
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/gin-gonic/gin"
)

func NewUserHttpHandler(service user.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Handler 通过一个实体类，把内部接口用HTTP暴露出去【控制层Controller】
type Handler struct {
	service user.Service
}

// Registry 用于注册Handler所需要暴露的路由
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/douyin/user/register/")
}
