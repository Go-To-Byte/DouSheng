// Created by yczbest at 2023/02/18 14:53

package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"

	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
)

// FavoriteAction 视频点赞接口实现:成功返回nil,失败返回错误信息
func (f *favoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (
	*favorite.FavoriteActionResponse, error) {
	//参数校验
	if err := req.Validate(); err != nil {
		f.l.Errorf("interaction: FavoriteAction 参数校验失败！%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}
	switch req.ActionType {
	//点赞操作
	case 1:
		_, err := f.InsertFavoriteRecord(ctx, req)
		//点赞失败处理
		if err != nil {
			f.l.Errorf("视频点赞失败:%s", err.Error())
			return nil, status.Error(codes.PermissionDenied,
				constant.Code2Msg(constant.ERROR_SAVE))
		}
		//点赞成功处理
		return favorite.NewFavoriteActionResponse(), err
	//取消点赞操作
	case 2:
		_, err := f.DeleteFavoriteRecord(ctx, req)
		//取消点赞失败
		if err != nil {
			f.l.Errorf("取消视频点赞失败:%s", err.Error())
			return nil, status.Error(codes.PermissionDenied,
				constant.Code2Msg(constant.ERROR_REMOVE))
		}
		//取消点赞成功
		return favorite.NewFavoriteActionResponse(), nil
	default:
		//ActionType参数错误
		f.l.Errorf("ActionType参数错误")
		return nil, status.Error(codes.PermissionDenied,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}
}

// FavoriteList 实现获取喜欢视频列表
func (f *favoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (
	*favorite.FavoriteListResponse, error) {

	// 参数校验
	if err := req.Validate(); err != nil {
		f.l.Errorf("interaction: FavoriteAction 参数校验失败！%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	//	获取喜欢视频列表
	pos, err := f.getFavoriteListPo(ctx, req)
	if err != nil {
		f.l.Errorf("interaction: FavoriteAction 参数校验失败！%s", err.Error())
		return nil, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	resp := favorite.NewFavoriteListResponse()

	// 根据列表分别获取用户信息
	videoList := make([]*video.Video, len(pos))
	resp.VideoList = videoList

	if len(videoList) == 0 {
		return resp, nil
	}

	// 将用户信息和视频信息组合成Response TODO：换成批量查询
	videoReq := video.NewGetVideoRequest()
	videoReq.Token = req.Token
	for i, po := range pos {
		videoReq.VideoId = po.VideoId
		videoVo, err := f.videoService.GetVideo(ctx, videoReq)
		if err != nil {
			f.l.Errorf(err.Error())
			return nil, err
		}

		resp.VideoList[i] = videoVo
	}

	return resp, nil
}

// FavoriteCount 获取 1、用户喜欢列表的数目 2、获取视频点赞数
func (f *favoriteServiceImpl) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (
	*favorite.FavoriteCountResponse, error) {

	resp, err := f.getFavoriteCount(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	return resp, nil
}
