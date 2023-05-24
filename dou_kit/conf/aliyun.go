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
	Endpoint     string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKey    string `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	AccessSecret string `mapstructure:"access_secret" json:"access_secret" yaml:"access_secret"`
	Bucket       string `mapstructure:"bucket" json:"bucket" yaml:"buck"`
	VideoDir     string `mapstructure:"video_dir" json:"video_dir" yaml:"video_dir"`
	ImageDir     string `mapstructure:"image_dir" json:"image_dir" yaml:"image_dir"`

	// 用于截取封面
	CoverStyle string `mapstructure:"cover_style" json:"cover_style" yaml:"cover_style"`

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
