// @Author: Ciusyan 2023/2/7
package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou_common/ioc"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/Go-To-Byte/DouSheng/video_service/conf"
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

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
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
	ioc.GrpcDI(impl)
}
