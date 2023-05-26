// @Author: Hexiaoming 2023/2/7
package impl

import (
	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user-service/apps/relation"
)

var (
	impl = &relationServiceImpl{}
)

type relationServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	relation.UnimplementedServiceServer

	// 引入用户模块（引入内部服务）
	userServer user.ServiceServer
}

func (s *relationServiceImpl) Init() error {
	s.l = zap.L().Named(relation.AppName)

	// 获取MySQL 驱动
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.db = db

	// 注入用户服务的依赖
	s.userServer = ioc.GetGrpcDependency(user.AppName).(user.ServiceServer)

	return nil
}

func (s *relationServiceImpl) Name() string {
	return relation.AppName
}

func (s *relationServiceImpl) Registry(server *grpc.Server) {
	relation.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
