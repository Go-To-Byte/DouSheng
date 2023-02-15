// @Author: Ciusyan 2023/2/14
package utils

// V2P 将 value -> ptr
func V2P[T any](n T) *T {
	return &n
}

const (
	// 视频和封面前缀URI
	urlPrefix = "https://ciusyan-dousheng.oss-cn-shenzhen.aliyuncs.com/"
)

// URLPrefix 添加上URL的前缀
func URLPrefix(uri string) string {
	return urlPrefix + uri
}
