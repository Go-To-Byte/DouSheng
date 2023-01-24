// @Author: Ciusyan 2023/1/24
package cmd

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps"
	_ "github.com/Go-To-Byte/DouSheng/apps/all"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

// log 为全局变量, 只需要load 即可全局使用, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 从Config里面的日志配置，来配置全局Logger对象
	lc := conf.C().Log
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
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没有把日志输出到文件
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	// 配置日志的输出格式
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	// 把日志配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "启动 Dousheng API服务",
	Short:   "启动 Dousheng API服务",
	Example: "go run main start",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1、加载配置文件&全局Logger对象

		if err := conf.LoadConfigFromToml(configFile); err != nil {
			return err
		}
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 2、类似于Mysql注入驱动的方式加载UserServiceImpl的 init方法，将依赖注入IOC
		// _ "github.com/Go-To-Byte/DouSheng/apps/user/impl"【User模块的ServiceImpl服务注入IOC】
		// 然后再初始化IOC容器里的服务
		apps.InitService()

		// 3、通过Gin启动服务
		// _ "github.com/Go-To-Byte/DouSheng/apps/user/http"【User模块的Gin服务注入IOC】
		// 初始化Gin服务、还有注册Gin服务对象的路由
		g := gin.Default()
		apps.InitGin(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	f := StartCmd.PersistentFlags()
	f.StringVarP(&configFile, "config", "f",
		"etc/dousheng.toml", "dousheng api 的配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
