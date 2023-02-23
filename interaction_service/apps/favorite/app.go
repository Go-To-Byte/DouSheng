// Created by yczbest at 2023/02/18 13:30

package favorite

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "favorite"
)

var (
	validate = validator.New()
)

func NewFavoriteActionRequest() *FavoriteActionRequest {
	return &FavoriteActionRequest{}
}

// 点赞/取消赞 参数校验
func (r *FavoriteActionRequest) Validate() error {
	return validate.Struct(r)
}

// 获取喜欢视频列表 参数校验
func (r *GetFavoriteListRequest) Validate() error {
	return validate.Struct(r)
}

// 创建视频点赞响应体
func NewFavoriteActionResponse() *FavoriteActionResponse {
	return &FavoriteActionResponse{}
}

// 构建Po
func NewDefaultFavoritePo() *FavoritePo {
	return &FavoritePo{}
}

// TableName 指明表名 -> gorm 参数映射
func (*FavoritePo) TableName() string {
	return AppName
}

func po2vo(ctx context.Context, po *video.VideoPo) (*video.Video, error) {
	//获取User对象

	return &video.Video{
		Id: po.Id,
		//Author:
		PlayUrl:  po.PlayUrl,
		CoverUrl: po.CoverUrl,
		//FavoriteCount:
		//CommentCount:
		//IsFavorite:
		Title: po.Title,
	}, nil
}

func NewDefaultGetFavoriteListRequest() *GetFavoriteListRequest {
	return &GetFavoriteListRequest{}
}

// 获取视频列表响应体
func NewDefaultGetFavoriteListResponse() *GetFavoriteListResponse {
	return &GetFavoriteListResponse{}
}

func NewGetFavoriteListResponse(ctx context.Context, pos []*video.VideoPo) (*GetFavoriteListResponse, error) {
	set := make([]*video.Video, len(pos))
	if pos == nil || len(pos) <= 0 {
		// 只是没有查到，不应该抛异常出去
		res := NewDefaultGetFavoriteListResponse()
		return res, nil
	}
	for i, po := range pos {
		vo, err := po2vo(ctx, po)
		if err != nil {
			return nil, err
		}
		set[i] = vo
	}
	res := NewDefaultGetFavoriteListResponse()
	res.VideoList = set
	return res, nil
}
