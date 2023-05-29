// @Author: Ciusyan 2023/2/7

package protocol

import (
	"github.com/gin-gonic/gin"

	authRpc "github.com/Go-To-Byte/DouSheng/auth-service/client/rpc"
	"github.com/Go-To-Byte/DouSheng/dou-kit/cmd"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/Go-To-Byte/DouSheng/dou-kit/ioc"
	"github.com/Go-To-Byte/DouSheng/dou-kit/protocol"
)

// 获取路由中间件
func getMiddle() ([]gin.HandlerFunc, error) {

	// 获取认证服务客户端
	auth, err := authRpc.NewAuthServiceClientFromCfg()
	if err != nil {
		// 这里可以直接 panic，相当于认证服务都加载不到

		panic(err)
	}

	// 添加Token认证中间件
	middles := []gin.HandlerFunc{exception.GinErrWrapper(auth.GinAuthHandlerFunc())}

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
