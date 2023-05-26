// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	tkRpc "github.com/Go-To-Byte/DouSheng/api-rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
	userRpc "github.com/Go-To-Byte/DouSheng/user-service/client/rpc"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
)

type commentServiceImpl struct {
	db *gorm.DB
	l  logger.Logger
	comment.UnimplementedServiceServer
	tokenService token.ServiceClient
	userService  user.ServiceClient
}

// 定义全局服务组件
var impl = &commentServiceImpl{}

func NewCommentServiceImpl() *commentServiceImpl {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}
	return &commentServiceImpl{
		//Interaction的子模块Comment
		l:  zap.L().Named("Comment"),
		db: db,
	}
}

// 对封装内容的初始化
func (c *commentServiceImpl) Init() error {
	c.l = zap.L().Named("Comment")
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}
	c.db = db

	tkClient, err := tkRpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}
	c.tokenService = tkClient.TokenService()

	userCilent, err := userRpc.NewUserCenterClientFromCfg()
	if err != nil {
		return err
	}
	c.userService = userCilent.UserService()

	return nil
}

// 服务名
func (c *commentServiceImpl) Name() string {
	return comment.AppName
}

// 注册PPC服务
func (c *commentServiceImpl) Registry(s *grpc.Server) {
	comment.RegisterServiceServer(s, impl)
}

func init() {
	//注入IOC
	ioc.GrpcDI(impl)
}
