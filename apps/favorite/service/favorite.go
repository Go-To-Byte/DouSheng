// Author: BeYoung
// Date: 2023/1/30 23:28
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/dao"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/favorite/proto"
)

func (f *Favorite) Favorite(ctx context.Context, req *proto.FavoriteRequest) (*proto.FavoriteResponse, error) {
	favorite := model.Favorite{
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

	return &proto.FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
	}, nil
}

func (f *Favorite) FavoriteList(ctx context.Context, req *proto.FavoriteListRequest) (*proto.FavoriteListResponse, error) {
	r := dao.FavoriteFindByUserID(req.UserId)
	list := make([]int64, len(r))
	for i := range r {
		list = append(list, r[i].VideoID)
	}
	return &proto.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		VideoList:  list,
	}, nil
}
