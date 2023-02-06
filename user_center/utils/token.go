// @Author: Ciusyan 2023/2/6
package utils

import (
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
