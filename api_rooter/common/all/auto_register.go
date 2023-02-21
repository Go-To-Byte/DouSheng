// @Author: Ciusyan 2023/1/25
package all

// 在这里统一自动注册IOC
import (
	_ "github.com/Go-To-Byte/DouSheng/api_rooter/apps/token/impl"

	_ "github.com/Go-To-Byte/DouSheng/api_rooter/apps/user/api"
	_ "github.com/Go-To-Byte/DouSheng/api_rooter/apps/video/api"

	// 加载切面对象
	_ "github.com/Go-To-Byte/DouSheng/api_rooter/protocol"
)
