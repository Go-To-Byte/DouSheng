// Author: BeYoung
// Date: 2023/2/1 15:18
// Software: GoLand

package services

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Follow(ctx *gin.Context) {
	zap.S().Debugf("Follow")
	c := proto.NewRelationClient(models.GrpcConn)
	request := proto.FollowRequest{
		UserId:     0,
		ToUserId:   0,
		ActionType: 0,
	}

	// TODO: JWT Authorization
	var err error
	token := ctx.Query("token")
	if request.UserId, err = strconv.ParseInt(token, 10, 64); err == nil {
		zap.S().Panicf("Invalid token value failed(token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Invalid token value failed: %v", token),
		})
		ctx.Abort()
		return
	}
	// Parse to_user_id to int64
	if request.ToUserId, err = strconv.ParseInt(ctx.Query("to_user_id"), 10, 64); err != nil {
		zap.S().Panicf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Parse to_user_id value failed: %v", ctx.Query("to_user_id")),
		})
		ctx.Abort()
		return
	}
	// Parse ActionType to int32
	ActionType, err := strconv.ParseInt(ctx.Query("ActionType"), 10, 32)
	if request.ActionType = int32(ActionType); err != nil {
		zap.S().Panicf("Parse ActionType value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Parse ActionType value failed: %v", ctx.Query("to_user_id")),
		})
		ctx.Abort()
		return
	}

	// 发出请求 && 处理响应
	if response, err := c.Follow(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  response.StatusMsg,
		})
	} else {
		ctx.JSON(http.StatusOK, models.FollowResponse{
			StatusCode: 0,
			StatusMsg:  response.StatusMsg,
		})
	}

}

func FollowList(ctx *gin.Context) {
	zap.S().Debugf("Follow")
	c := proto.NewRelationClient(models.GrpcConn)
	request := proto.FollowListRequest{UserId: 0}

	// TODO: JWT Authorization
	var err error
	token := ctx.Query("token")
	if request.UserId, err = strconv.ParseInt(token, 10, 64); err == nil {
		zap.S().Panicf("Invalid token value failed(token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Invalid token value failed: %v", token),
		})
		ctx.Abort()
		return
	}

	// 发出请求 && 处理响应
	if response, err := c.FollowList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FollowListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
			UserList:   nil,
		})
	} else {
		// 依据 user_id 提前 user_info
		responseJson := models.FollowListResponse{
			StatusCode: string(response.StatusCode),
			StatusMsg:  response.StatusMsg,
		}

		for i := 0; i < len(response.UserList); i++ {

		}
		ctx.JSON(http.StatusOK, responseJson)
	}
}

func FollowerList(ctx *gin.Context) {

}

func FriendList(ctx *gin.Context) {

}
