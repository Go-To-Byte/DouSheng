// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &userServiceImpl{}

// userServiceImpl 基于Mysql实现的Service
type userServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user.UnimplementedServiceServer
}

func (u *userServiceImpl) Init() error {
	u.l = zap.L().Named("User")

	// 从配置中获取 Mysql 驱动
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	u.db = db

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
