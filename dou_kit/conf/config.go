// @Author: Ciusyan 2023/2/7
package conf

//=====
// 项目配置汇总
//=====

// Config 将配置文件抽成一个对象
type Config struct {
	App     *app     `mapstructure:"app" json:"app" yaml:"app"`
	Log     *log     `mapstructure:"log" json:"log" yaml:"log"`
	MySQL   *mySQL   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mongodb *mongodb `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`

	// Consul 注册中心
	Consul *consul `mapstructure:"consul" json:"consul" yaml:"consul"`
	// aliyun oss 服务
	Aliyun *aliyun `mapstruct:"aliyun" json:"aliyun" yaml:"aliyun"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:     NewDefaultApp(),
		Log:     NewDefaultLog(),
		MySQL:   NewDefaultMySQL(),
		Mongodb: NewMongodb(),
		Consul:  NewDefaultConsul(),
	}
}

// 防止配置文件在运行时被更改，设置为私有的
var global *Config

// C 获取总的配置对象
func C() *Config {
	if global == nil {
		panic("加载全局配置失败")
	}
	return global
}
