// @Author: Ciusyan 2023/1/24
package cmd

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/protocol"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "启动 API服务",
	Short:   "启动 API服务",
	Example: "go run main start",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ========
		// 1、加载配置文件&全局Logger对象
		// ========

		if err := conf.LoadConfigFromToml(configFile); err != nil {
			return err
		}
		if err := conf.LoadGlobalLogger(); err != nil {
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
		return managerStartAndStop()
	},
}

var (
	// HttpStartAop 使用者可以传入一个切面，不传就使用默认值
	HttpStartAop = protocol.DefaultHttpStartBefore()
)

func NewManager() *manager {
	return &manager{
		http: protocol.NewHttpService(HttpStartAop),
		grpc: protocol.NewGRPCService(),
		l:    zap.L().Named("CLI"),
	}
}

// 用于管理服务的开启、和关闭
type manager struct {
	http *protocol.HttpService
	grpc *protocol.GRPCService
	l    logger.Logger
}

func (m *manager) start() error {

	// 打印加载好的服务
	m.l.Infof("已加载的 [Internal] 服务: %s", ioc.ExistingInternalDependencies())
	m.l.Infof("已加载的 [GRPC] 服务: %s", ioc.ExistingGrpcDependencies())
	m.l.Infof("已加载的 [HTTP] 服务: %s", ioc.ExistingGinDependencies())

	// 将GRPC放在后台跑
	go m.grpc.Start()

	// 注：这属于正常关闭："api: Server closed"
	if err := m.http.Start(); err != nil && err.Error() != "api: Server closed" {
		return err
	}
	return nil
}

// WaitStop 处理来自外部的中断信号，比如Terminal
func (m *manager) waitStop(ch <-chan os.Signal) {

	for v := range ch {
		switch v {
		default:
			m.l.Infof("接受到信号：%s", v)

			if err := m.http.Stop(); err != nil {
				m.l.Errorf("优雅关闭 [HTTP] 服务出错：%s", err.Error())
			}

			if err := m.grpc.Stop(); err != nil {
				m.l.Errorf("优雅关闭 [GRPC] 服务出错：%s", err.Error())
			}
		}
	}
}

// WaitSign 等待退出的信号，实现优雅退出
func (m *manager) waitSign() {
	// 用于接收信号的信道
	ch := make(chan os.Signal, 1)
	// 接收这几种信号
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)

	// 需要在后台等待关闭
	go m.waitStop(ch)
}

// 利用管理者来管理程序的启动和关闭
func managerStartAndStop() error {
	m := NewManager()
	m.waitSign()
	return m.start() // 开启服务
}

func init() {
	f := StartCmd.PersistentFlags()
	f.StringVarP(&configFile, "config", "f",
		"etc/config.toml", "用户中心的配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
