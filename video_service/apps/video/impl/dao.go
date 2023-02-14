// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

func (s *videoServiceImpl) Insert(ctx context.Context, req *video.PublishVideoRequest) (
	*video.VideoPo, error) {

	// 获取 video po 对象
	po, err := s.getVideoPo(ctx, req)
	if err != nil {
		// 因为 getVideoPo 已经包装过 err 了，直接返回即可
		return nil, err
	}

	// 2、保存到数据库
	tx := s.db.WithContext(ctx).Create(po)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}
	return nil, err
}

func (s *videoServiceImpl) getVideoPo(ctx context.Context, req *video.PublishVideoRequest) (
	*video.VideoPo, error) {

	tokenReq := token.NewValidateTokenRequest(req.Token)

	// 这里主要是为了获取 用户ID
	validatedToken, err := s.tokenService.ValidateToken(ctx, tokenReq)

	if err != nil {
		s.l.Errorf(err.Error())
		// GRPC 调用，不需要继续包装了
		return nil, err
	}

	VideoPo := video.NewVideoPo(req)
	VideoPo.AuthorId = validatedToken.GetUserId()
	return VideoPo, nil
}

// 查询出视频列表
func (s *videoServiceImpl) query(ctx context.Context, req *video.FeedVideosRequest) (
	[]*video.VideoPo, error) {

	if req == nil {
		req = video.NewFeedVideosRequest()
	} else {
		req.Page = video.NewPageRequest()
	}

	db := s.db.WithContext(ctx)

	// 如果有传 LatestTime
	if req.LatestTime != nil {
		db = db.Where("created_at <= ?", *req.LatestTime)
	}

	set := make([]*video.VideoPo, 10)

	// 构建分页 、排序、 查询
	db = db.Limit(int(req.Page.PageSize)).Offset(int(req.Page.Offset)).
		Order("created_at desc").Find(&set)

	if db.Error != nil {
		s.l.Errorf("video: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}
