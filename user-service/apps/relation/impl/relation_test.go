// @Author: Ciusyan 2023/3/1
package impl_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/relation"
)

// TODO：2、完善测试用例

func TestRelationServiceImpl_FollowAction(t *testing.T) {
	should := assert.New(t)

	req := relation.NewFollowActionRequest()
	// 16 -> 17
	req.LoginUserId = 16
	req.ToUserId = 17
	req.ActionType = relation.ActionType_UN_FOLLOW_ACTION

	_, err := service.FollowAction(context.Background(), req)
	if should.NoError(err) {
		t.Log("关注成功")
	}
}

func TestRelationServiceImpl_FollowerList(t *testing.T) {

}

func TestRelationServiceImpl_FollowList(t *testing.T) {

}

func TestRelationServiceImpl_FriendList(t *testing.T) {
	should := assert.New(t)

	req := relation.NewFriendListRequest()
	// 查询 16的朋友
	req.LoginUserId = 16
	req.UserId = 25

	resp, err := service.FriendList(context.Background(), req)

	if should.NoError(err) {
		t.Log(resp)
	}
}

var (
	service relation.ServiceServer
)

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
	service = ioc.GetGrpcDependency(relation.AppName).(relation.ServiceServer)
}
