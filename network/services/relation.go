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

	// TODO: JWT Authorization
	token := ctx.Query("token")
	if userID, err := strconv.ParseInt(token, 10, 64); err == nil {
		zap.S().Panicf("Invalid token value failed(token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Invalid token value failed: %v", token),
		})
		ctx.Abort()
	} else {
		// Parse to_user_id to int64
		if toUserID, err := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64); err != nil {
			zap.S().Panicf("Parse to_user_id value failed(id: %v): %v", ctx.Query("to_user_id"), err)
			ctx.JSON(http.StatusBadRequest, models.FollowResponse{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("Parse to_user_id value failed: %v", ctx.Query("to_user_id")),
			})
			ctx.Abort()
		} else {
			request := proto.FollowRequest{
				UserId:     userID,
				ToUserId:   toUserID,
				ActionType: 0,
			}
			if response, err := c.Follow(ctx, &request); err != nil {
				ctx.JSON(http.StatusOK, models.FollowResponse{
					StatusCode: 0,
					StatusMsg:  response.StatusMsg,
				})
			}
		}
	}
}

func Friend(ctx *gin.Context) {

}

func FollowList(ctx *gin.Context) {

}

func FollowerList(ctx *gin.Context) {

}

func FriendList(ctx *gin.Context) {

}
