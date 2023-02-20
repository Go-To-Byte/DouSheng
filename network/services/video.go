// Author: BeYoung
// Date: 2023/2/3 20:53
// Software: GoLand

package services

import (
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func Publish(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("userID: %v", userID)
	c := proto.NewVideoClient(models.Dials["video"])

	VideoId, exists := ctx.Get("video_id")
	if exists == false {
		zap.S().Errorf("failed get video_id")
	}
	VideoUrl, exists := ctx.Get("video_url")
	if exists == false {
		zap.S().Errorf("failed get video_url")
	}
	CoverUrl, exists := ctx.Get("cover_url")
	if exists == false {
		zap.S().Errorf("failed get cover_url")
	}

	// 发起 grpc 请求
	request := proto.PublishRequest{
		UserId:   userID.(int64),
		VideoId:  VideoId.(int64),
		VideoUrl: VideoUrl.(string),
		CoverUrl: CoverUrl.(string),
		Title:    ctx.Query("title"),
	}

	if _, err := c.Publish(ctx, &request); err != nil {
		zap.S().Errorf("failed to publish: %v", err)
		ctx.JSON(http.StatusBadRequest, models.PublishResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
	}

	ctx.JSON(http.StatusOK, models.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func PublishList(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("userID: %v", userID)
	c := proto.NewVideoClient(models.Dials["video"])

	// Parse userID
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("user_id"), 10, 64); err != nil {
		zap.S().Errorf("Parse user_id value failed(id: %v): %v", ctx.Query("user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.PublishListResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	var response *proto.PublishListResponse
	request := proto.PublishListRequest{UserId: toUserID}
	if response, err = c.PublishList(ctx, &request); err != nil {
		zap.S().Errorf("failed to publish list: %v", err)
		ctx.JSON(http.StatusBadRequest, models.PublishListResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
	}

	publishList := models.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	for i := 0; i < len(response.VideoList); i++ {
		video, err := getVideoInfo(userID.(int64), response.VideoList[i])
		if err != nil {
			continue
		}
		publishList.VideoList = append(publishList.VideoList, video)
	}
	ctx.JSON(http.StatusOK, publishList)
}

func Feed(ctx *gin.Context) {
	token := ctx.Query("token")
	zap.S().Debugf("user: %v", token)
	c := proto.NewVideoClient(models.Dials["video"])

	// Parse timeStamp
	var err error
	var timeStamp int64
	if timeStamp, err = strconv.ParseInt(ctx.Query("latest_time"), 10, 64); err != nil {
		zap.S().Errorf("Parse latest_time value failed(latest_time: %v): %v", ctx.Query("latest_time"), err)
		timeStamp = time.Now().UnixMicro() / 1000
	}

	var response *proto.FeedResponse
	request := proto.FeedRequest{LatestTime: &timeStamp}
	if response, err = c.Feed(ctx, &request); err != nil {
		zap.S().Errorf("failed feed client: %+v", &request)
		ctx.JSON(http.StatusBadRequest, models.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
		ctx.Abort()
		return
	}

	feed := models.FeedResponse{
		NextTime:   response.NextTime,
		StatusCode: 0,
		StatusMsg:  "success:",
	}
	for i := 0; i < len(response.VideoList); i++ {
		video, _ := getVideoInfo(0, response.VideoList[i])
		if err != nil {
			continue
		}
		feed.VideoList = append(feed.VideoList, video)
	}
	ctx.JSON(http.StatusOK, feed)
}
