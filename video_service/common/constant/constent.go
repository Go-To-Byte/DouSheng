// @Author: Ciusyan 2023/2/7
package constant

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
)

// video_service 服务的常量、枚举

var (
	BAD_NO_FILE     = custom.New(80001, "读取文件数据失败")
	BAD_UPLOAD_FILE = custom.New(80002, "上传文件失败")
)

const (
	REQUEST_FILE = "data"
)
