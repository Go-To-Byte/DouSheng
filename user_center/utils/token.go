// @Author: Ciusyan 2023/2/6
package utils

import (
	"github.com/Go-To-Byte/DouSheng/user_center/common/constant"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
)

// MakeBearer 生成Base64的字符串
func MakeBearer(length int) string {
	charList := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	t := make([]string, length)
	rand.Seed(time.Now().UnixNano() + int64(length) + rand.Int63n(10000))
	for i := 0; i < length; i++ {
		rn := rand.Intn(len(charList))
		w := charList[rn : rn+1]
		t = append(t, w)
	}

	return strings.Join(t, "")
}

// GetToken 从gin的Ctx种获取Token
func GetToken(ctx *gin.Context) string {

	// 1、从 header 中获取
	// ...我们这里都是从参数中获取的

	// 2、从query string 中获取
	tk := ctx.Query(constant.REQUEST_TOKEN)
	if tk != "" {
		return tk
	}

	// 3、从 body 中获取
	tk = ctx.PostForm(constant.REQUEST_TOKEN)
	if tk != "" {
		return tk
	}

	// 4、都没有，就返回 ""
	return ""
}
