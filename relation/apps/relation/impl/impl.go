// Author: BeYoung
// Date: 2023/2/22 21:43
// Software: GoLand

package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/relation/apps/relation"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// 用于注入IOC中
var impl = &RelationServiceImpl{}

func NewRelationServiceImpl() *RelationServiceImpl {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		panic(err)
	}

	return &RelationServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("Relation"),
		db: db,
	}
}

// RelationServiceImpl  基于Mysql实现的Service
type RelationServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	relation.UnimplementedRelationServer

	// 用于管理Token
	tokenService token.ServiceClient
}

func (r *RelationServiceImpl) Init() error {
	r.l = zap.L().Named("Relation")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	r.db = db

	client, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	r.tokenService = client.TokenService()

	return nil
}

func (r *RelationServiceImpl) Name() string {
	return relation.AppName
}

func (r *RelationServiceImpl) Registry(s *grpc.Server) {
	relation.RegisterRelationServer(s, impl)
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
