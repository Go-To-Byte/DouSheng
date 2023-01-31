// Author: BeYoung
// Date: 2023/2/1 0:24
// Software: GoLand

package models

type ConfigYAML struct {
	ID         int64      `mapstructure:"ID"`
	GrpcConfig GrpcConfig `mapstructure:"grpc"`
	DBConfig   DBConfig   `mapstructure:"sql"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type GrpcConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
