// @Author: Ciusyan 2023/2/7
package all

// 在这里统一自动注册实例 [主要是放入IOC]

import (
	_ "github.com/Go-To-Byte/DouSheng/video_service/apps/video/api"
	_ "github.com/Go-To-Byte/DouSheng/video_service/apps/video/impl"

	// 加载切面对象
	_ "github.com/Go-To-Byte/DouSheng/video_service/protocol"
)
