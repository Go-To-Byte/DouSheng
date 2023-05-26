// @Author: Ciusyan 2023/2/7
package constant

import (
	"fmt"
)

const (
	REQUEST_TOKEN = "token"
)

type StatusCode int32

const (
	CODE_OK    = 0
	CODE_ERROR = 1

	INTERNAL = "内部错误"
)

const (
	// OPERATE_OK 想要通过 code -> msg 那么定义code时，也必须唯一，所以只提供一个枚举值
	OPERATE_OK StatusCode = 0

	// SAVE_OK 下面的可表示其他成功：[如 custom.Ok(SAVE_OK)]
	SAVE_OK    = "保存成功"
	REMOVE_OK  = "删除成功"
	ACQUIRE_OK = "获取成功"
)

const (
	BAD_REQUEST      StatusCode = 400
	BAD_UNAUTHORIZED StatusCode = 401
	BAD_FORBIDDEN    StatusCode = 403
	BAD_NOT_FOUND    StatusCode = 404
	BAD_SERVER_ERROR StatusCode = 500
)

const (
	ERROR_OPERATE        StatusCode = 40001
	ERROR_SAVE           StatusCode = 40002
	ERROR_REMOVE         StatusCode = 40003
	ERROR_UPLOAD         StatusCode = 40004
	ERROR_TOKEN_VALIDATE StatusCode = 40005
	ERROR_ARGS_VALIDATE  StatusCode = 40006
	ERROR_ACQUIRE        StatusCode = 400007
)

const (
	WRONG_USER_NOT_EXIST StatusCode = 50001
	WRONG_PASSWORD       StatusCode = 50002
	WRONG_CAPTCHA        StatusCode = 50004
	WRONG_EXIST_USERS    StatusCode = 50005
	WRONG_NO_TOKEN       StatusCode = 50006
	WRONG_TOKEN_EXPIRED  StatusCode = 50007
	WRONG_NO_PERMISSION  StatusCode = 50008
)

var msgToCode = map[string]StatusCode{
	"操作成功": OPERATE_OK,

	"请求出错":    BAD_REQUEST,
	"未授权":     BAD_UNAUTHORIZED,
	"禁止访问":    BAD_FORBIDDEN,
	"资源不存在":   BAD_NOT_FOUND,
	"服务器内部错误": BAD_SERVER_ERROR,

	"操作失败":      ERROR_OPERATE,
	"保存失败":      ERROR_SAVE,
	"删除失败":      ERROR_REMOVE,
	"获取失败":      ERROR_ACQUIRE,
	"上传失败":      ERROR_UPLOAD,
	"Token校验失败": ERROR_TOKEN_VALIDATE,
	"参数校验失败":    ERROR_ARGS_VALIDATE,

	"用户不存在":         WRONG_USER_NOT_EXIST,
	"密码错误":          WRONG_PASSWORD,
	"验证码错误":         WRONG_CAPTCHA,
	"用户已存在":         WRONG_EXIST_USERS,
	"没有Token，请登录":   WRONG_NO_TOKEN,
	"Token过期，请重新登录": WRONG_NO_TOKEN,
	"没有相关的操作权限":     WRONG_NO_PERMISSION,

	// user-service
	"用户名或密码错误": BAD_NAME_PASSWORD,

	// video-service
	"读取文件数据失败": BAD_NO_FILE,
	"上传文件失败":   BAD_UPLOAD_FILE,
}

// user-service 服务的常量、枚举
const (
	USER_ID                      = "user_id"
	OK_REGISTER                  = "注册成功"
	BAD_NAME_PASSWORD StatusCode = 70001
)

// video-service 服务的常量、枚举
const (
	REQUEST_FILE               = "data"
	BAD_NO_FILE     StatusCode = 80001
	BAD_UPLOAD_FILE StatusCode = 80002
)

// relation service
const (
	FOLLOW_ACTION   int32 = 1
	UNFOLLOW_ACTION int32 = 2
)

func Msg2Code(msg string) int32 {
	if msg == "" {
		return CODE_ERROR
	}
	if _, ok := msgToCode[msg]; !ok {
		return CODE_ERROR
	}
	return int32(msgToCode[msg])
}

func Code2Msg(code StatusCode) string {
	for k, v := range msgToCode {
		if v != code {
			continue
		}
		return k
	}
	return fmt.Sprintf("未知错误，code = %d", code)
}
