// @Author: Ciusyan 2023/1/24
package main

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
