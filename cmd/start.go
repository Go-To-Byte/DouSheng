// @Author: Ciusyan 2023/1/24
package cmd

import (
	"github.com/Go-To-Byte/DouSheng/apps"
	_ "github.com/Go-To-Byte/DouSheng/apps/all"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "启动 Dousheng API服务",
	Short:   "启动 Dousheng API服务",
	Example: "go run main start",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1、加载配置文件

		if err := conf.LoadConfigFromToml(configFile); err != nil {
			panic(err)
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
