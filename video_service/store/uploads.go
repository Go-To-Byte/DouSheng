// @Author: Ciusyan 2023/2/8
package store

import (
	"mime/multipart"
)

// Uploader 文件上传接口
type Uploader interface {
	// Upload 上传文件到云端
	Upload(*UploadParam) (*UploadResult, error)
	// Delete 删除云端的视频，也需要知道是删除哪里的资源，叫什么名字，可删除多个
	Delete(bucketName string, filePaths ...string) error
}

// NewUploadParam 从文件流中上传
func NewUploadParam(bucketName string, filePath string, file multipart.File) *UploadParam {
	return &UploadParam{
		BucketName: bucketName,
		FilePath:   filePath,
		FileStream: file,
	}
}

// UploadParam 文件上传参数
type UploadParam struct {
	// 操作云上哪一个资源
	BucketName string `json:"bucket_name"`
	// 上传的文件路径
	FilePath string
	// 待上传的文件流
	FileStream multipart.File `json:"file_stream"`
	// ...
}

// Validate 简单校验参数
func (u *UploadParam) Validate() bool {
	return u.BucketName != "" && u.FilePath != "" && u.FileStream != nil
}

func NewUploadResult() *UploadResult {
	return &UploadResult{}
}

// UploadResult 文件上传返回结果
type UploadResult struct {
	// 访问的绝对路径 url
	AbsoluteURL string `json:"absolute_url"`
	// 访问的相对路径 uri
	RelativeURI string `json:"relative_uri"`

	// 如果是视频，将封面相对路径返回
	CoverRelativeURI string `json:"cover_relative_uri"`
	// ...
}
