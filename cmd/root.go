// @Author: Ciusyan 2023/1/24
package cmd

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/version"
	"github.com/spf13/cobra"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "Dousheng",
	Long:    "dousheng API",
	Short:   "dusheng API",
	Example: "dousheng cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
		}
		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&vers, "version", "v", false, "Dousheng的版本信息")
}
