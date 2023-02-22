// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/chat/apps/chat"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"google.golang.org/grpc"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &ChatServiceImpl{}

func NewCommentServiceImpl() *ChatServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &ChatServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("Comment"),
		db: db,
	}
}

// ChatServiceImpl  基于Mysql实现的Service
type ChatServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	chat.UnimplementedChatServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (c *ChatServiceImpl) Init() error {
	c.l = zap.L().Named("Comment")

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

func (c *ChatServiceImpl) Name() string {
	return chat.AppName
}

func (c *ChatServiceImpl) Registry(s *grpc.Server) {
	chat.RegisterChatServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
