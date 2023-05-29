// Created by yczbest at 2023/02/23 10:33

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/favorite"
)

// loginAndRegisterResp 登录和注册的响应对象
type favoriteActionResponse struct {
	*custom.CodeMsg
}

// userInfoResp 用户信息的响应对象
type getFavoriteListResponse struct {
	*custom.CodeMsg
	*favorite.FavoriteListResponse
}

func (h *Handler) FavoriteAction(c *gin.Context) error {

	req := favorite.NewFavoriteActionRequest()
	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	_, err := h.favoriteService.FavoriteAction(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		favoriteActionResponse{
			CodeMsg: custom.NewWithCode(constant.OPERATE_OK),
		})
	return nil
}

func (h *Handler) GetFavoriteList(c *gin.Context) error {

	req := favorite.NewFavoriteListRequest()

	// 1、接收参数
	if err := c.ShouldBind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.favoriteService.FavoriteList(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		getFavoriteListResponse{
			CodeMsg:              custom.NewWithCode(constant.OPERATE_OK),
			FavoriteListResponse: resp,
		})
	return nil
}
