// Author: BeYoung
// Date: 2023/2/1 15:18
// Software: GoLand

package services

import (
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Follow(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Follow: %v", userID.(int64))
	c := proto.NewRelationClient(models.Dials["relation"])

	// Parse to_user_id to int64
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("to_user_id"), 10, 64); err != nil {
		zap.S().Errorf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// Parse ActionType
	var actionType int64
	if actionType, err = strconv.ParseInt(ctx.Query("action_type"), 10, 32); err != nil {
		zap.S().Errorf("Parse ActionType value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 发出请求 && 处理响应
	request := proto.FollowRequest{
		UserId:     userID.(int64),
		ToUserId:   toUserID,
		ActionType: int32(actionType),
	}
	if _, err := c.Follow(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 处理response 响应
	ctx.JSON(http.StatusOK, models.FollowResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func FollowList(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("FollowList: %v", userID)
	c := proto.NewRelationClient(models.Dials["relation"])

	// Parse to_user_id to int64
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("user_id"), 10, 64); err != nil {
		zap.S().Errorf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 发出请求, 获取关注列表 && 处理响应
	var response *proto.FollowListResponse
	request := proto.FollowListRequest{UserId: toUserID}
	if response, err = c.FollowList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowListResponse{
			StatusCode: "1",
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 获得关注用户的id列表后，处理响应数据
	followListResponse := models.FollowListResponse{
		StatusCode: "0",
		StatusMsg:  "success",
	}
	zap.S().Debugf("relation follow list: len(user_list) = %d", len(response.UserList))

	// 依据 user_id 提取 user_info
	for i := 0; i < len(response.UserList); i++ {
		user, err := getUserInfo(request.UserId, response.UserList[i])
		if err != nil {
			continue
		}
		followListResponse.UserList = append(followListResponse.UserList, user)
	}
	ctx.JSON(http.StatusOK, followListResponse)
}

func FollowerList(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("FollowerList: %v", userID)
	c := proto.NewRelationClient(models.Dials["relation"])

	// Parse to_user_id to int64
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("user_id"), 10, 64); err != nil {
		zap.S().Errorf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 发出请求, 获取粉丝列表 && 处理响应
	var response *proto.FollowerListResponse
	request := proto.FollowerListRequest{UserId: toUserID}
	if response, err = c.FollowerList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowerListResponse{
			StatusCode: "1",
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 获得关注用户的id列表后，处理响应数据
	followerListResponse := models.FollowerListResponse{
		StatusCode: "0",
		StatusMsg:  "success",
	}
	zap.S().Debugf("relation follower list: len(user_list) = %d", len(response.UserList))

	// 依据 user_id 提取 user_info
	for i := 0; i < len(response.UserList); i++ {
		user, err := getUserInfo(request.UserId, response.UserList[i])
		if err != nil {
			continue
		}
		followerListResponse.UserList = append(followerListResponse.UserList, user)
	}
	ctx.JSON(http.StatusOK, followerListResponse)
}

func FriendList(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("FollowerList: %v", userID)
	c := proto.NewRelationClient(models.Dials["relation"])

	// Parse to_user_id to int64
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("user_id"), 10, 64); err != nil {
		zap.S().Errorf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 发出请求, 获取好友列表 && 处理响应
	var response *proto.FriendListResponse
	request := proto.FriendListRequest{UserId: toUserID}
	if response, err = c.FriendList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FriendListResponse{
			StatusCode: "1",
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 获得好友的id列表后，处理响应数据
	friendListResponse := models.FriendListResponse{
		StatusCode: "0",
		StatusMsg:  "success",
	}
	zap.S().Debugf("relation follower list: len(user_list) = %d", len(response.UserList))

	// 依据 user_id 提取 user_info
	for i := 0; i < len(response.UserList); i++ {
		user, err := getUserInfo(request.UserId, response.UserList[i])
		if err != nil {
			continue
		}
		friendListResponse.UserList = append(friendListResponse.UserList, user)
	}
	ctx.JSON(http.StatusOK, friendListResponse)
}
