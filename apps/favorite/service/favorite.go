// Author: BeYoung
// Date: 2023/1/30 23:28
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/dao"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/models"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/proto"
)

func (f *Favorite) Favorite(ctx context.Context, req *proto.FavoriteRequest) (*proto.FavoriteResponse, error) {
	if req.ActionType == 1 {
		favorite := model.Favorite{
			ID:      models.Node.Generate().Int64(),
			UserID:  req.UserId,
			VideoID: req.VideoId,
			Flag:    1,
		}

		if err := dao.Add(favorite); err != nil {
			return &proto.FavoriteResponse{
				StatusCode: 1,
				StatusMsg:  "favorite add failed",
			}, err
		}
	} else {
		favorite := model.Favorite{
			ID:      0,
			UserID:  req.UserId,
			VideoID: req.VideoId,
			Flag:    1,
		}

		if err := dao.Delete(favorite); err != nil {
			return &proto.FavoriteResponse{
				StatusCode: 1,
				StatusMsg:  "favorite delete failed",
			}, err
		}
	}

	return &proto.FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
	}, nil
}

func (f *Favorite) FavoredList(ctx context.Context, req *proto.FavoredListRequest) (*proto.FavoredListResponse, error) {
	r := dao.FavoriteFindByVideoID(model.Favorite{VideoID: req.VideoId})
	list := make([]int64, 0)
	for i := range r {
		list = append(list, r[i].UserID)
	}
	return &proto.FavoredListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil
}

func (f *Favorite) FavoriteList(ctx context.Context, req *proto.FavoriteListRequest) (*proto.FavoriteListResponse, error) {
	r := dao.FavoriteFindByUserID(model.Favorite{UserID: req.UserId})
	list := make([]int64, 0)
	for i := range r {
		list = append(list, r[i].VideoID)
	}
	return &proto.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		VideoList:  list,
	}, nil
}

func (f *Favorite) FavoriteJudge(ctx context.Context, req *proto.FavoriteJudgeRequest) (*proto.FavoriteJudgeResponse, error) {
	r := dao.FavoriteFindByUserIDWithVideoID(model.Favorite{UserID: req.UserId, VideoID: req.VideoId})
	list := make([]int64, 0)
	for i := range r {
		list = append(list, r[i].VideoID)
	}
	return &proto.FavoriteJudgeResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		IsFavorite: int32(len(list)),
	}, nil
}
