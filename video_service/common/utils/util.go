// @Author: Ciusyan 2023/2/14
package utils

const (
	// 视频和封面前缀URI
	urlPrefix = "https://ciusyan-dousheng.oss-cn-shenzhen.aliyuncs.com/"
)

// URLPrefix 添加上URL的前缀
func URLPrefix(uri string) string {
	return urlPrefix + uri
}
