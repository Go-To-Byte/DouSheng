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
var impl = &UserServiceImpl{}

func NewUserServiceImpl() *UserServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &UserServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("User"),
		db: db,
	}
}

// UserServiceImpl 基于Mysql实现的Service
type UserServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user.UnimplementedServiceServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (u *UserServiceImpl) Init() error {
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

func (u *UserServiceImpl) Name() string {
	return user.AppName
}

func (u *UserServiceImpl) Registry(s *grpc.Server) {
	user.RegisterServiceServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
