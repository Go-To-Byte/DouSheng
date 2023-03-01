// Package impl @Author: Ciusyan 2023/2/28
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
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
		return append(errs, err)
	}

	uResp.FollowCount = &relationCount.FollowCount
	uResp.FollowerCount = &relationCount.FollowerCount

	loginUId, err := s.getLoginUID(ctx)
	if err != nil {
		s.l.Errorf("user composeRelation：根据Token获取ID失败，%s", err.Error())
		return append(errs, err)
	}

	// 是否关注查询用户
	req := relation.NewDefaultUserFollowerPo()
	req.UserId = uResp.Id // 查看此用户登录用户是否是查询用户的粉丝
	req.FollowerId = loginUId
	isFollow, err := s.relation.IsFollower(ctx, req)
	if err != nil {
		s.l.Error(err)
		return append(errs, err)
	}

	uResp.IsFollow = isFollow.MyFollower

	return
}

func (s *userServiceImpl) getLoginUID(ctx context.Context) (int64, error) {
	// user += isFollow （relation 添加 isFollow 方法）

	loginUserId := ctx.Value(constant.USER_ID)
	if loginUserId == nil {
		// 从Token中获取
		tkReq := token.NewValidateTokenRequest(ctx.Value(constant.REQUEST_TOKEN).(string))
		uIdResp, err := s.tokenService.GetUIDFromTk(ctx, tkReq)
		loginUserId = uIdResp.UserId
		if err != nil {
			return 0, err
		}
	}

	return loginUserId.(int64), nil
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
	publishCount, err := s.video.ComposeVideoCount(ctx, publishCountReq)
	if err != nil {
		s.l.Error(err)
		errs = append(errs, err)
	}

	uResp.WorkCount = &publishCount.PublishCount
	uResp.TotalFavorited = &publishCount.AcquireTotalFavorite

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

	favoriteListReq := favorite.NewFavoriteCountRequest()
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
