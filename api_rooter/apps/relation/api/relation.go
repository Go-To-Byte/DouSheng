// Package api @Author: Hexiaoming 2023/2/15
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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

func (h *Handler) followList(ctx *gin.Context) error {
	req := relation.NewFollowListRequest()

	// 1、接收参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	// TODO 需替换为绑定参数
	userId, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	req.UserId = userId
	req.Token = ctx.Query("token")

	// 业务请求
	resp, err := h.service.FollowList(ctx, req)
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

func (h *Handler) followerList(ctx *gin.Context) error {
	req := relation.NewFollowerListRequest()
	// 1、接收参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	// TODO 需替换为绑定参数
	userId, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	req.UserId = userId
	req.Token = ctx.Query("token")

	// 业务请求
	resp, err := h.service.FollowerList(ctx, req)
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

func (h *Handler) followAction(ctx *gin.Context) error {
	req := relation.NewFollowActionRequest()
	// 1、接收参数
	// if err := ctx.Bind(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	// TODO 需替换为绑定参数
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 32)
	req.ToUserId = toUserId
	req.Token = ctx.Query("token")
	req.ActionType = int32(actionType)

	// 业务请求
	resp, err := h.service.FollowAction(ctx, req)
	if err != nil {
		h.l.Errorf("relation: Token：%s \n", req.Token)
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
func (h *Handler) friendList(ctx *gin.Context) error {
	req := relation.NewFriendListRequest()
	// 1、接收参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	// TODO 需替换为绑定参数
	userId, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	req.UserId = userId
	req.Token = ctx.Query("token")

	// 业务请求
	resp, err := h.service.FriendList(ctx, req)
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
