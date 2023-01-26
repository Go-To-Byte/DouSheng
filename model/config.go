// Author: BeYoung
// Date: 2023/1/27 5:33
// Software: GoLand

package model

type Config struct {
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
