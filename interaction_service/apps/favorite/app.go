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

// 点赞/取消赞 参数校验
func (r *FavoriteActionRequest) Validate() error {
	return validate.Struct(r)
}

// 获取喜欢视频列表 参数校验
func (r *GetFavoriteListRequest) Validate() error {
	return validate.Struct(r)
}

// 通过video_id获取喜欢视频数量参数校验
func (r *GetFavoriteCountByIdRequest) Validate() error {
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

func NewDefaultGetFavoriteListRequest() *GetFavoriteListRequest {
	return &GetFavoriteListRequest{}
}

// 获取视频列表响应体
func NewDefaultGetFavoriteListResponse() *GetFavoriteListResponse {
	return &GetFavoriteListResponse{}
}

// 获取视频点赞总数请求体
func NewDefaultGetFavoriteCountByIdRequest() *GetFavoriteCountByIdRequest {
	return &GetFavoriteCountByIdRequest{}
}

// 获取视频点赞总数响应体
func NewDefaultGetFavoriteCountByIdResponse() *GetfavoriteCountByIdResponse {
	return &GetfavoriteCountByIdResponse{}
}
