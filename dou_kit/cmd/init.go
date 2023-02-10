package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
)

var (
	createTableFilePath string
)

// initCmd represents the start command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "DouSheng 服务初始化",
	Long:  "DouSheng 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := conf.LoadConfigFromToml(configFile); err != nil {
			return err
		}

		err := createTables()
		if err != nil {
			return err
		}

		return nil
	},
}

func createTables() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 读取SQL文件
	sqlFile, err := os.ReadFile(createTableFilePath)
	if err != nil {
		return err
	}

	fmt.Println("执行的SQL: ")
	fmt.Println(string(sqlFile))

	// 执行SQL文件
	res := db.WithContext(ctx).Exec(string(sqlFile))
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func init() {
	initCmd.PersistentFlags().StringVarP(&createTableFilePath, "sql-file-path", "s", "docs/sql/tables.sql", "SQL脚本路径")
	RootCmd.AddCommand(initCmd)
}
