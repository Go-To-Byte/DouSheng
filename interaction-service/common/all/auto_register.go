// @Author: Ciusyan 2023/2/7
package all

// 在这里统一自动注册实例 [主要是放入IOC]

import (
	_ "github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment/impl"
	_ "github.com/Go-To-Byte/DouSheng/interaction-service/apps/message/impl"
)
