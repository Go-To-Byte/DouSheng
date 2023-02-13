// @Author: Ciusyan 2023/2/8
package utils

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/rs/xid"
	"mime/multipart"
	"path"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/video_service/store"
	"github.com/Go-To-Byte/DouSheng/video_service/store/aliyun"
)

func UploadFile(file *multipart.FileHeader) (*store.UploadResult, error) {
	if file == nil || file.Size <= 0 {
		return nil, exception.WithStatusMsg("请正确的上传文件")
	}

	// 获取文件扩展名
	fileName := fmt.Sprintf("%s%s", xid.New().String(), path.Ext(file.Filename))

	// 这里需要面向接口
	var uploader store.Uploader = aliyun.NewAliOssStore(true)

	open, err := file.Open()
	if err != nil {
		return nil, err
	}

	// 构建请求参数
	ossCfg := conf.C().Aliyun
	param := aliyun.NewParamByStream(
		ossCfg.Bucket, ossCfg.VideoDir, fileName, open,
	)

	// 上传并返回
	return uploader.Upload(param)
}
