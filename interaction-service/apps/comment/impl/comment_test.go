// Created by yczbest at 2023/02/22 20:51

package impl

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
)

var service comment.ServiceServer

// 发送评论
func TestCommentServiceImpl_CommentAction(t *testing.T) {
	should := assert.New(t)
	newComment := comment.NewCommentActionRequest()
	newComment.Token = "sPDgHB87RaWwMCP1vJlDrIdG" //15
	newComment.VideoId = 2
	newComment.ActionType = 1
	newComment.CommentText = "测试数据"
	_, err := service.CommentAction(context.Background(), newComment)
	if should.NoError(err) {
		fmt.Println("评论成功！")
	}
}

func TestCommentServiceImpl_DeleteCommentById(t *testing.T) {
	should := assert.New(t)
	newComment := comment.NewCommentActionRequest()
	newComment.Token = "sPDgHB87RaWwMCP1vJlDrIdG"
	newComment.VideoId = 2
	newComment.ActionType = 2
	newComment.CommentId = 1677196402757710700

	_, err := service.CommentAction(context.Background(), newComment)
	if should.NoError(err) {
		fmt.Println("删除评论成功！")
	}
}

func TestCommentServiceImpl_GetCommentList(t *testing.T) {
	should := assert.New(t)
	newComment := comment.NewDefaultGetCommentListRequest()
	newComment.Token = "sPDgHB87RaWwMCP1vJlDrIdG"
	newComment.VideoId = 2

	rsp, err := service.GetCommentList(context.Background(), newComment)
	if should.NoError(err) {
		fmt.Println("获取评论列表成功！")
		fmt.Println(rsp.CommentList)
	}
}

func TestCommentServiceImpl_GetCommentCountById(t *testing.T) {
	should := assert.New(t)
	newComment := comment.NewDefaultGetCommentCountByIdRequest()
	newComment.VideoId = 2

	rsp, err := service.GetCommentCountById(context.Background(), newComment)
	if should.NoError(err) {
		fmt.Println("获取评论总数成功！")
		fmt.Println(rsp.CommentCount)
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
	service = ioc.GetGrpcDependency(comment.AppName).(comment.ServiceServer)
}
