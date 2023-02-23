// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/video/apps/video"
	"google.golang.org/grpc"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &VideoServiceImpl{}

func NewCommentServiceImpl() *VideoServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &VideoServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named(video.AppName),
		db: db,
	}
}

// VideoServiceImpl  基于Mysql实现的Service
type VideoServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	video.UnimplementedVideoServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (c *VideoServiceImpl) Init() error {
	c.l = zap.L().Named(video.AppName)

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

func (c *VideoServiceImpl) Name() string {
	return video.AppName
}

func (c *VideoServiceImpl) Registry(s *grpc.Server) {
	video.RegisterVideoServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
