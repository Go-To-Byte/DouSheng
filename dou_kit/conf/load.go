// @Author: Ciusyan 2023/2/7
package conf

import (
	"encoding/json"
	"fmt"
	envir "github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/caarlos0/env"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"os"
)

//=====
// 用于加载全局配置
//=====

// LoadConfigFromToml 从Toml配置文件加载
func LoadConfig(filePath string) error {
	// 初始化全局对象
	cfg := InitConfig(filePath)

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
	consulCfg.Address = c.Consul.Register.Addr()
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

func InitConfig(flag string) *Config {
	/* 配置文件选择 参数>flag解析>自定义常量>gin环境 */
	var config string
	if flag == "" {
		if os.Getenv(envir.DebugEnv) == "" {
			switch gin.Mode() {
			case gin.DebugMode:
				config = envir.DebugEnv
			case gin.ReleaseMode:
				config = envir.ProdEnv
			}
		} else {
			config = envir.DebugEnv

		}
	} else {
		config = flag
	}
	// viper配置
	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	// 读
	err := vip.ReadInConfig()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	// 监控
	vip.WatchConfig()
	// 开启监控
	vip.OnConfigChange(func(in fsnotify.Event) {
		change := in.Name
		fmt.Printf("%s changed\n", change)
		// 绑定nacos
		if err = vip.Unmarshal(&Nacos); err != nil {
			panic(err.Error())
		}
	})
	// 绑定 nacos
	if err = vip.Unmarshal(&Nacos); err != nil {
		panic(err.Error())
	}

	/*                  从ncos读取配置                              */
	sc := []constant.ServerConfig{
		{
			IpAddr: Nacos.Host,
			Port:   Nacos.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         Nacos.Namespace, // 如果需要支持多namespace，可以场景有多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: Nacos.DataId,
		Group:  Nacos.Group})

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	err = json.Unmarshal([]byte(content), &global)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	return global
}
