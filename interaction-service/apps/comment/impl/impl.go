// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
)

type commentServiceImpl struct {
	db *gorm.DB
	l  logger.Logger
	comment.UnimplementedServiceServer
}

// 定义全局服务组件
var impl = &commentServiceImpl{}

func (c *commentServiceImpl) Init() error {
	c.l = zap.L().Named(comment.AppName)

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}
	c.db = db

	return nil
}

func (c *commentServiceImpl) Name() string {
	return comment.AppName
}

func (c *commentServiceImpl) Registry(s *grpc.Server) {
	comment.RegisterServiceServer(s, impl)
}

func init() {
	//注入IOC
	ioc.GrpcDI(impl)
}
