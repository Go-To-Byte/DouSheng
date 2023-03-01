// Package impl @Author: Ciusyan 2023/2/28
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

// get follow list count, user += followListCount, user += followerListCount
func (s *userServiceImpl) composeRelation(ctx context.Context, uResp *user.User) (errs []error) {

	// 携程内部必须捕获 panic()
	defer func() {
		if r := recover(); r != nil {
			s.l.Errorf("user composeRelation：panic %v", r)
			return
		}
	}()

	relationCountReq := relation.NewListCountRequest()
	relationCountReq.UserId = uResp.Id
	relationCountReq.Type = relation.CountType_ALL
	relationCount, err := s.relation.ListCount(ctx, relationCountReq)
	if err != nil {
		s.l.Error(err)
		errs = append(errs, err)
	}
	uResp.FollowCount = &relationCount.FollowCount
	uResp.FollowerCount = &relationCount.FollowerCount
	// TODO：user += isFollow （relation 添加 isFollow 方法）

	return
}

// get publish list, user += publishCount
func (s *userServiceImpl) composeVideo(ctx context.Context, uResp *user.User) (errs []error) {

	// 携程内部必须捕获 panic()
	defer func() {
		if r := recover(); r != nil {
			s.l.Errorf("user composeRelation：panic %v", r)
			return
		}
	}()

	publishCountReq := video.NewPublishListCountRequest(uResp.Id)
	publishCount, err := s.video.PublishListCount(ctx, publishCountReq)
	if err != nil {
		s.l.Error(err)
		errs = append(errs, err)
	}

	uResp.WorkCount = &publishCount.PublishCount

	return
}

// get favorite list, user += favoriteCount
func (s *userServiceImpl) composeFavorite(ctx context.Context, uResp *user.User) (errs []error) {

	// 携程内部必须捕获 panic()
	defer func() {
		if r := recover(); r != nil {
			s.l.Errorf("user composeRelation：panic %v", r)
			return
		}
	}()

	favoriteListReq := favorite.NewFavoritePo()
	favoriteListReq.UserId = uResp.Id
	favoriteCountResp, err := s.favorite.FavoriteCount(ctx, favoriteListReq)
	if err != nil {
		s.l.Error(err)
		errs = append(errs, err)
	}

	uResp.FavoriteCount = &favoriteCountResp.FavoriteCount

	// TODO：获赞数量

	return
}
