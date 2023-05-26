// Created by yczbest at 2023/02/21 14:59

package comment

import (
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "comment"
)

var (
	validate = validator.New()
)

func NewCommentActionRequest() *CommentActionRequest {
	return &CommentActionRequest{}
}

func (c *CommentActionRequest) Validate() error {
	return validate.Struct(c)
}

func (r *GetCommentListRequest) Validate() error {
	return validate.Struct(r)
}

func NewDefaultCommentActionResponse() *CommentActionResponse {
	return &CommentActionResponse{}
}

func NewDefaultCommentPo() *CommentPo {
	return &CommentPo{}
}

// TableName 指明表名 -> gorm 参数映射
func (*CommentPo) TableName() string {
	return AppName
}

func NewDefaultGetCommentListResponse() *GetCommentListResponse {
	return &GetCommentListResponse{}
}

func NewDefaultGetCommentListRequest() *GetCommentListRequest {
	return &GetCommentListRequest{}
}
