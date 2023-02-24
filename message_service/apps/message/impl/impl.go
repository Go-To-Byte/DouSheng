// @Author: Hexiaoming 2023/2/7
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
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	userRpc "github.com/Go-To-Byte/DouSheng/user_center/client/rpc"

	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
)

var (
	impl = &messageServiceImpl{}
)

type messageServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	message.UnimplementedServiceServer

	// 依赖Token的客户端
	tokenService token.ServiceClient

	// 依赖User 的客户端
	userServer user.ServiceClient
}

func (s *messageServiceImpl) Init() error {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
	s.l = zap.L().Named(message.AppName)

	// 获取API网关客户端[GRPC调用]
	apiRooter, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		s.l.Errorf("message: 获取API rooter 出现错误：%s", err.Error())
		return err
	}
	s.tokenService = apiRooter.TokenService()

	// 获取用户中心的客户端[GRPC调用]
	userCenter, err := userRpc.NewUserCenterClientFromCfg()
	if err != nil {
		return err
	}
	s.userServer = userCenter.UserService()

	return nil
}

func (s *messageServiceImpl) Name() string {
	return message.AppName
}

func (s *messageServiceImpl) Registry(server *grpc.Server) {
	message.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
