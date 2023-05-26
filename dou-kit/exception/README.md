# 包装网络层对象通用异常返回

注意：这里仅用于包装网络层的异常返回结果。因为官方给的接口文档有点头大

这里需要配合constant里的枚举类：CodeMsg

提供两个方法，WithCodeMsg(枚举值)、WithMsg(错误消息)

```go

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

```

之后可继续封装，这里比较随意
