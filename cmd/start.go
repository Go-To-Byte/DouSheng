// @Author: Ciusyan 2023/1/24
package cmd

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/Go-To-Byte/DouSheng/protocol"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"

	// 驱动加载所有需要放入IOC的实例
	_ "github.com/Go-To-Byte/DouSheng/apps/all"
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
		// ========
		// 1、加载配置文件&全局Logger对象
		// ========

		if err := conf.LoadConfigFromToml(configFile); err != nil {
			return err
		}
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// ========
		// 2、初始化IOC容器中的所有服务
		// ========

		// 类似于Mysql注入驱动的方式加载UserServiceImpl的 init方法，将依赖注入IOC
		// _ "github.com/Go-To-Byte/DouSheng/apps/user/impl"【User模块的ServiceImpl服务注入IOC】
		if err := ioc.InitAllDependencies(); err != nil {
			return err
		}

		// ========
		// 3、使用服务管理者来处理服务的关闭和开启
		// ========
		serviceManager := NewManager()

		// 用于接收信号的信道
		ch := make(chan os.Signal, 1)
		defer close(ch)
		// 接收这几种信号
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)

		// 需要在后台等待关闭
		go serviceManager.WaitStop(ch)

		return serviceManager.start() // 开启服务
	},
}

func NewManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("CLI"),
	}
}

// 用于管理服务的开启、和关闭
type manager struct {
	http *protocol.HttpService
	l    logger.Logger
}

func (m *manager) start() error {

	// 打印加载好的服务
	m.l.Infof("已加载的内部服务: %s", ioc.ExistingInternalDependencies())
	m.l.Infof("已加载的Gin HTTP服务: %s", ioc.ExistingGinDependencies())

	// 注：这属于正常关闭："http: Server closed"
	if err := m.http.Start(); err != nil && err.Error() != "http: Server closed" {
		return err
	}
	return nil
}

// WaitStop 处理来自外部的中断信号，比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Infof("接受到信号：%s", v)
			m.http.Stop()
		}
	}
}

func init() {
	f := StartCmd.PersistentFlags()
	f.StringVarP(&configFile, "config", "f",
		"etc/dousheng.toml", "dousheng api 的配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
