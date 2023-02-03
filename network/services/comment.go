// Author: BeYoung
// Date: 2023/2/3 2:02
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

func Comment(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Comment: %v", userID)
	c := proto.NewCommentClient(models.Dials["comment"])

	// parse videoID
	var err error
	var videoID int64
	if videoID, err = strconv.ParseInt(ctx.Query("video_id"), 10, 64); err != nil {
		zap.S().Errorf("Error parsing videoID: %v", ctx.Query("video_id"))
		ctx.JSON(http.StatusBadRequest, models.CommentResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Error parsing videoID: %v", ctx.Query("video_id")),
		})
	}

	// Parse ActionType
	var actionType int64
	if actionType, err = strconv.ParseInt(ctx.Query("ActionType"), 10, 32); err != nil {
		zap.S().Errorf("Parse ActionType value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Parse ActionType value failed: %v", ctx.Query("to_user_id")),
		})
		ctx.Abort()
		return
	}

	// 发出请求 && 处理响应
	var err error
	var response *proto.CommentResponse
	request := proto.CommentRequest{
		UserId:     userID.(int64),
		VideoId:    videoID,
		ActionType: int32(actionType),
		Content:    ctx.Query("comment_text"),
		CommentId:  0,
	}
	if response, err = c.Comment(request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "failed to comment",
		})

	}
}

func CommentList(ctx *gin.Context) {

}
