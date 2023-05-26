// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"sync"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api-rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/favorite"
	interactionrpc "github.com/Go-To-Byte/DouSheng/interaction-service/client/rpc"
	"github.com/Go-To-Byte/DouSheng/relation-service/apps/relation"
	relationrpc "github.com/Go-To-Byte/DouSheng/relation-service/client/rpc"
	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
	videoRpc "github.com/Go-To-Byte/DouSheng/video-service/client/rpc"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
)

// 用于注入IOC中
var impl = &userServiceImpl{}

// userServiceImpl 基于Mysql实现的Service
type userServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user.UnimplementedServiceServer
	// 用于管理Token
	video        video.ServiceClient
	relation     relation.ServiceClient
	favorite     favorite.ServiceClient
	tokenService token.ServiceClient
}

func (u *userServiceImpl) Init() error {
	u.l = zap.L().Named("User")
	errors := make([]error, 0)

	wait := sync.WaitGroup{}
	wait.Add(1)
	defer wait.Wait()
	go func() {
		defer wait.Done()
		if db, err := conf.C().MySQL.GetDB(); err != nil {
			errors = append(errors, err)
		} else {
			u.db = db
		}
	}()

	go func() {
		wait.Add(1)
		defer wait.Done()
		if client, err := rpc.NewApiRooterClientFromCfg(); err != nil {
			errors = append(errors, err)
		} else {
			u.tokenService = client.TokenService()
		}
	}()

	go func() {
		wait.Add(1)
		defer wait.Done()
		if client, err := videoRpc.NewVideoServiceClientFromCfg(); err != nil {
			errors = append(errors, err)
		} else {
			u.video = client.VideoService()
		}
	}()

	go func() {
		wait.Add(1)
		defer wait.Done()
		if client, err := relationrpc.NewRelationServiceClientFromCfg(); err != nil {
			errors = append(errors, err)
		} else {
			u.relation = client.RelationService()
		}
	}()

	go func() {
		wait.Add(1)
		defer wait.Done()
		if client, err := interactionrpc.NewInteractionServiceClientFromConfig(); err != nil {
			errors = append(errors, err)
		} else {
			u.favorite = client.FavoriteService()
		}
	}()

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
