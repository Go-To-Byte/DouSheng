// Author: BeYoung
// Date: 2023/1/30 23:28
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite/impl/models"
)

func (f *FavoriteServiceImpl) Favorite(ctx context.Context, req *favorite.FavoriteRequest) (*favorite.FavoriteResponse, error) {
	if req.ActionType == 1 {
		fav := model.Favorite{
			ID:      models.Node.Generate().Int64(),
			UserID:  req.UserId,
			VideoID: req.VideoId,
			Flag:    1,
		}

		if err := f.Add(fav); err != nil {
			return &favorite.FavoriteResponse{
				StatusCode: 1,
				StatusMsg:  "favorite add failed",
			}, err
		}
	} else {
		fav := model.Favorite{
			ID:      0,
			UserID:  req.UserId,
			VideoID: req.VideoId,
			Flag:    1,
		}

		if err := f.Delete(fav); err != nil {
			return &favorite.FavoriteResponse{
				StatusCode: 1,
				StatusMsg:  "favorite delete failed",
			}, err
		}
	}

	return &favorite.FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
	}, nil
}

func (f *FavoriteServiceImpl) FavoredList(ctx context.Context, req *favorite.FavoredListRequest) (*favorite.FavoredListResponse, error) {
	r := f.FavoriteFindByVideoID(model.Favorite{VideoID: req.VideoId})
	list := make([]int64, 0)
	for i := range r {
		list = append(list, r[i].UserID)
	}
	return &favorite.FavoredListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil
}

func (f *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	r := f.FavoriteFindByUserID(model.Favorite{UserID: req.UserId})
	list := make([]int64, 0)
	for i := range r {
		list = append(list, r[i].VideoID)
	}
	return &favorite.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		VideoList:  list,
	}, nil
}

func (f *FavoriteServiceImpl) FavoriteJudge(ctx context.Context, req *favorite.FavoriteJudgeRequest) (*favorite.FavoriteJudgeResponse, error) {
	r := f.FavoriteFindByUserIDAndVideoID(model.Favorite{UserID: req.UserId, VideoID: req.VideoId})
	list := make([]int64, 0)
	for i := range r {
		list = append(list, r[i].VideoID)
	}
	return &favorite.FavoriteJudgeResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		IsFavorite: int32(len(list)),
	}, nil
}
