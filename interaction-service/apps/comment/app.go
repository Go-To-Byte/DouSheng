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

// 发布/删除评论 参数校验
func (c *CommentActionRequest) Validate() error {
	return validate.Struct(c)
}

// 获取评论列表 参数校验
func (r *GetCommentListRequest) Validate() error {
	return validate.Struct(r)
}

// 获取评论列表 参数校验
func (r *GetCommentCountByIdRequest) Validate() error {
	return validate.Struct(r)
}

// 创建评论操作响应体
func NewDefaultCommentActionResponse() *CommentActionResponse {
	return &CommentActionResponse{}
}

// 构建Po
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

func NewDefaultGetCommentCountByIdRequest() *GetCommentCountByIdRequest {
	return &GetCommentCountByIdRequest{}
}

func NewDefaultGetCommentCountByIdResponse() *GetCommentCountByIdResponse {
	return &GetCommentCountByIdResponse{}
}

func NewCommentMapRequest() *CommentMapRequest {
	return &CommentMapRequest{}
}

func NewCommentMapResponse() *CommentMapResponse {
	return &CommentMapResponse{}
}
