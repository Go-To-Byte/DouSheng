// Created by yczbest at 2023/02/19 19:01

package impl

import (
	"github.com/infraboard/mcube/logger/zap"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"

	"github.com/Go-To-Byte/DouSheng/video-service/apps/favorite"
)

var service favorite.ServiceServer

func TestFavoriteServiceImpl_FavoriteAction(t *testing.T) {

}

func TestFavoriteServiceImpl_Delete(t *testing.T) {

}

func TestFavoriteServiceImpl_FavoriteList(t *testing.T) {

}

func init() {

	// 加载配置文件
	if err := conf.LoadConfigFromToml("../../../etc/config.toml"); err != nil {
		panic(err)
	}

	// 初始化全局Logger
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	// 初始化IOC容器
	if err := ioc.InitAllDependencies(); err != nil {
		panic(err)
	}

	// 从IOC中获取接口实现
	service = ioc.GetGrpcDependency(favorite.AppName).(favorite.ServiceServer)
}
