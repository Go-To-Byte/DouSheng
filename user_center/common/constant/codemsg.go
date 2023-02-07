// @Author: Ciusyan 2023/2/7
package constant

func NewCodeMsg(code int32, msg string) *CodeMsg {
	return &CodeMsg{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

// CodeMsg 自定义 msg+code 枚举
type CodeMsg struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

var (
	BAD_OPERATER = NewCodeMsg(30001, "操作失败")
	BAD_SAVE     = NewCodeMsg(30002, "保存失败")
)

var (
	OK_OPERATER = NewCodeMsg(40001, "操作成功")
	OK_REGISTER = NewCodeMsg(40002, "注册成功")
)

var (
	BAD_TOKEN_ERROR   = NewCodeMsg(50001, "Token校验失败")
	BAD_ARGS_VALIDATE = NewCodeMsg(50002, "参数校验失败")
	BAD_NAME_PASSWORD = NewCodeMsg(50003, "用户名或密码错误")
)

var (
	WARNING_USER_EXIST = NewCodeMsg(60001, "用户已存在")
)
