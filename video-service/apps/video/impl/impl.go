// @Author: Ciusyan 2023/2/7
package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
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
	s.l = zap.L().Named(video.AppName)

	// 获取MySQL 驱动
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *videoServiceImpl) Name() string {
	return video.AppName
}

func (s *videoServiceImpl) Registry(server *grpc.Server) {
	video.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
