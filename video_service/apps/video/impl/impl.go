// @Author: Ciusyan 2023/2/7
package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
)

var (
	impl = &videoServiceImpl{}
)

type videoServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	video.UnimplementedServiceServer

	// 依赖Token的客户端
	tokenService token.ServiceClient
}

func (s *videoServiceImpl) Init() error {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
	s.l = zap.L().Named(video.AppName)

	// 获取用户中心的客户端[GRPC调用]
	apiRooter, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		s.l.Errorf("video: getVideoPo 出现错误：%s", err.Error())
		return err
	}
	s.tokenService = apiRooter.TokenService()
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
