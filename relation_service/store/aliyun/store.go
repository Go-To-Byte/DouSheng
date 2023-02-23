package aliyun

import (
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/rs/xid"
	"mime/multipart"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/video_service/store"
)

func NewAliOssStore(isSaveCover bool) *AliOssStore {
	a := conf.C().Aliyun
	c, err := a.GetClient()
	if err != nil {
		panic(err)
	}

	ossStore := &AliOssStore{
		client: c,
	}

	// 是否需要保存封面
	if isSaveCover {
		ossStore.coverStyle = a.CoverStyle
		ossStore.coverObjKey = fmt.Sprintf("%s%s.jpg", a.ImageDir, xid.New().String())
	}

	return ossStore
}

type AliOssStore struct {
	client *oss.Client

	// 用于截取视频封面
	coverStyle string

	// 封面路径
	coverObjKey string

	// 可以扩展其他功能[如：上传进度条等]
}

// NewParamByPath 通过本地路径上传到 oss
func NewParamByPath(bucketName, objDir, objName string, localPath string) *store.UploadParam {
	param := newCommonParam(bucketName, objDir, objName)
	// 本地文件路径
	param.LocalPath = localPath
	param.FromTo = store.FromLocalPath
	return param
}

// NewParamByStream 通过文件流上传到 oss
func NewParamByStream(bucketName, objDir, objName string, fileStream multipart.File) *store.UploadParam {
	param := newCommonParam(bucketName, objDir, objName)
	// 从文件流中获取
	param.FileStream = fileStream
	param.FromTo = store.FromStream
	return param
}

func newCommonParam(bucketName, objDir, objName string) *store.UploadParam {
	return &store.UploadParam{
		BucketName:     bucketName,
		ObjectDir:      objDir,
		ObjectFileName: objName,
	}
}

// Upload 上传文件到 阿里云的 OSS
func (s *AliOssStore) Upload(param *store.UploadParam) (*store.UploadResult, error) {
	if param == nil || param.FromTo == "" || param.BucketName == "" {
		return nil, fmt.Errorf("请正确使用上传参数")
	}

	bucket, err := s.client.Bucket(param.BucketName)
	if err != nil {
		return nil, err
	}

	objKey := param.ObjectKey()
	switch param.FromTo {
	case store.FromLocalPath:
		if err = bucket.PutObjectFromFile(objKey, param.LocalPath); err != nil {
			return nil, err
		}
	case store.FromStream:
		if err = bucket.PutObject(objKey, param.FileStream); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("暂不支持此From：%s, 可以选择：[%s、%s]",
			param.FromTo, store.FromLocalPath, store.FromStream)
	}

	result := store.NewUploadResult()
	result.RelativeURI = objKey

	// signedURL, err := bucket.SignURL(objKey, oss.HTTPGet, 600, oss.Process(s.coverStyle))

	// 看看是否需要保存封面
	if s.coverObjKey == "" {
		return result, nil
	}

	// 将封面持久化
	if err = s.saveCover(bucket, objKey); err != nil {
		return nil, err
	}
	result.CoverRelativeURI = s.coverObjKey
	return result, nil
}

// bucket：之前 bucket
// oldObjKey：在原来旧的文件相对路径
// https://help.aliyun.com/document_detail/55811.htm?spm=a2c4g.11186623.0.0.5b1f24cbXasqWY#concept-bf1-ssc-wdb
func (s *AliOssStore) saveCover(bucket *oss.Bucket, oldObjKey string) error {

	process := fmt.Sprintf(
		"%s|sys/saveas,o_%v,b_%v",
		s.coverStyle,
		base64.URLEncoding.EncodeToString([]byte(s.coverObjKey)),
		base64.URLEncoding.EncodeToString([]byte(bucket.BucketName)),
	)
	_, err := bucket.ProcessObject(oldObjKey, process)
	return err
}
