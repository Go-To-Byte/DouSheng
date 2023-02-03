// Author: BeYoung
// Date: 2023/2/1 0:24
// Software: GoLand

package models

type ConfigYAML struct {
	ID       int64        `mapstructure:"ID"`
	DB       dbConfig     `mapstructure:"sql"`
	Jwt      jwtConfig    `mapstructure:"jwt"`
	GrpcName grpcConfig   `mapstructure:"grpc"`
	Consul   consulConfig `mapstructure:"consul"`
}

type jwtConfig struct {
	Key string `mapstructure:"key"`
}

type consulConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type dbConfig struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type grpcConfig struct {
	Feed     string `mapstructure:"feed"`
	User     string `mapstructure:"user"`
	Video    string `mapstructure:"video"`
	Comment  string `mapstructure:"comment"`
	Message  string `mapstructure:"message"`
	Relation string `mapstructure:"relation"`
	Favorite string `mapstructure:"favorite"`
}
