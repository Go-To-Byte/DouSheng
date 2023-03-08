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

const (
	// 点赞操作
	FAVORITE_ACTION = 1
	// 取消点赞
	FAVORITE_CANCEL_ACTION = 2
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
	case FAVORITE_ACTION:
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
	case FAVORITE_CANCEL_ACTION:
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

	po := favorite.NewFavoritePo()
	po.UserId = req.UserId
	//	获取喜欢视频列表
	pos, err := f.getFavoriteListPo(ctx, po)
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

func (s *favoriteServiceImpl) FavoriteCountMap(ctx context.Context, req *favorite.FavoriteMapRequest) (
	*favorite.FavoriteMapResponse, error) {

	// 1、参数校验
	resp := favorite.NewFavoriteMapResponse()
	if req.VideoIds == nil || len(req.VideoIds) <= 0 {
		s.l.Errorf("user userList：你的参数可能有问题哟~")
		return resp, nil
	}

	// 2、获取每个视频的点赞数+是否关注
	resp.FavoriteMap = make(map[int64]*favorite.FavoriteMap, 10)
	for _, v := range req.VideoIds {
		// 获取点赞数
		favoriteReq := favorite.NewFavoriteCountRequest()
		favoriteReq.VideoIds = []int64{v}
		favoriteResp, err := s.getFavoriteCount(ctx, favoriteReq)
		if err != nil {
			return resp, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
		}

		// 查询是否点赞
		po := favorite.NewFavoritePo()
		po.UserId = req.UserId
		po.VideoId = v
		isFavorite, err := s.isFavorite(ctx, po)
		if err != nil {
			return resp, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
		}

		resp.FavoriteMap[v] = favorite.NewFavoriteMap(favoriteResp.AcquireFavoriteCount, isFavorite)
	}

	return resp, nil
}
