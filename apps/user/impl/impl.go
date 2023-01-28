// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &userServiceImpl{}

func NewUserServiceImpl() *userServiceImpl {
	return &userServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("User"),
		db: conf.C().MySQL.GetDB(),
	}
}

// userServiceImpl 基于Mysql实现的Service
type userServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user.UnimplementedServiceServer
}

func (u *userServiceImpl) Init() error {
	u.l = zap.L().Named("User")
	u.db = conf.C().MySQL.GetDB()

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
