// @Author: Ciusyan 2023/1/23
package conf

// LogFormat 日志格式
type LogFormat string

const (
	TextFormat = LogFormat("text") // 文本格式
	JSONFormat = LogFormat("json") // Json 格式
)

// LogTo 日志记录到哪
type LogTo string

const (
	ToFile    = LogTo("file")   // 文件
	ToStdoutL = LogTo("stdout") // 标准输出
)
