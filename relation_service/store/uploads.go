package store

import (
	"mime/multipart"
)

// Uploader 文件上传接口
type Uploader interface {
	// Upload 上传文件到云端
	Upload(*UploadParam) (*UploadResult, error)
}

// From 从哪里获取文件
type From string

const (
	FromStream    = From("stream")
	FromLocalPath = From("localPath")
)

// UploadParam 文件上传参数
type UploadParam struct {
	// 可用于云上操作
	BucketName string `json:"bucket_name"`

	// 目标文件夹
	ObjectDir string `json:"object_dir"`
	// 目标文件名称
	ObjectFileName string `json:"object_file_name"`

	// 待上传文件的本地路径
	LocalPath string `json:"local_path"`
	// 待上传的文件流
	FileStream multipart.File `json:"file_stream"`

	// 从哪里加载文件
	FromTo From `json:"from_to"`

	// ...
}

// ObjectKey 获取 ObjectKey
func (p *UploadParam) ObjectKey() string {
	return p.ObjectDir + p.ObjectFileName
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
