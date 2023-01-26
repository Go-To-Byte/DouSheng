// Author: BeYoung
// Date: 2023/1/26 0:41
// Software: GoLand

package run

import "github.com/gin-gonic/gin"

// GetRouter 获取路由
func GetRouter() *gin.Engine {
	var r *gin.Engine // 单例模式（懒汉模式）
	if r == nil {
		r = gin.Default()
	}
	return r
}
