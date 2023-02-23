// Created by yczbest at 2023/02/23 10:33

package api

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/gin-gonic/gin"
	"net/http"
)

// loginAndRegisterResp 登录和注册的响应对象
type favoriteActionResponse struct {
	*custom.CodeMsg
}

// userInfoResp 用户信息的响应对象
type getFavoriteListResponse struct {
	*custom.CodeMsg
	*favorite.GetFavoriteListResponse
}

func (h *Handler) FavoriteAction(c *gin.Context) error {

	req := favorite.NewFavoriteActionRequest()
	// 1、接收参数
	if err := c.Bind(req); err != nil {
		fmt.Printf("参数为：%s", req.VideoId)
		fmt.Printf("参数为：%s", req.Token)
		fmt.Printf("参数为：%s", req.ActionType)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	_, err := h.favoriteService.FavoriteAction(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		favoriteActionResponse{
			CodeMsg: custom.Ok(constant.OK_REGISTER),
		})
	return nil
}

func (h *Handler) GetFavoriteList(c *gin.Context) error {

	req := favorite.NewDefaultGetFavoriteListRequest()

	// 1、接收参数
	if err := c.Bind(req); err != nil {
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.favoriteService.GetFavoriteList(c.Request.Context(), req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	c.JSON(http.StatusOK,
		getFavoriteListResponse{
			CodeMsg:                 custom.NewWithCode(constant.OPERATE_OK),
			GetFavoriteListResponse: resp,
		})
	return nil
}
