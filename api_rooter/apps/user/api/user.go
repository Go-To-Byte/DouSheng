// @Author: Ciusyan 2023/1/24

package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
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

func (h *Handler) Register(c context.Context, ctx *app.RequestContext) error {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.service.Register(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       custom.Ok(constant.OK_REGISTER),
			TokenResponse: resp,
		})
	return nil
}

func (h *Handler) Login(c context.Context, ctx *app.RequestContext) error {

	req := user.NewLoginAndRegisterRequest()

	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 2、进行接口调用
	resp, err := h.service.Login(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK,
		loginAndRegisterResp{
			CodeMsg:       custom.NewWithCode(constant.OPERATE_OK),
			TokenResponse: resp,
		})
	return nil
}

func (h *Handler) GetUserInfo(c context.Context, ctx *app.RequestContext) error {
	req := user.NewUserInfoRequest()

	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		fmt.Println(err, "====================================")
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	info, err := h.service.UserInfo(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	ctx.JSON(http.StatusOK,
		userInfoResp{
			CodeMsg:          custom.Ok(constant.ACQUIRE_OK),
			UserInfoResponse: info,
		})
	return nil
}
