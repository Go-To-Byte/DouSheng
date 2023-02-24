// Created by yczbest at 2023/02/18 14:53

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
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
				constant.Code2Msg(constant.WRONG_USER_NOT_EXIST))
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
	}
	//ActionType参数错误
	f.l.Errorf("ActionType参数错误")
	return nil, status.Error(codes.PermissionDenied,
		constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
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
	//将用户信息和视频信息组合成Response
	for index, po := range pos {
		userReq := user.NewUserInfoRequest()
		userReq.UserId = po.AuthorId
		userRsp, err := f.userService.UserInfo(ctx, userReq)
		if err != nil {
			return nil, err
		}
		videoVo, err := f.videoPo2videoPo(po, userRsp.User)
		if err != nil {
			return nil, err
		}
		videoList[index] = videoVo
	}
	res := favorite.NewDefaultGetFavoriteListResponse()
	res.VideoList = videoList
	return res, nil
}

func (f *favoriteServiceImpl) videoPo2videoPo(po *video.VideoPo,
	userInfo *user.User) (*video.Video, error) {
	// 也可以是单个查询
	if userInfo == nil {
		// 走GRPC调用，获取用户信息
		req := user.NewUserInfoRequest()
		req.UserId = po.AuthorId
	}

	// po -> vo
	return &video.Video{
		Id:       po.Id,
		Author:   userInfo,
		PlayUrl:  utils.URLPrefix(po.PlayUrl),
		CoverUrl: utils.URLPrefix(po.CoverUrl),
		Title:    po.Title,
	}, nil
}
