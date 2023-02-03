// Author: BeYoung
// Date: 2023/2/3 18:12
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

func Message(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Comment: %v", userID)
	c := proto.NewMessageClient(models.Dials["message"])

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
	request := proto.MessageRequest{
		UserId:     userID.(int64),
		ToUserId:   toUserID,
		ActionType: int32(actionType),
		Content:    ctx.Query("content"),
	}
	if _, err = c.MessageAction(ctx, &request); err != nil {
		zap.S().Debugf("MessageAction failed: %v", err)
		ctx.JSON(http.StatusBadRequest, models.MessageResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 处理 response 响应
	ctx.JSON(http.StatusOK, models.MessageResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func MessageList(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Comment: %v", userID)
	c := proto.NewMessageClient(models.Dials["message"])

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

	// 发出请求, 获取消息列表 && 处理响应
	var response *proto.MessageListResponse
	request := proto.MessageListRequest{UserId: userID.(int64), ToUserId: toUserID}
	if response, err = c.MessageHistory(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.MessageListResponse{
			StatusCode: "1",
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 处理 response 响应
	messageResponse := models.MessageListResponse{
		StatusCode: "0",
		StatusMsg:  "success",
	}
	for i := 0; i < len(response.MessageList); i++ {
		message := models.Message{
			Content:    response.MessageList[i].Content,
			CreateTime: response.MessageList[i].CreateTime,
			ID:         response.MessageList[i].Id,
			UserID:     response.MessageList[i].UserId,
			ToUserID:   response.MessageList[i].ToUserId,
		}
		messageResponse.MessageList = append(messageResponse.MessageList, message)
	}
	ctx.JSON(http.StatusOK, messageResponse)
}
