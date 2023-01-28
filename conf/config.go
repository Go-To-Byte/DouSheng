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
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

func NewDefaultApp() *app {
	return &app{
		Name: "dousheng",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func (a *app) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "",
		Password:    "",
		Database:    "",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

func (m *MySQL) GetDB() *gorm.DB {
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
func (m *MySQL) getDBConn() (*gorm.DB, error) {
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

// MySQL todo
type MySQL struct {
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

func NewDefaultLog() *Log {
	return &Log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}
