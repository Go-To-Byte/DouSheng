// Author: BeYoung
// Date: 2023/1/28 2:04
// Software: GoLand

package models

type ConfigYAML struct {
	ID       int64    `mapstructure:"ID"`
	DBConfig DBConfig `mapstructure:"sql"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}
