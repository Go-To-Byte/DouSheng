// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user/apps/user"
	"google.golang.org/grpc"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &UserServiceImpl{}

func NewCommentServiceImpl() *UserServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &UserServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("Comment"),
		db: db,
	}
}

// UserServiceImpl  基于Mysql实现的Service
type UserServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user.UnimplementedUserServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (s *UserServiceImpl) Init() error {
	s.l = zap.L().Named("Comment")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db

	client, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	s.tokenService = client.TokenService()

	return nil
}

func (s *UserServiceImpl) Name() string {
	return user.AppName
}

func (s *UserServiceImpl) Registry(srv *grpc.Server) {
	user.RegisterUserServer(srv, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
