// Author: BeYoung
// Date: 2023/1/28 2:04
// Software: GoLand

package models

type ConfigYAML struct {
	ID        int64           `mapstructure:"ID"`
	DB        dbConfig        `mapstructure:"sql"`
	Consul    consulConfig    `mapstructure:"consul"`
	Localhost localhostConfig `mapstructure:"localhost"`
}

type dbConfig struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type consulConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Tags string `mapstructure:"tags"`
}

type localhostConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
