// @Author: Ciusyan 2023/1/25
package all

// 在这里统一自动注册IOC
import (
	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/token/impl"

	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/comment/api"
	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/favorite/api"
	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/user/api"
	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/video/api"

	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/message/api"
	_ "github.com/Go-To-Byte/DouSheng/api-rooter/apps/relation/api"

	// 加载切面对象
	_ "github.com/Go-To-Byte/DouSheng/api-rooter/protocol"
)
