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
	list := make([]*proto.Video, len(r))

	for i := range r {
		video := &proto.Video{
			VideoId:       r[i].ID,
			UserId:        r[i].AuthID,
			PlayUrl:       r[i].PlayURL,
			CoverUrl:      r[i].CoverURL,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         r[i].Titel,
		}
		list = append(list, video)
	}

	zap.S().Debugf("success to get comment list ==> len(comment_list): %v", len(list))
	return &proto.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  list,
	}, nil
}
