// @Author: Ciusyan 2023/2/8
package aliyun

import (
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"

	"github.com/Go-To-Byte/DouSheng/video-service/store"
)

// NewAliOssStore 是否需要直接保存持久化封面
func NewAliOssStore(client *oss.Client, coverPath, coverStyle string) *AliOssStore {
	return &AliOssStore{
		client:      client,
		coverObjKey: coverPath,
		coverStyle:  coverStyle,
	}
}

type AliOssStore struct {
	client *oss.Client

	// 用于截取视频封面
	coverStyle string

	// 封面路径
	coverObjKey string

	// 可以扩展其他功能[如：上传进度条等]
}

// Upload 上传文件到 阿里云的 OSS
func (s *AliOssStore) Upload(param *store.UploadParam) (*store.UploadResult, error) {
	if param == nil || !param.Validate() {
		return nil, fmt.Errorf("请正确使用上传参数")
	}

	// 获取对应 Bucket 资源抽象
	bucket, err := s.client.Bucket(param.BucketName)
	if err != nil {
		return nil, err
	}

	// 将 file 文件上传到对应bucket，变成一个对象，名字为 param.FilePath,
	if err = bucket.PutObject(param.FilePath, param.FileStream); err != nil {
		return nil, err
	}

	result := store.NewUploadResult()
	// 这个路径其实也可以不用给出去，
	result.RelativeURI = param.FilePath

	// 看看是否需要保存封面
	if s.coverObjKey == "" {
		return result, nil
	}

	// 将封面持久化
	if err = s.saveCover(bucket, param.FilePath); err != nil {
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

// Delete 删除 Object
func (a *AliOssStore) Delete(bucketName string, filePaths ...string) error {
	if bucketName == "" || filePaths == nil || len(filePaths) == 0 {
		return fmt.Errorf("请正确使用上传参数")
	}

	// 获取Bucket
	bucket, err := a.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	for i := 0; i < len(filePaths); i++ {
		// 挨个删除删除
		if err := bucket.DeleteObject(filePaths[i]); err != nil {
			// 这里之记录日志，别影响删除后面的 Object // TODO 日志
			log.Printf("删除第[%d]个Object失败：%s", i+1, err)
		}
	}

	return nil
}
