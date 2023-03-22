// @Author: Ciusyan 2023/1/25
package protocol

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
)

// =====
// 使用HTTP对外暴露服务
// =====

type HttpService struct {
	Server *server.Hertz
	L      logger.Logger // 用于打印日志
	C      *conf.Config  // 用于获取项目名称

	before StartFuncAop // 执行start前的逻辑
}

func NewHttpService(f StartFuncAop) *HttpService {

	service := &HttpService{
		L:      zap.L().Named("Http Service"),
		C:      conf.C(),
		before: f,
	}

	// HTTP服务对象

	r := server.Default(
		server.WithHostPorts(service.C.App.HTTP.Addr()),
		server.WithReadTimeout(60*time.Second),
		server.WithWriteTimeout(60*time.Second),
		server.WithIdleTimeout(60*time.Second),
		server.WithMaxRequestBodySize(1<<20),
		// 配置地址：https://www.cloudwego.io/zh/docs/hertz/reference/config/

	)
	// 添加日志中间件
	r.Use(AccessLog())

	service.Server = r

	return service
}

// Start 开启服务
func (s *HttpService) Start() error {

	// 调用执行前的逻辑
	if err := s.before.Before(s); err != nil {
		s.L.Errorf("Start前的逻辑执行失败")
		return err
	}

	s.L.Infof("[HTTP] 服务监听地址：%s", s.C.App.HTTP.Addr())
	s.Server.Spin()
	return nil
}

// Stop 停止服务
func (s *HttpService) Stop() error {
	s.L.Infof("服务开始 stop")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		s.L.Warnf("关闭服务异常：%s", err)
		return err
	}
	return nil
}
