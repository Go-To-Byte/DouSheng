// @Author: Ciusyan 2023/1/29
package impl

import (
	"github.com/Go-To-Byte/DouSheng/apps/comment"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	impl = &commentServiceImpl{}
)

type commentServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	comment.UnimplementedServiceServer
}

func (s *commentServiceImpl) Init() error {
	s.db = conf.C().MySQL.GetDB()
	s.l = zap.L().Named(comment.AppName)
	return nil
}

func (s *commentServiceImpl) Name() string {
	return comment.AppName
}

func (s *commentServiceImpl) Registry(server *grpc.Server) {
	comment.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
