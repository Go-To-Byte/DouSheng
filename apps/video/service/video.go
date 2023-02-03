// Author: BeYoung
// Date: 2023/1/30 23:28
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/video/dao"
	"github.com/Go-To-Byte/DouSheng/apps/video/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/video/models"
	"github.com/Go-To-Byte/DouSheng/apps/video/proto"
	"go.uber.org/zap"
	"time"
)

func (v *Video) Publish(ctx context.Context, req *proto.PublishRequest) (*proto.PublishResponse, error) {
	// 从 request 获取评论信息
	video := model.Video{
		ID:       models.Node.Generate().Int64(),
		AuthID:   req.UserId,
		Titel:    req.Title,
		CoverURL: req.CoverUrl,
		PlayURL:  req.VideoUrl,
	}
	// 添加评论
	if err := dao.Add(video); err != nil {
		zap.S().Errorf("failed to add video: %+v", video)

		return &proto.PublishResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
		}, err
	}

	zap.S().Debugf("success to add video: %+v", video)
	return &proto.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (v *Video) PublishList(ctx context.Context, req *proto.PublishListRequest) (*proto.PublishListResponse, error) {
	r := dao.VideoFindByUserID(req.UserId)
	list := make([]int64, 0)

	for i := range r {
		list = append(list, r[i].ID)
	}

	zap.S().Debugf("success to get comment list ==> len(comment_list): %v", len(list))
	return &proto.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  list,
	}, nil
}

func (v *Video) Info(ctx context.Context, req *proto.VideoInfoRequest) (*proto.VideoInfoResponse, error) {
	r := dao.VideoFindByVideoID(req.VideoId)
	list := make([]*proto.Video, 0)

	for i := range r {
		video := &proto.Video{
			Id:            r[i].ID,
			Author:        r[i].AuthID,
			PlayUrl:       r[i].PlayURL,
			CoverUrl:      r[i].CoverURL,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         r[i].Titel,
		}
		list = append(list, video)
	}

	zap.S().Debugf("success to get video info %+v", list[0])
	return &proto.VideoInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Video:      list[0],
	}, nil
}

func (v *Video) Feed(ctx context.Context, req *proto.FeedRequest) (*proto.FeedResponse, error) {
	var timeStamp = (time.Now().UnixNano() - 1288834974657) << 22
	if req.LatestTime != nil {
		timeStamp = *req.LatestTime
	}
	r := dao.VideoFindByTimeStamp(timeStamp)
	list := make([]int64, 0)

	for i := range r {
		list = append(list, r[i].ID)
	}

	zap.S().Debugf("success to get video list ==> len(comment_list): %v", len(list))
	return &proto.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  list,
		NextTime:   r[len(r)-1].ID,
	}, nil
}
