// Created by yczbest at 2023/02/18 14:53

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 视频点赞接口实现:成功返回nil,失败返回错误信息
func (f *favoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
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

// 实现获取喜欢视频列表
func (f *favoriteServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.GetFavoriteListRequest) (*favorite.GetFavoriteListResponse, error) {
	//参数校验
	if err := req.Validate(); err != nil {
		f.l.Errorf("interaction: FavoriteAction 参数校验失败！%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}
	//	获取喜欢视频列表
	pos, err := f.GetFavoriteListPo(ctx, req)
	if err != nil {
		return nil, err
	}
	// 根据列表分别获取用户信息
	videoList := make([]*video.Video, len(pos))
	rsp := favorite.NewDefaultGetFavoriteListResponse()
	if len(videoList) == 0 {
		rsp.VideoList = videoList
		return rsp, nil
	}
	//将用户信息和视频信息组合成Response
	for index, po := range pos {
		videoReq := video.NewGetVideoRequest()
		videoReq.VideoId = po.VideoId
		videoVo, err := f.videoService.GetVideo(ctx, videoReq)
		if err != nil {
			return nil, err
		}
		videoList[index] = videoVo
	}
	rsp = favorite.NewDefaultGetFavoriteListResponse()
	rsp.VideoList = videoList
	return rsp, nil
}
