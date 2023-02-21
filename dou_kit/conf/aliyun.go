// @Author: Ciusyan 2023/2/8
package conf

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

//=====
// Aliyun 配置对象
//=====

type aliyun struct {
	Endpoint     string `toml:"endpoint" env:"ALIYUN_ENDPOINT"`
	AccessKey    string `toml:"access_key" env:"ALIYUN_ACCESS_KEY"`
	AccessSecret string `toml:"access_secret" env:"ALIYUN_ACCESS_SECRET"`
	Bucket       string `toml:"bucket" env:"BUCKET"`
	VideoDir     string `toml:"video_dir" env:"VIDEO_DIR"`
	ImageDir     string `toml:"image_dir" env:"IMAGE_DIR"`

	// 用于截取封面
	CoverStyle string `toml:"cover_style" env:"COVER_STYLE"`

	lock sync.Mutex
}

func NewDefaultAliyun() *aliyun {
	return &aliyun{
		Endpoint:     "<yourEndpoint>",
		AccessKey:    "<yourAccessKeyId>",
		AccessSecret: "<yourAccessKeySecret>",
	}
}

var (
	ossClient *oss.Client
)

// GetClient 获取 OSS Bucket
func (a *aliyun) GetClient() (*oss.Client, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if ossClient == nil {
		c, err := a.client()
		if err != nil {
			return nil, err
		}
		ossClient = c
	}

	// 根据配置的名称，获取具体的 Bucket
	return ossClient, nil
}

func (a *aliyun) client() (*oss.Client, error) {
	c, err := oss.New(a.Endpoint, a.AccessKey, a.AccessSecret)
	if err != nil {
		return nil, err
	}
	return c, nil
}
