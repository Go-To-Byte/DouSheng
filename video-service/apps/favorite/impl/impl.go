// Created by yczbest at 2023/02/18 15:00

package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
)

// 定义全局组件
var impl = &favoriteServiceImpl{}

// 定义结构
type favoriteServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	favorite.UnimplementedServiceServer

	// 视频模块
	videoService video.ServiceServer
}

func (s *favoriteServiceImpl) Init() error {
	s.l = zap.L().Named("Favorite")

	// 获取 MySQL 驱动
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.db = db

	// 注入 视频模块 的依赖
	s.videoService = ioc.GetGrpcDependency(video.AppName).(video.ServiceServer)

	return nil
}

func (u *favoriteServiceImpl) Name() string {
	return favorite.AppName
}

func (f *favoriteServiceImpl) Registry(s *grpc.Server) {
	favorite.RegisterServiceServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
