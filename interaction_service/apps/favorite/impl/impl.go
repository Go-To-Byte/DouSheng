// Created by yczbest at 2023/02/18 15:00

package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	apiRpc "github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	userRpc "github.com/Go-To-Byte/DouSheng/user_center/client/rpc"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	videoRpc "github.com/Go-To-Byte/DouSheng/video_service/client/rpc"

	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
)

// 定义全局组件
var impl = &favoriteServiceImpl{}

// 定义结构
type favoriteServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	favorite.UnimplementedServiceServer
	//token
	tokenService token.ServiceClient
	userService  user.ServiceClient
	videoService video.ServiceClient
}

func NewFavoriteServiceImpl() *favoriteServiceImpl {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &favoriteServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("Favorite"),
		db: db,
	}
}

func (f *favoriteServiceImpl) Init() error {
	f.l = zap.L().Named("Favorite")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	f.db = db

	tkClient, err := apiRpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	f.tokenService = tkClient.TokenService()

	userClient, err := userRpc.NewUserCenterClientFromCfg()
	if err != nil {
		return err
	}

	f.userService = userClient.UserService()

	videoClient, err := videoRpc.NewVideoServiceClientFromCfg()
	if err != nil {
		return err
	}

	f.videoService = videoClient.VideoService()

	return nil
}

// 注入服务名
func (u *favoriteServiceImpl) Name() string {
	return favorite.AppName
}

// 注册GRPC服务
func (f *favoriteServiceImpl) Registry(s *grpc.Server) {
	favorite.RegisterServiceServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
