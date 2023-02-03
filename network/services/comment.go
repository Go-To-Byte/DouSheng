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
		ctx.Abort()
		return
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

	// Parse commentID
	var commentID int64
	if commentID, err = strconv.ParseInt(ctx.Query("ActionType"), 10, 32); err != nil {
		zap.S().Errorf("Parse ActionType value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FollowResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Parse ActionType value failed: %v", ctx.Query("to_user_id")),
		})
		ctx.Abort()
		return
	}

	var response *proto.CommentResponse
	request := proto.CommentRequest{
		UserId:     userID.(int64),
		VideoId:    videoID,
		ActionType: int32(actionType),
		Content:    ctx.Query("comment_text"),
		CommentId:  commentID,
	}

	// 发出请求 && 处理响应
	if response, err = c.Comment(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "failed to comment",
		})
		ctx.Abort()
		return
	}

	user, _ := getUserInfo(userID.(int64), userID.(int64))
	ctx.JSON(http.StatusOK, models.CommentResponse{
		Comment: models.Comment{
			Content:    response.Comment.Content,
			CreateDate: response.Comment.CreateDate,
			ID:         response.Comment.Id,
			User:       user,
		},
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func CommentList(ctx *gin.Context) {
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
		ctx.Abort()
		return
	}

	// 请求评论列表
	var response *proto.CommentListResponse
	request := proto.CommentListRequest{VideoId: videoID}
	if response, err = c.CommentList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CommentListResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 处理 response 结构
	commentList := models.CommentListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}

	// 填充 response
	for i := 0; i < len(response.CommentList); i++ {
		user, err := getUserInfo(userID.(int64), response.CommentList[i].User)
		if err != nil {
			continue
		}
		comment := models.Comment{
			Content:    response.CommentList[i].Content,
			CreateDate: response.CommentList[i].CreateDate,
			ID:         response.CommentList[i].Id,
			User:       user,
		}
		commentList.Comments = append(commentList.Comments, comment)
	}

	ctx.JSON(http.StatusOK, commentList)
}
