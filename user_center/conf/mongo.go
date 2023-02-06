// @Author: Ciusyan 2023/2/6
package conf

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

//=====
// Mysql配置对象
//=====

type mongodb struct {
	Hosts    []string `toml:"hosts" env:"MONGO_HOSTS"`
	Username string   `toml:"username" env:"MONGO_USERNAME"`
	Password string   `toml:"password" env:"MONGO_PASSWORD"`
	Database string   `toml:"database" env:"MONGO_DATABASE"`

	lock sync.Mutex
}

func NewMongodb() *mongodb {
	return &mongodb{
		Hosts:    []string{"127.0.0.1:27017"},
		Database: "",
	}
}

func (m *mongodb) GetDB() *mongo.Database {
	client, err := m.MongoClient()
	if err != nil {
		panic(err)
	}
	// 根据mongo客户端获取具体的数据库连接
	return client.Database(m.Database)
}

var mongoClient *mongo.Client

// MongoClient 获取mongo客户端连接
func (m *mongodb) MongoClient() (*mongo.Client, error) {
	// 保证线程安全
	m.lock.Lock()
	defer m.lock.Unlock()

	if mongoClient == nil {
		client, err := m.getClient()
		if err != nil {
			return nil, err
		}
		mongoClient = client
	}

	return mongoClient, nil
}

// 连接mongo客户端
func (m *mongodb) getClient() (*mongo.Client, error) {
	opt := options.Client()

	// AuthSource：代表认证数据库，mongo的用户是针对DB的，认证用户与对应的数据库一起创建
	credential := options.Credential{
		AuthSource: m.Database,
	}

	// PasswordSet：使用 password 认证
	if m.Username != "" && m.Password != "" {
		credential.Username = m.Username
		credential.Password = m.Password
		credential.PasswordSet = true
		opt.SetAuth(credential)
	}

	// Mongo 地址
	opt.SetHosts(m.Hosts)
	opt.SetConnectTimeout(5 * time.Second)

	// 注：mongo 这里是惰性连接
	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return nil, fmt.Errorf("新建mongo连接失败：%s", err.Error())
	}

	// 查看是否真正连接上了
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("ping mongodb 服务(%s) 失败： %s", m.Hosts, err)
	}

	return client, nil
}
