// @Author: Ciusyan 2023/2/7
package exception

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/user_center/common/constant"
)

type CustomException interface {
	error
	GetCodeMsg() *constant.CodeMsg
}

// Exception 异常对象
type Exception struct {
	Code int32       `json:"error_code"`
	Msg  string      `json:"message"`
	Meta interface{} `json:"meta"`
}

// WithCodeMsg 传递CodeMsg
func WithCodeMsg(msg *constant.CodeMsg) *Exception {
	return &Exception{
		Code: msg.StatusCode,
		Msg:  msg.StatusMsg,
		Meta: msg,
	}
}

// WithMsg 传递Msg
func WithMsg(format string, a ...any) *Exception {
	// 错误只带 msg，返回的 code 默认为 1
	return WithCodeMsg(constant.NewCodeMsg(1, fmt.Sprintf(format, a...)))
}

func (e *Exception) GetCodeMsg() *constant.CodeMsg {
	return e.Meta.(*constant.CodeMsg)
}

func (e *Exception) Error() string {
	return e.Msg
}
