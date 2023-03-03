// @Author: Ciusyan 2023/3/1
package impl_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
)

// TODO：完善测试用例

func TestRelationServiceImpl_FollowAction(t *testing.T) {
	should := assert.New(t)

	req := relation.NewFollowActionRequest()
	// 16 -> 17
	req.Token = "nuo7A79Qp4Ms7144BAyQwW4H"
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
	req.Token = "nuo7A79Qp4Ms7144BAyQwW4H"
	req.UserId = 16

	resp, err := service.FriendList(context.Background(), req)

	if should.NoError(err) {
		t.Log(resp)
	}
}

func TestRelationServiceImpl_ListCount(t *testing.T) {
	should := assert.New(t)

	req := relation.NewListCountRequest()
	req.UserId = 1
	req.Type = relation.CountType_ALL

	resp, err := service.ListCount(context.Background(), req)

	if should.NoError(err) {
		t.Log("关注数 = ", resp.FollowCount)
		t.Log("========", resp)
		t.Log("粉丝数 = ", resp.FollowerCount)
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
