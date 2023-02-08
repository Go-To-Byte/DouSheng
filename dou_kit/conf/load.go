// @Author: Ciusyan 2023/2/7
package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	"github.com/hashicorp/consul/api"
	"github.com/infraboard/mcube/logger/zap"
)

//=====
// 用于加载全局配置
//=====

// LoadConfigFromToml 从Toml配置文件加载
func LoadConfigFromToml(filePath string) error {
	// 初始化全局对象
	cfg := NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, cfg)
	if err != nil {
		return fmt.Errorf("load config file error，path：%s，%s", filePath, err)
	}

	return cfg.LoadGlobal()
}

// LoadConfigFromEnv 从环境变量加载
func LoadConfigFromEnv() error {
	config := NewDefaultConfig()
	if err := env.Parse(config); err != nil {
		return err
	}
	return config.LoadGlobal()
}

// LoadGlobal 或者可以这样加载全局实例s
func (c *Config) LoadGlobal() error {
	// 给全局配置赋值
	global = c

	// 初始化 Consul 客户端
	consulCfg := api.DefaultConfig()
	consulCfg.Address = c.Consul.Addr()
	client, err := api.NewClient(consulCfg)
	if err != nil {
		return err
	}

	globalConsulClient = client
	return nil
}

// LoadGlobalLogger log 为全局变量, 只需要load 即可全局使用, 依赖全局配置先初始化
func LoadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 从Config里面的日志配置，来配置全局Logger对象
	lc := C().Log
	// 解析日志Level配置
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	// 使用默认配置初始化Logger全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的Level级别
	zapConfig.Level = level
	// 程序每启动一次，不必都生成一个新的日志文件
	zapConfig.Files.RotateOnStartup = false

	switch lc.To {
	case ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没有把日志输出到文件
		zapConfig.ToFiles = false
	case ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	// 配置日志的输出格式
	switch lc.Format {
	case JSONFormat:
		zapConfig.JSON = true
	}
	// 把日志配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}
