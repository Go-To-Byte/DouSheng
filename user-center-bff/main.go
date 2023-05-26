// @Author: Ciusyan 2023/1/24
package main

import (
	"github.com/Go-To-Byte/DouSheng/dou-kit/cmd"

	// 驱动加载所有变量，主要是[IOC的实例]
	_ "github.com/Go-To-Byte/DouSheng/user-center-bff/common/all"
)

func main() {
	// 交由CLI启动
	cmd.Main()
}
