// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/favorite"
	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/video_service/apps/video"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"sync"
)

// 用于注入IOC中
var impl = &userServiceImpl{}

// userServiceImpl 基于Mysql实现的Service
type userServiceImpl struct {
	l  logger.Logger
	db *gorm.DB

	user     user.UnimplementedServiceServer
	video    video.UnimplementedServiceServer
	relation relation.UnimplementedServiceServer
	favorite favorite.UnimplementedServiceServer
	// 用于管理Token
	tokenService token.ServiceClient
}

func (u *userServiceImpl) Init() error {
	u.l = zap.L().Named("User")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	u.db = db

	client, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	u.tokenService = client.TokenService()

	return nil
}

func (u *userServiceImpl) Name() string {
	return user.AppName
}

func (u *userServiceImpl) Registry(s *grpc.Server) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	wait := sync.WaitGroup{}
	wait.Add(4)
	go func() {
		defer wait.Done()
		user.RegisterServiceServer(s, impl.user)
	}()
	go func() {
		defer wait.Done()
		video.RegisterServiceServer(s, impl.video)
	}()
	go func() {
		defer wait.Done()
		relation.RegisterServiceServer(s, impl.relation)
	}()
	go func() {
		defer wait.Done()
		favorite.RegisterServiceServer(s, impl.favorite)
	}()
	wait.Wait()
}

func init() {
	// 将此UserService注入GRPC服务的IOC容器中
	ioc.GrpcDI(impl)
}
