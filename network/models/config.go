// Author: BeYoung
// Date: 2023/2/1 0:24
// Software: GoLand

package models

type ConfigYAML struct {
	ID         int64      `mapstructure:"ID"`
	JwtConfig  jwtConfig  `mapstructure:"jwt"`
	GrpcConfig grpcConfig `mapstructure:"grpc"`
	DBConfig   dbConfig   `mapstructure:"sql"`
}

type jwtConfig struct {
	Key string `mapstructure:"key"`
}

type dbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type grpcConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
