// Created by yczbest at 2023/02/18 15:02

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// 点赞数据插入 成功返回po，失败返回错误信息
func (f *favoriteServiceImpl) InsertFavoriteRecord(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoritePo, error) {
	//构造新的实例
	po, err := f.NewFavoritePo(ctx, req)
	if err != nil {
		return nil, err
	}
	//检查是否已经存在此纪录
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := f.tokenService.ValidateToken(ctx, tokenReq)
	po = favorite.NewFavoritePo()
	db := f.db.WithContext(ctx)
	db = db.Where("user_id = ? AND video_id = ?", validatedToken.GetUserId(), req.VideoId).Order("video_id").Find(&po)
	if db.Error != nil {
		f.l.Errorf("查询点赞记录失败：%s", db.Error.Error())
		return nil, db.Error
	}
	if db.RowsAffected != 0 {
		f.l.Errorf("记录已经存在，无需重复点赞")
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}
	po, err = f.NewFavoritePo(ctx, req)
	if err != nil {
		f.l.Errorf("由请求创建Po实例失败：%s", err)
		return nil, err
	}
	// 2、保存到数据库
	tx := f.db.WithContext(ctx).Create(&po)
	if tx.Error != nil {
		f.l.Errorf("保存数据失败：", err.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}
	return po, err
}

func (f *favoriteServiceImpl) DeleteFavoriteRecord(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoritePo, error) {
	//构造新的实例
	po, err := f.NewFavoritePo(ctx, req)
	if err != nil {
		return nil, err
	}
	//检查是否已经存在此纪录
	db := f.db.WithContext(ctx)
	db = db.Where("user_id = ? AND video_id = ?", po.UserId, po.VideoId).First(&favorite.FavoritePo{})
	db.Delete(&favorite.FavoritePo{})
	//记录不存在
	if db.RowsAffected == 0 {
		f.l.Errorf("点赞记录不存在：%s", err.Error())
		return po, exception.WithStatusCode(constant.ERROR_REMOVE)
	}
	return nil, nil
}

// 构建Favorite点赞实例
func (f *favoriteServiceImpl) NewFavoritePo(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoritePo, error) {
	//根据Token获取User
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := f.tokenService.ValidateToken(ctx, tokenReq)
	if err != nil {
		f.l.Errorf(err.Error())
		return nil, err
	}
	//构造请求
	favoritePo := favorite.NewFavoritePo()
	//TODO 雪花算法实现
	favoritePo.Id = time.Now().UnixNano()
	favoritePo.UserId = validatedToken.GetUserId()
	favoritePo.VideoId = req.VideoId
	return favoritePo, nil
}

// 获取喜欢视频列表
func (s *favoriteServiceImpl) getFavoriteListPo(ctx context.Context, po *favorite.FavoritePo) (
	[]*favorite.FavoritePo, error) {

	//向数据库查询所有数据
	db := s.db.WithContext(ctx)

	pos := make([]*favorite.FavoritePo, 0)
	if po.UserId > 0 && po.VideoId <= 0 {
		db = db.Where("user_id = ?", po.UserId)
	} else if po.VideoId > 0 && po.UserId <= 0 {
		db = db.Where("video_id = ?", po.VideoId)
	} else {
		s.l.Errorf("favorite getFavoriteListPo：你的参数可能有问题哟~")
		return pos, nil
	}

	//在favorite表中查找对应用户点赞的记录
	db.Find(&pos)

	if db.Error != nil {
		return nil, db.Error
	}

	return pos, nil
}

// getFavoriteCount 点赞数、被赞数 TODO  (这里可以精确控制、可按照 relation Count 里面做)
func (f *favoriteServiceImpl) getFavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (
	*favorite.FavoriteCountResponse, error) {

	resp := favorite.NewFavoriteCountResponse()
	db := f.db.WithContext(ctx).Model(&favorite.FavoritePo{})

	if req.UserId > 0 {
		// 查询点赞总数
		db.Where("user_id = ?", req.UserId).Count(&resp.FavoriteCount)
	}

	if db.Error != nil {
		f.l.Errorf("favorite getFavoriteCount：喜欢视频总数查询失败，%s", db.Error.Error())
		return resp, db.Error
	}

	if req.VideoIds != nil && len(req.VideoIds) > 0 {
		// 查询获得的点赞数
		db.Where("video_id IN ?", req.VideoIds).Count(&resp.AcquireFavoriteCount)
	}

	if db.Error != nil {
		f.l.Errorf("favorite getFavoriteCount：查询获得点赞数目失败，%s", db.Error.Error())
		return resp, db.Error
	}

	return resp, nil
}

func (s *favoriteServiceImpl) isFavorite(ctx context.Context, po *favorite.FavoritePo) (bool, error) {

	if po.UserId == 0 {
		return false, nil
	}

	// 只是查询，看看是否有条记录
	db := s.db.WithContext(ctx).
		Where("user_id = ? AND video_id = ?", po.UserId, po.VideoId).Find(favorite.NewFavoritePo())

	if db.Error != nil {
		s.l.Errorf("relation isFollowerByUId：查询错误，%s", db.Error.Error())
		return false, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	return db.RowsAffected == 1, nil
}
