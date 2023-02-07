// @Author: Ciusyan 2023/2/7
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/Go-To-Byte/DouSheng/video_service/version"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "video",
	Long:    "视频服务",
	Short:   "视频服务",
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
	f.BoolVarP(&vers, "version", "v", false, "视频服务的版本信息")
}
