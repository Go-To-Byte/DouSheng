// @Author: Ciusyan 2023/3/1
package utils

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
)

// V2P 将 value -> ptr
func V2P[T any](n T) *T {
	return &n
}

func TokenStrFromCtx(ctx context.Context) string {
	value := ctx.Value(constant.REQUEST_TOKEN)
	if value == nil {
		return ""
	}
	return value.(string)
}
