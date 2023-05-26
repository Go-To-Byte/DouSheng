// @Author: Ciusyan 2023/2/7
package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	apiRpc "github.com/Go-To-Byte/DouSheng/api-rooter/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/favorite"
	interactionRpc "github.com/Go-To-Byte/DouSheng/interaction-service/client/rpc"
	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"
	userRpc "github.com/Go-To-Byte/DouSheng/user-service/client/rpc"

	"github.com/Go-To-Byte/DouSheng/video-service/apps/video"
)

var (
	impl = &videoServiceImpl{}
)

type videoServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	video.UnimplementedServiceServer

	// 依赖Token的客户端
	tokenService token.ServiceClient

	// 依赖User 的客户端
	userServer user.ServiceClient

	// 依赖 Favorite 的客户端
	favoriteService favorite.ServiceClient

	// 依赖 Comment 的客户端
	commentService comment.ServiceClient
}

func (s *videoServiceImpl) Init() error {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
	s.l = zap.L().Named(video.AppName)

	// 获取ApiRooter的客户端[GRPC调用]
	apiRooter, err := apiRpc.NewApiRooterClientFromCfg()
	if err != nil {
		return err
	}

	s.tokenService = apiRooter.TokenService()

	// 获取用户中心的客户端[GRPC调用]
	userCenter, err := userRpc.NewUserCenterClientFromCfg()
	if err != nil {
		return err
	}

	s.userServer = userCenter.UserService()

	// 获取用户中心的客户端[GRPC调用]
	favoriteCenter, err := interactionRpc.NewInteractionServiceClientFromConfig()
	if err != nil {
		return err
	}

	s.favoriteService = favoriteCenter.FavoriteService()
	s.commentService = favoriteCenter.CommentService()

	return nil
}

func (s *videoServiceImpl) Name() string {
	return video.AppName
}

func (s *videoServiceImpl) Registry(server *grpc.Server) {
	video.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
