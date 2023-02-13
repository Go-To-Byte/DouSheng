// @Author: Ciusyan 2023/2/7
package exception

import (
	"fmt"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
)

// WithStatusMsgf 传递Msg
func WithStatusMsgf(format string, a ...any) *custom.Exception {
	// 错误只带 msg，返回的 code 默认为 1
	return WithStatusMsg(fmt.Sprintf(format, a...))
}

func WithStatusMsg(msg string) *custom.Exception {
	return WithCodeMsg(custom.NewWithMsg(msg))
}

func WithStatusCode(code constant.StatusCode) *custom.Exception {
	return WithStatusMsg(constant.Code2Msg(code))
}

// WithCodeMsg 传递CodeMsg
func WithCodeMsg(codeMsg *custom.CodeMsg) *custom.Exception {
	return WithDetails(codeMsg.Code(), codeMsg.Message())
}

// WithDetails 传递Details
func WithDetails(code int32, msg string, details ...interface{}) *custom.Exception {
	return &custom.Exception{
		S:       custom.New(code, msg),
		Details: details,
	}
}

// Err returns an error representing c and msg.  If c is OK, returns nil.
func Err(c int32, msg string) error {
	return custom.New(c, msg).Err()
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c int32, format string, a ...interface{}) error {
	return Err(c, fmt.Sprintf(format, a...))
}
