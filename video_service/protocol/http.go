// @Author: Ciusyan 2023/2/7
package protocol

import (
	"github.com/gin-gonic/gin"

	"github.com/Go-To-Byte/DouSheng/dou_kit/cmd"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/dou_kit/protocol"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc"
)

// =====
// 使用HTTP对外暴露服务
// =====

// 获取路由中间件
func getMiddle() ([]gin.HandlerFunc, error) {
	// 从配置文件中获取 user_center 的Client
	client, err := rpc.NewUserCenterClientFromCfg()

	if err != nil {
		return nil, err
	}

	// 添加Token认证中间件
	middles := []gin.HandlerFunc{client.GinAuthHandlerFunc()}
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