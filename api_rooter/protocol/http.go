// @Author: Ciusyan 2023/2/7
package protocol

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/cmd"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/protocol"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/common/utils"
)

// =====
// 使用HTTP对外暴露服务
// =====

// Auther 因为认证放在了此模块内部，所以直接从Ioc中获取依赖即可
type Auther struct {
	l logger.Logger

	t token.ServiceServer
}

func NewAuther() *Auther {
	return &Auther{
		l: zap.L().Named("Token_Server"),

		// 从IOC中取出Server端，因为是在自己的模块内部
		t: ioc.GetGrpcDependency(token.AppName).(token.ServiceServer),
	}
}

// GinAuthHandlerFunc HTTP auth中间件
func (a *Auther) GinAuthHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 从请求中解析出Token
		ak := utils.GetToken(ctx)

		// 验证Token
		req := token.NewValidateTokenRequest(ak)
		tk, err := a.t.ValidateToken(ctx.Request.Context(), req)

		if err != nil {
			a.l.Errorf("Token认证失败：%s", err.Error())
			ctx.JSON(http.StatusBadRequest, constant.ERROR_TOKEN_VALIDATE)
			// 有错误、直接终止传递
			ctx.Abort()
		} else {
			a.l.Infof("Token认证成功")
		}

		// 把Token传递给下一个链路
		ctx.Set(constant.REQUEST_TOKEN, tk)
		// 把请求传递下去
		ctx.Next()
	}
}

// 获取路由中间件
func getMiddle() ([]gin.HandlerFunc, error) {
	server := NewAuther()
	// 添加Token认证中间件
	middles := []gin.HandlerFunc{server.GinAuthHandlerFunc()}
	return middles, nil
}

func init() {
	// 给需要使用的切面赋值
	cmd.HttpStartAop = func(s *protocol.HttpService) error {
		// 1、获取中间件
		mids, err := getMiddle()
		if err != nil {
			return err
		}
		// 2、将所有的Gin服务对象注册到IOC中
		option := ioc.NewGinOption(s.R, "/"+s.C.App.Name, mids...)
		option.NotVersion = true
		option.NotName = true
		ioc.RegistryGin(option)
		return nil
	}
}
