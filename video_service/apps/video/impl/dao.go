// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
	"time"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

func (s *videoServiceImpl) Insert(ctx context.Context, req *video.PublishVideoRequest) (
	*video.VideoPo, error) {

	// 获取 video po 对象
	po := video.NewVideoPoWithSave(req)

	// 2、保存到数据库
	tx := s.db.WithContext(ctx).Create(po)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}

	return po, nil
}

// 查询出视频列表
func (s *videoServiceImpl) query(ctx context.Context, req *video.FeedVideosRequest) (
	[]*video.VideoPo, error) {

	if req == nil {
		req = video.NewFeedVideosRequest()
	} else {
		// 接收到参数，那么设置分页对象
		req.Page = video.NewPageRequest()
	}

	db := s.db.WithContext(ctx)

	// 如果没有传 LatestTime
	if req.LatestTime == nil || *req.LatestTime == 0 {
		req.LatestTime = utils.V2P(time.Now().UnixMilli())
	}

	set := make([]*video.VideoPo, 10)

	// 构建分页 、排序、 查询
	db = db.Where("created_at < ?", req.LatestTime).
		Limit(int(req.Page.PageSize)).Offset(int(req.Page.Offset)).
		Order("created_at desc").Find(&set)

	if db.Error != nil {
		s.l.Errorf("video: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}

func (s *videoServiceImpl) listFromUserId(ctx context.Context, userId int64) (
	[]*video.VideoPo, error) {

	db := s.db.WithContext(ctx)
	set := make([]*video.VideoPo, 10)

	// 构建条件、排序、 查询
	s.db.Where("author_id = ?", userId).Order("created_at desc").Find(&set)
	if db.Error != nil {
		s.l.Errorf("video: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}
