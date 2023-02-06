// @Author: Ciusyan 2023/2/7
package constant

func NewResp(code int32, msg string) *Response {
	return &Response{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var (
	BAD_TOKEN_ERROR = NewResp(1, "ss")
)
