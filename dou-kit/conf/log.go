// @Author: Ciusyan 2023/2/7
package conf

//=====
// 日志配置对象
//=====

// log 日志配置
type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

func NewDefaultLog() *log {
	return &log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// LogFormat 日志格式
type LogFormat string

const (
	TextFormat = LogFormat("text") // 文本格式
	JSONFormat = LogFormat("json") // Json 格式
)

// LogTo 日志记录到哪
type LogTo string

const (
	ToFile   = LogTo("file")   // 文件
	ToStdout = LogTo("stdout") // 标准输出
)
