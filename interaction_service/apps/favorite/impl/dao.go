// Created by yczbest at 2023/02/18 15:02

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
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
	po = favorite.NewDefaultFavoritePo()
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
	favoritePo := favorite.NewDefaultFavoritePo()
	//TODO 雪花算法实现
	favoritePo.Id = time.Now().UnixNano()
	favoritePo.UserId = validatedToken.GetUserId()
	favoritePo.VideoId = req.VideoId
	return favoritePo, nil
}

// 获取喜欢视频列表
func (f *favoriteServiceImpl) GetFavoriteListPo(ctx context.Context, req *favorite.GetFavoriteListRequest) ([]*video.VideoPo, error) {
	//根据Token获取User
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := f.tokenService.ValidateToken(ctx, tokenReq)
	if err != nil {
		f.l.Errorf(err.Error())
		return nil, err
	}
	//向数据库查询所有数据
	db := f.db.WithContext(ctx)
	// 鉴权token解析出的ID暂不做处理
	_ = validatedToken.GetUserId()
	//统计记录数量
	var count int64 = 0
	//左连接查询Video信息
	db.Table("favorite").Where("user_id = ?", req.UserId).Joins("left join video on video.id = favorite.user_id").Count(&count)
	if db.Error != nil {
		return nil, db.Error
	}
	pos := make([]*video.VideoPo, count)

	db.Joins("left join favorite on favorite.user_id = ? AND favorite.video_id = video.id", req.UserId).Find(&pos)
	if db.Error != nil {
		return nil, db.Error
	}
	return pos, nil

}

//TODO 返回指定video_id视频点赞数量
