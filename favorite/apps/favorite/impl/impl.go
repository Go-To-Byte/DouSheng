// Author: BeYoung
// Date: 2023/2/22 20:37
// Software: GoLand

package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/favorite/apps/favorite"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &FavoriteServiceImpl{}

func NewFavoriteServiceImpl() *FavoriteServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &FavoriteServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("Favorite"),
		db: db,
	}
}

// FavoriteServiceImpl favoriteServiceImpl 基于Mysql实现的Service
type FavoriteServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	favorite.UnimplementedFavoriteServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (f *FavoriteServiceImpl) Init() error {
	f.l = zap.L().Named("Favorite")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	f.db = db

	client, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	f.tokenService = client.TokenService()

	return nil
}

func (f *FavoriteServiceImpl) Name() string {
	return favorite.AppName
}

func (f *FavoriteServiceImpl) Registry(s *grpc.Server) {
	favorite.RegisterFavoriteServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}