// Package impl @Author: Ciusyan 2023/3/3
package impl

import (
	"context"
	kitUtils "github.com/Go-To-Byte/DouSheng/dou_kit/utils"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

// 组合用户信息
func (s *videoServiceImpl) composeUser(ctx context.Context, pos []*video.VideoPo, tk string) (map[int64]*user.User, error) {
	// 携程内部必须捕获 panic()
	defer func() {
		if r := recover(); r != nil {
			s.l.Errorf("user composeRelation：panic %v", r)
			return
		}
	}()

	// 1、取出视频列表的 userIds
	userIds := kitUtils.NewSet()

	for _, po := range pos {
		userIds.Add(po.AuthorId)
	}

	// 2、获取UserMap列表
	userMapReq := user.NewUserMapRequest()
	userMapReq.UserIds = userIds.Items()
	userMapReq.Token = tk

	// GRPC调用
	userMap, err := s.userServer.UserMap(ctx, userMapReq)
	if err != nil {
		return nil, err
	}

	return userMap.GetUserMap(), nil
}

// 组合点赞信息
func (s *videoServiceImpl) composeFavorite(ctx context.Context, videoIds []int64, loginUid int64) (
	map[int64]*favorite.FavoriteMap, error) {

	// 携程内部必须捕获 panic()
	defer func() {
		if r := recover(); r != nil {
			s.l.Errorf("user composeRelation：panic %v", r)
			return
		}
	}()

	req := favorite.NewFavoriteMapRequest()
	req.VideoIds = videoIds
	req.UserId = loginUid

	// GRPC调用
	favoriteMap, err := s.favoriteService.FavoriteCountMap(ctx, req)
	if err != nil {
		return nil, err
	}

	return favoriteMap.GetFavoriteMap(), nil
}

// 组合评论信息
func (s *videoServiceImpl) composeComment(ctx context.Context, videoIds []int64) (map[int64]int64, error) {
	// 携程内部必须捕获 panic()
	defer func() {
		if r := recover(); r != nil {
			s.l.Errorf("user composeRelation：panic %v", r)
			return
		}
	}()
	req := comment.NewCommentMapRequest()
	req.VideoIds = videoIds

	// GRPC调用
	countMap, err := s.commentService.CommentCountMap(ctx, req)
	if err != nil {
		return nil, err
	}

	return countMap.GetCommentCountMap(), nil
}
