// Created by yczbest at 2023/02/19 19:01

package impl

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
)

var service favorite.ServiceServer

func TestFavoriteServiceImpl_FavoriteAction(t *testing.T) {
	should := assert.New(t)
	newFavorite := favorite.NewFavoriteActionRequest()
	newFavorite.Token = "sPDgHB87RaWwMCP1vJlDrIdG"
	newFavorite.ActionType = 1
	newFavorite.VideoId = 2
	_, err := service.FavoriteAction(context.Background(), newFavorite)

	if should.NoError(err) {
		fmt.Println("点赞成功！")
	}
}

func TestFavoriteServiceImpl_Delete(t *testing.T) {
	should := assert.New(t)
	newFavorite := favorite.NewFavoriteActionRequest()
	newFavorite.Token = "bpnAkzIpL6mUFNcn148JTdUF"
	newFavorite.ActionType = 1
	newFavorite.VideoId = 23
	_, err := service.FavoriteAction(context.Background(), newFavorite)
	if should.NoError(err) {
		fmt.Println("取消点赞成功！")
	}
}

func TestFavoriteServiceImpl_FavoriteList(t *testing.T) {
	should := assert.New(t)
	newReq := favorite.NewFavoriteListRequest()
	newReq.UserId = 17
	newReq.Token = "bpnAkzIpL6mUFNcn148JTdUF"
	res, err := service.FavoriteList(context.Background(), newReq)
	fmt.Println(err)
	if should.NoError(err) {
		fmt.Println("获取喜欢视频列表成功！")
		fmt.Println(res.VideoList)
	}
}

func TestFavoriteServiceImpl_FavoriteCount(t *testing.T) {
	should := assert.New(t)
	newReq := favorite.NewFavoriteCountRequest()
	newReq.VideoIds = []int64{17, 18, 24, 25}
	res, err := service.FavoriteCount(context.Background(), newReq)
	if should.NoError(err) {
		fmt.Println("获取点赞总数成功")
		fmt.Println(res.AcquireFavoriteCount)
		fmt.Println(res.FavoriteCount)
	}
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
