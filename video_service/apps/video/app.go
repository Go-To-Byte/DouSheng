// @Author: Ciusyan 2023/2/7
package video

import "github.com/go-playground/validator/v10"

const (
	AppName = "video"
)

var (
	validate = validator.New()
)

// Validate 校验参数
func (r *PublishVideoRequest) Validate() error {
	return validate.Struct(r)
}

func NewPublishVideoRequest() *PublishVideoRequest {
	return &PublishVideoRequest{}
}

// TableName 指明表名 -> gorm 参数映射
func (*Video) TableName() string {
	return AppName
}

func NewVideo(req *PublishVideoRequest) *Video {
	return &Video{
		Title:    req.Title,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
	}
}
