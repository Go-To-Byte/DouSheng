// @Author: Ciusyan 2023/2/14
package custom

import (
	"errors"
	"fmt"

	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
)

type CodeMsg struct {
	// 状态码
	StatusCode int32 `json:"status_code"`
	// 消息
	StatusMsg string `json:"status_msg"`
}

func Ok(msg string) *CodeMsg {
	return New(constant.CODE_OK, msg)
}

func Bad(msg string) *CodeMsg {
	return New(constant.CODE_ERROR, msg)
}

// New returns a CodeMsg representing c and msg.
func New(c int32, msg string) *CodeMsg {
	return &CodeMsg{StatusCode: c, StatusMsg: msg}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c int32, format string, a ...interface{}) *CodeMsg {
	return New(c, fmt.Sprintf(format, a...))
}

func NewWithMsg(msg string) *CodeMsg {
	return New(constant.Msg2Code(msg), msg)
}

func NewWithCode(code constant.StatusCode) *CodeMsg {
	return NewWithMsg(constant.Code2Msg(code))
}

// Code returns the status code contained in s.
func (s *CodeMsg) Code() int32 {
	if s == nil {
		return constant.CODE_OK
	}
	return s.StatusCode
}

// Message returns the message contained in s.
func (s *CodeMsg) Message() string {
	if s == nil {
		return ""
	}
	return s.StatusMsg
}

// Err returns an immutable error representing s; returns nil if s.StatusCode() is OK.
func (s *CodeMsg) Err() error {
	if s.StatusCode == constant.CODE_OK {
		return nil
	}
	return &Exception{S: s}
}

func (s *CodeMsg) String() string {
	return fmt.Sprintf("custom error: code = %d desc = %s", s.StatusCode, s.StatusMsg)
}

/*---------------------------------------------*/
/*------------------分割线----------------------*/
/*---------------------------------------------*/

// Exception 异常对象
type Exception struct {
	S *CodeMsg
	// 可携带额外消息
	Details []interface{} `json:"details"`
}

func (e *Exception) Error() string {
	return e.S.String()
}

func (e *Exception) CodeMsg() *CodeMsg {
	return e.S
}

// Is 判断是否是自定义异常
func (e *Exception) Is(target error) bool {
	tse, ok := target.(*Exception)
	if !ok {
		return false
	}
	return errors.Is(e, tse)
}

// GetDetails 获取额外信息
func (e *Exception) GetDetails() []interface{} {
	if e == nil {
		return nil
	}
	return e.Details
}
