// @Author: Ciusyan 2023/2/7
package conf_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/Go-To-Byte/DouSheng/video_service/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/config.toml")

	if should.NoError(err) {
		should.Equal("DouSheng", conf.C().App.Name)
	}

}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")

	err := conf.LoadConfigFromEnv()

	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
	}

}

func TestMySQLGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if should.NoError(err) {
		fmt.Println(conf.C().MySQL.GetDB())
	}
}
