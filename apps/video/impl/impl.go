// @Author: Ciusyan 2023/1/29
package impl

import (
	"github.com/Go-To-Byte/DouSheng/apps/video"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	impl = &videoServiceImpl{}
)

type videoServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	video.UnimplementedServiceServer
}

func (s *videoServiceImpl) Init() error {
	s.db = conf.C().MySQL.GetDB()
	s.l = zap.L().Named(video.AppName)

	return nil
}

func (s *videoServiceImpl) Name() string {
	return video.AppName
}

func (s *videoServiceImpl) Registry(server *grpc.Server) {
	video.RegisterServiceServer(server, impl)
}

func init() {
	// 注入IOC容器
	ioc.GrpcDI(impl)
}
