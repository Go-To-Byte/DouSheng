// Package api @Author: Hexiaoming 2023/2/15
package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
)

type followListResp struct {
	*custom.CodeMsg
	*relation.FollowListResponse
}

type followerListResp struct {
	*custom.CodeMsg
	*relation.FollowerListResponse
}

type friendListResp struct {
	*custom.CodeMsg
	*relation.FriendListResponse
}

type followActionResp struct {
	*custom.CodeMsg
	*relation.FollowActionResponse
}

func (h *Handler) followList(c context.Context, ctx *app.RequestContext) error {
	req := relation.NewFollowListRequest()

	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.FollowList(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, followListResp{
		CodeMsg:            custom.Ok(constant.ACQUIRE_OK),
		FollowListResponse: resp,
	})
	return nil
}

func (h *Handler) followerList(c context.Context, ctx *app.RequestContext) error {
	req := relation.NewFollowerListRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.FollowerList(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, followerListResp{
		CodeMsg:              custom.Ok(constant.ACQUIRE_OK),
		FollowerListResponse: resp,
	})
	return nil
}

func (h *Handler) followAction(c context.Context, ctx *app.RequestContext) error {
	req := relation.NewFollowActionRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.FollowAction(c, req)
	if err != nil {
		h.log.Errorf("relation: Token：%s \n", req.Token)
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, followActionResp{
		CodeMsg:              custom.Ok(constant.ACQUIRE_OK),
		FollowActionResponse: resp,
	})
	return nil
}

// 好友列表, 与粉丝列表的差别在于多了最新消息
func (h *Handler) friendList(c context.Context, ctx *app.RequestContext) error {
	req := relation.NewFriendListRequest()
	// 1、接收参数
	if err := ctx.BindAndValidate(req); err != nil {
		h.log.Error(err)
		return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	}

	// 业务请求
	resp, err := h.service.FriendList(c, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, friendListResp{
		CodeMsg:            custom.Ok(constant.ACQUIRE_OK),
		FriendListResponse: resp,
	})
	return nil
}
