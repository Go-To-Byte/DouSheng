// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
)

// 用于注入IOC中
var impl = &userServiceImpl{}

func NewUserServiceImpl() *userServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &userServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("User"),
		db: db,
	}
}

// userServiceImpl 基于Mysql实现的Service
type userServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user.UnimplementedServiceServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (u *userServiceImpl) Init() error {
	u.l = zap.L().Named("User")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	u.db = db

	client, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	u.tokenService = client.TokenService()

	return nil
}

func (u *userServiceImpl) Name() string {
	return user.AppName
}

func (u *userServiceImpl) Registry(s *grpc.Server) {
	user.RegisterServiceServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
