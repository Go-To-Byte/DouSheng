// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/video/apps/video"
	"github.com/Go-To-Byte/DouSheng/video/apps/video/impl/dal/model"

	_ "github.com/Go-To-Byte/DouSheng/video/apps/video/impl/init"

	"go.uber.org/zap"
	"time"
)

func (c *VideoServiceImpl) Publish(ctx context.Context, req *video.PublishRequest) (*video.PublishResponse, error) {
	// 从 request 获取评论信息
	v := model.Video{
		ID:       req.VideoId,
		AuthID:   req.UserId,
		Titel:    req.Title,
		CoverURL: req.CoverUrl,
		PlayURL:  req.VideoUrl,
	}
	// 添加评论
	if err := c.Add(v); err != nil {
		zap.S().Errorf("failed to add video: %+v", v)

		return &video.PublishResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		}, err
	}

	zap.S().Debugf("success to add video: %+v", v)
	return &video.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (c *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	r := c.VideoFindByUserID(req.UserId)
	list := make([]int64, 0)

	for i := range r {
		list = append(list, r[i].ID)
	}

	zap.S().Debugf("success to get comment list ==> len(comment_list): %v", len(list))
	return &video.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  list,
	}, nil
}

func (c *VideoServiceImpl) Info(ctx context.Context, req *video.VideoInfoRequest) (*video.VideoInfoResponse, error) {
	r := c.VideoFindByVideoID(req.VideoId)
	list := make([]*video.Video, 0)

	for i := range r {
		v := &video.Video{
			Id:            r[i].ID,
			Author:        r[i].AuthID,
			PlayUrl:       r[i].PlayURL,
			CoverUrl:      r[i].CoverURL,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         r[i].Titel,
		}
		list = append(list, v)
	}

	zap.S().Debugf("success to get video info %+v", list[0])
	return &video.VideoInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Video:      list[0],
	}, nil
}

func (c *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	var timeStamp = (time.Now().UnixMicro()/1000 - 1288834974657) << 22
	zap.S().Debugf("feed time stamp %+v", timeStamp)
	if req.LatestTime != nil {
		timeStamp = (*req.LatestTime - 1288834974657) << 22
	}
	r := c.VideoFindByTimeStamp(timeStamp)
	list := make([]int64, 0)

	for i := range r {
		list = append(list, r[i].ID)
	}

	nextTime := time.Now().UnixMicro() / 1000
	if len(list) > 0 {
		nextTime = r[len(list)-1].ID
	}
	zap.S().Debugf("success to get video list ==> len(video): %v", len(list))
	return &video.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  list,
		NextTime:   nextTime,
	}, nil
}
