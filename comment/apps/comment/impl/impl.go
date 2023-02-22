// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"google.golang.org/grpc"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &CommentServiceImpl{}

func NewUserServiceImpl() *CommentServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &CommentServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("User"),
		db: db,
	}
}

// CommentServiceImpl userServiceImpl 基于Mysql实现的Service
type CommentServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	comment.UnimplementedCommentServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (c *CommentServiceImpl) Init() error {
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

func (c *CommentServiceImpl) Name() string {
	return comment.AppName
}

func (c *CommentServiceImpl) Registry(s *grpc.Server) {
	comment.RegisterCommentServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
