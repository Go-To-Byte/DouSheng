// @Author: Hexiaoming 2023/2/15
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
	// "github.com/Go-To-Byte/DouSheng/relation_service/common/utils"
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
	// TODO 临时测试请求, 需替换为绑定参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	req.UserId = 10
	req.Token = "VzVifO2phBxKJsvCNygmbGQO"

	// 业务请求
	resp, err := h.service.FollowList(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, followListResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		FollowListResponse: resp,
	})
	return nil
}

func (h *Handler) followerList(ctx *gin.Context) error {
	req := relation.NewFollowerListRequest()
	// 1、接收参数
	// TODO 临时测试请求, 需替换为绑定参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	req.UserId = 10
	req.Token = "VzVifO2phBxKJsvCNygmbGQO"

	// 业务请求
	resp, err := h.service.FollowerList(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, followerListResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		FollowerListResponse: resp,
	})
	return nil
}

func (h *Handler) followAction(ctx *gin.Context) error {
	req := relation.NewFollowActionRequest()
	// 1、接收参数
	// TODO 临时测试请求, 需替换为绑定参数
	// if err := ctx.Bind(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	req.ToUserId = 12
	req.Token = "VzVifO2phBxKJsvCNygmbGQO"
	req.ActionType = "1"

	// 业务请求
	resp, err := h.service.FollowAction(ctx, req)
	if err != nil {
		h.l.Errorf("relation: Token：%s \n", req.Token)
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, followActionResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		FollowActionResponse: resp,
	})
	return nil
}

// 好友列表, 与粉丝列表的差别在于多了最新消息
func (h *Handler) friendList(ctx *gin.Context) error {
	req := relation.NewFriendListRequest()
	// 1、接收参数
	// TODO 临时测试请求, 需替换为绑定参数
	// if err := ctx.ShouldBindQuery(req); err != nil {
	// 	return exception.WithStatusCode(constant.ERROR_ARGS_VALIDATE)
	// }

	req.UserId = 10
	req.Token = "0N7Ser1RrITQO92mz0eka7El"

	// 业务请求
	resp, err := h.service.FriendList(ctx, req)
	if err != nil {
		return exception.GrpcErrWrapper(err)
	}

	// 获取成功
	ctx.JSON(http.StatusOK, friendListResp{
		CodeMsg:             custom.Ok(constant.ACQUIRE_OK),
		FriendListResponse: resp,
	})
	return nil
}