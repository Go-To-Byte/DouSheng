// @Author: Ciusyan 2023/2/7
package video

import (
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
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

func (r *PublishListRequest) Validate() error {
	return validate.Struct(r)
}

func NewPublishVideoRequest() *PublishVideoRequest {
	return &PublishVideoRequest{}
}

// TableName 指明表名 -> gorm 参数映射
func (*VideoPo) TableName() string {
	return AppName
}

func NewVideoPo() *VideoPo {
	return &VideoPo{}
}

func NewVideoPoWithSave(req *PublishVideoRequest) *VideoPo {
	return &VideoPo{
		Title:     req.Title,
		PlayUrl:   req.PlayUrl,
		CoverUrl:  req.CoverUrl,
		CreatedAt: time.Now().UnixMilli(),
		AuthorId:  req.UserId,
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

func NewFeedSet() *FeedSetResponse {
	return &FeedSetResponse{}
}

func NewPublishListRequest() *PublishListRequest {
	return &PublishListRequest{}
}

func NewPublishListResponse() *PublishListResponse {
	return &PublishListResponse{}
}

// Po2vo 将 videoPo -> video，并且会组合用户信息
// userMap：用户信息 [userId] = User
func (po *VideoPo) Po2vo(userMap map[int64]*user.User) *Video {
	// po -> vo
	return &Video{
		Id:       po.Id,
		Author:   userMap[po.AuthorId],
		PlayUrl:  utils.URLPrefix(po.PlayUrl),
		CoverUrl: utils.URLPrefix(po.CoverUrl),
		Title:    po.Title,
	}
}

func NewGetVideoRequest() *GetVideoRequest {
	return &GetVideoRequest{}
}
