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
	OPERATE_OK = NewCodeMsg(0, "操作成功")
	SAVE_OK    = NewCodeMsg(0, "保存成功")
	REMOVE_OK  = NewCodeMsg(0, "删除成功")
	ACQUIRE_OK = NewCodeMsg(0, "获取成功")
)

var (
	BAD_REQUEST      = NewCodeMsg(400, "请求出错")
	BAD_UNAUTHORIZED = NewCodeMsg(401, "未授权")
	BAD_FORBIDDEN    = NewCodeMsg(403, "禁止访问")
	BAD_NOT_FOUND    = NewCodeMsg(404, "资源不存在")
	BAD_SERVER_ERROR = NewCodeMsg(500, "服务器内部错误")
)

var (
	ERROR_OPERATE       = NewCodeMsg(40001, "操作失败")
	ERROR_SAVE          = NewCodeMsg(40002, "保存失败")
	ERROR_REMOVE        = NewCodeMsg(40003, "删除失败")
	ERROR_UPLOAD_IMG    = NewCodeMsg(40004, "图片上传失败")
	ERROR_USER_INFO     = NewCodeMsg(40005, "Token校验失败")
	ERROR_ARGS_VALIDATE = NewCodeMsg(40006, "参数校验失败")
)

var (
	WRONG_USERNAME_NOT_EXIST = NewCodeMsg(50001, "用户不存在")
	WRONG_PASSWORD           = NewCodeMsg(50002, "密码错误")
	WRONG_CAPTCHA            = NewCodeMsg(50004, "验证码错误")
	WRONG_EXIST_USERS        = NewCodeMsg(50005, "用户已存在")
	WRONG_NO_TOKEN           = NewCodeMsg(60001, "没有Token，请登录")
	WRONG_TOKEN_EXPIRED      = NewCodeMsg(60002, "Token过期，请重新登录")
	WRONG_NO_PERMISSION      = NewCodeMsg(60005, "没有相关的操作权限")
)
