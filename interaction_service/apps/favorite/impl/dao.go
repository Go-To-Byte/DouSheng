// Created by yczbest at 2023/02/18 15:02

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
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
func (f *favoriteServiceImpl) getFavoriteListPo(ctx context.Context, req *favorite.FavoriteListRequest) (
	[]*favorite.FavoritePo, error) {

	//向数据库查询所有数据
	db := f.db.WithContext(ctx)
	//统计记录数量
	//在favorite表中查找对应用户点赞的记录
	pos := make([]*favorite.FavoritePo, 0)
	db.Where("user_id = ?", req.UserId).Find(&pos)

	if db.Error != nil {
		return nil, db.Error
	}

	return pos, nil
}

// FavoriteCount TODO 返回指定video_id视频点赞数量
func (f *favoriteServiceImpl) getFavoriteCount(ctx context.Context, req *favorite.FavoritePo) (int64, error) {

	db := f.db.WithContext(ctx).Model(&favorite.FavoritePo{})
	if req.UserId > 0 && req.VideoId <= 0 { // 查用户点赞视频数
		db.Where(" user_id = ?", req.UserId)
	} else if req.VideoId > 0 && req.UserId <= 0 { // 查视频点赞数
		db.Where(" video_id = ?", req.VideoId)
	} else {
		f.l.Error("favorite getFavoriteCount：你的参数可能有问题哟~")
		return 0, nil
	}

	var count int64
	db.Count(&count)

	if db.Error != nil {
		f.l.Errorf("喜欢视频总数查询失败:%s", db.Error.Error())
		return 0, db.Error
	}

	return count, nil
}
