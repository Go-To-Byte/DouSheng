// Created by yczbest at 2023/02/18 14:53

package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/video-service/apps/favorite"
)

// TODO 6、实现点赞模块

// FavoriteAction 视频点赞接口实现:成功返回nil,失败返回错误信息
func (f *favoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (
	*favorite.FavoriteActionResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}

// FavoriteList 实现获取喜欢视频列表
func (f *favoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (
	*favorite.FavoriteListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}

func (f *favoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (
	*favorite.IsFavoriteResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method IsFavorite not implemented")
}
