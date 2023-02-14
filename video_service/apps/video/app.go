// @Author: Ciusyan 2023/2/7
package video

import (
	"github.com/Go-To-Byte/DouSheng/video_service/common/utils"
	"github.com/go-playground/validator/v10"
	"time"
)

const (
	AppName = "video"
)

var (
	validate = validator.New()
)

// Validate 参数校验
func (r *PublishVideoRequest) Validate() error {
	return validate.Struct(r)
}

func NewPublishVideoRequest() *PublishVideoRequest {
	return &PublishVideoRequest{}
}

// TableName 指明表名 -> gorm 参数映射
func (*VideoPo) TableName() string {
	return AppName
}

func NewVideoPo(req *PublishVideoRequest) *VideoPo {
	return &VideoPo{
		Title:     req.Title,
		PlayUrl:   req.PlayUrl,
		CoverUrl:  req.CoverUrl,
		CreatedAt: time.Now().UnixMilli(),
	}
}

func NewPublishVideoResponse() *PublishVideoResponse {
	return &PublishVideoResponse{}
}

func NewFeedVideosRequest() *FeedVideosRequest {
	return &FeedVideosRequest{
		Page:       NewPageRequest(),
		LatestTime: utils.V2P(time.Now().UnixMilli()),
	}
}

func NewPageRequest() *PageRequest {
	return &PageRequest{
		Offset:   0,
		PageSize: 20,
	}
}

func (r *FeedSetResponse) Length() int64 {
	return int64(len(r.VideoList))
}

func NewFeedSet() *FeedSetResponse {
	return &FeedSetResponse{}
}
