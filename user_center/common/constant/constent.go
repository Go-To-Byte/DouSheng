// @Author: Ciusyan 2023/2/7
package constant

import "github.com/Go-To-Byte/DouSheng/dou_kit/constant"

// user_center 服务的常量、枚举

var (
	OK_REGISTER       = constant.NewCodeMsg(0, "注册成功")
	BAD_NAME_PASSWORD = constant.NewCodeMsg(70001, "用户名或密码错误")
)

const (
	USER_ID = "user_id"
)
