// Author: BeYoung
// Date: 2023/2/3 15:21
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

func Favorite(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Favorite: %v", userID)
	c := proto.NewFavoriteClient(models.Dials["favorite"])

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

	// 处理grpc请求数据结构
	request := proto.FavoriteRequest{
		UserId:     userID.(int64),
		VideoId:    videoID,
		ActionType: int32(actionType),
	}

	// 发送grpc请求
	if _, err = c.Favorite(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FavoriteResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 处理 response
	ctx.JSON(http.StatusOK, models.FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func FavoriteList(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Favorite: %v", userID)
	c := proto.NewFavoriteClient(models.Dials["favorite"])

	// parse toUserID
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("user_id"), 10, 63); err != nil {
		zap.S().Errorf("Parse toUserID value failed(id: %v): %v", ctx.Query("to_user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.FavoriteListResponse{
			StatusCode: "1",
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	// 处理grpc请求
	var response *proto.FavoriteListResponse
	request := proto.FavoriteListRequest{UserId: toUserID}
	if response, err = c.FavoriteList(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.FavoriteListResponse{
			StatusCode: "1",
			StatusMsg:  "failed",
		})
	}

	// 处理 response
	favoriteList := models.FavoriteListResponse{
		StatusCode: "0",
		StatusMsg:  "success",
	}
	for i := 0; i < len(response.VideoList); i++ {
		video, err := getVideoInfo(userID.(int64), response.VideoList[i])
		if err != nil {
			continue
		}
		favoriteList.VideoList = append(favoriteList.VideoList, video)
	}
	ctx.JSON(http.StatusOK, favoriteList)
}
