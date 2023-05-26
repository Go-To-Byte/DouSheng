// @Author: Hexiaoming 2023/2/7
package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/message"
)

var (
	impl = &messageServiceImpl{}
)

type messageServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	message.UnimplementedServiceServer
}

func (s *messageServiceImpl) Init() error {
	s.l = zap.L().Named(message.AppName)

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *messageServiceImpl) Name() string {
	return message.AppName
}

func (s *messageServiceImpl) Registry(server *grpc.Server) {
	message.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
