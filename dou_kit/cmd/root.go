// @Author: Ciusyan 2023/1/24
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/Go-To-Byte/DouSheng/dou_kit/version"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "user",
	Long:    "用户中心",
	Short:   "用户中心",
	Example: "go run main.go start",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
		}
		return cmd.Help()
	},
}

// Main 程序启动交由CLI
func Main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&vers, "version", "v", false, "用户中心的版本信息")
}
