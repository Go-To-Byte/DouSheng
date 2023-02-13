// @Author: Ciusyan 2023/2/7
package video

import (
	"github.com/go-playground/validator/v10"
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
		Title:    req.Title,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
	}
}

func NewPublishVideoResponse() *PublishVideoResponse {
	return &PublishVideoResponse{}
}
