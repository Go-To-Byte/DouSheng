// @Author: Ciusyan 2023/1/23
package conf

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

// 防止配置文件在运行时被更改，设置为私有的
var config *Config

func C() *Config {
	return config
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

// Config 将配置文件抽成一个对象
type Config struct {
	App   *app   `toml:"app"`
	Log   *log   `toml:"log"`
	MySQL *mySQL `toml:"mysql"`
}

func NewDefaultApp() *app {
	return &app{
		Name: "dousheng",
		HTTP: newDefaultHTTP(),
		GRPC: newDefaultGRPC(),
	}
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	HTTP *http  `toml:"api"`
	GRPC *grpc  `toml:"grpc"`
}

func NewDefaultMySQL() *mySQL {
	return &mySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "",
		Password:    "",
		Database:    "",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock() // 锁住临界区，保证线程安全
	defer m.lock.Unlock()

	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}

	return db
}

// gorm获取数据库连接
func (m *mySQL) getDBConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接Mysql：%s，error：%s", dsn, err.Error())
	}

	// 维护连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)
	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))

	// 用于测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql %s，error：%s", dsn, err.Error())
	}
	return db, nil
}

var db *gorm.DB

// mySQL mysql配置
type mySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的是 Mysql的连接池，需要对连接池做一些规划配置
	// 控制当前程序的 Mysql打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制 Mysql复用，比如 5， 最多运行5个复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期，这个和 Mysql Server配置有关系，必须小于 Server 配置
	// 比如一个链接用 12 h 换一个 conn，保证一点的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle 连接 最多允许存货多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	// 作为私有变量，用于控制DetDB
	lock sync.Mutex
}

func NewDefaultLog() *log {
	return &log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// log 日志配置
type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8050",
	}
}

// HTTP 服务配置
type http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port string `toml:"port" env:"HTTP_PORT"`
}

// Addr 获取 HTTP 服务配置的 IP + 端口
func (h *http) Addr() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}

func newDefaultGRPC() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: "8505",
	}
}

// GRPC 服务配置
type grpc struct {
	Host string `toml:"host" env:"GRPC_HOST"`
	Port string `toml:"port" env:"GRPC_PORT"`
}

// Addr 获取 GRPC 服务配置的 IP + 端口
func (g *grpc) Addr() string {
	return fmt.Sprintf("%s:%s", g.Host, g.Port)
}
