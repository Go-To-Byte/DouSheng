// @Author: Ciusyan 2023/1/23
package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

/** 用于加载全局配置 **/

// LoadConfigFromToml 从Toml配置文件加载
func LoadConfigFromToml(filePath string) error {
	// 初始化全局对象
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config file error，path：%s，%s", filePath, err)
	}

	return nil
}

// LoadConfigFromEnv 从环境变量加载
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}

// LoadGlobal 或者可以这样加载全局实例s
func LoadGlobal() error {
	var err error
	db, err = config.MySQL.getDBConn()
	return err
}
