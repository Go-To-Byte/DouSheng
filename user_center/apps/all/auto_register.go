// @Author: Ciusyan 2023/1/25
package all

// 在这里统一自动注册IOC
import (
	_ "github.com/Go-To-Byte/DouSheng/user_center/apps/user/api"
	_ "github.com/Go-To-Byte/DouSheng/user_center/apps/user/impl"

	_ "github.com/Go-To-Byte/DouSheng/user_center/apps/token/impl"
)
