// Created by yczbest at 2023/02/18 13:30

package favorite

import (
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

// Validate 点赞/取消赞 参数校验
func (r *FavoriteActionRequest) Validate() error {
	return validate.Struct(r)
}

// Validate 获取喜欢视频列表 参数校验
func (r *FavoriteListRequest) Validate() error {
	return validate.Struct(r)
}

// NewFavoriteActionResponse 创建视频点赞响应体
func NewFavoriteActionResponse() *FavoriteActionResponse {
	return &FavoriteActionResponse{}
}

// NewFavoritePo 构建Po
func NewFavoritePo() *FavoritePo {
	return &FavoritePo{}
}

// TableName 指明表名 -> gorm 参数映射
func (*FavoritePo) TableName() string {
	return AppName
}

func NewFavoriteListRequest() *FavoriteListRequest {
	return &FavoriteListRequest{}
}

// NewFavoriteListResponse 获取视频列表响应体
func NewFavoriteListResponse() *FavoriteListResponse {
	return &FavoriteListResponse{}
}

func NewFavoriteCountResponse() *FavoriteCountResponse {
	return &FavoriteCountResponse{}
}
