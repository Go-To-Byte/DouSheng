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

func (c *FavoriteServiceImpl) Init() error {
	c.l = zap.L().Named("Favorite")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	c.db = db

	client, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	c.tokenService = client.TokenService()

	return nil
}

func (c *FavoriteServiceImpl) Name() string {
	return favorite.AppName
}

func (c *FavoriteServiceImpl) Registry(s *grpc.Server) {
	favorite.RegisterFavoriteServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
