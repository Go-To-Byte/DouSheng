// @Author: Ciusyan 2023/1/24
package cmd

import (
	"github.com/Go-To-Byte/DouSheng/apps/user/http"
	"github.com/Go-To-Byte/DouSheng/apps/user/impl"
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

		// 2、加载Host Service 的实现类

		service := impl.NewUserServiceImpl()

		// 3、通过Gin启动服务
		api := http.NewUserHttpHandler(service)
		g := gin.Default()
		api.Registry(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	f := StartCmd.PersistentFlags()
	f.StringVarP(&configFile, "config", "f",
		"etc/dousheng.toml", "dousheng api 的配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
