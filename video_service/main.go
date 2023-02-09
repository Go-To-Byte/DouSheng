// @Author: Ciusyan 2023/2/7
package main

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/cmd"

	// 驱动加载所有变量，主要是[IOC的实例]
	_ "github.com/Go-To-Byte/DouSheng/video_service/common/all"
)

func main() {
	// 交由CLI启动
	cmd.Main()
}
