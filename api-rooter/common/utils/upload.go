// @Author: Ciusyan 2023/2/8
package utils

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"mime/multipart"
	"path"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/Go-To-Byte/DouSheng/video-service/store"
	"github.com/Go-To-Byte/DouSheng/video-service/store/aliyun"
)

func UploadFile(ctx context.Context, file *multipart.FileHeader, ch chan *UploadMsg) {
	if file == nil || file.Size <= 0 {
		// 传递一个 Err 出去
		ch <- newErrMsg(exception.WithStatusMsg("请正确的上传文件"))
	}

	// 获取文件扩展名
	fileName := fmt.Sprintf("%s%s", xid.New().String(), path.Ext(file.Filename))

	// 构建请求参数
	ossCfg := conf.C().Aliyun

	client, err := ossCfg.GetClient()
	if err != nil {
		// 传递一个 Err 出去
		ch <- newErrMsg(err)
	}

	coverPath := fmt.Sprintf("%s%s.jpg", ossCfg.ImageDir, xid.New().String())

	// 这里需要面向接口
	var uploader store.Uploader = aliyun.NewAliOssStore(client, coverPath, ossCfg.CoverStyle)

	open, err := file.Open()
	if err != nil {
		// 传递一个 Err 出去
		ch <- newErrMsg(err)
	}

	param := store.NewUploadParam(ossCfg.Bucket, ossCfg.VideoDir+fileName, open)

	result, err := uploader.Upload(param)
	if err != nil {
		// 传递一个 Err 出去
		ch <- newErrMsg(err)
	}

	// 上传成功并传递参数出去
	ch <- newSuccessMsg(uploader, result)
}

// UploadMsg 用于传递信息·
type UploadMsg struct {
	// 传递错误
	Err error
	// 传递上传的结果
	UploadResult *store.UploadResult
	// 传递Uploader
	uploader store.Uploader
}

func newSuccessMsg(uploader store.Uploader, uploadResult *store.UploadResult) *UploadMsg {
	return &UploadMsg{
		uploader:     uploader,
		UploadResult: uploadResult,
	}
}

func newErrMsg(err error) *UploadMsg {
	return &UploadMsg{
		Err: err,
	}
}

// Delete 删除上传成功的视频和封面，只有当文件上传成功、但是入库失败时才会调用
func (m *UploadMsg) Delete() {
	if m.Err != nil {
		return
	}

	ossCfg := conf.C().Aliyun
	// 调用对应删除的逻辑，这里放在后台删除即可，不管他成不成功
	go m.uploader.Delete(ossCfg.Bucket, m.UploadResult.RelativeURI, m.UploadResult.CoverRelativeURI)
}
