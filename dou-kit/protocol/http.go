// @Author: Ciusyan 2023/1/25
package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou-kit/conf"
)

// =====
// 使用HTTP对外暴露服务
// =====

type HttpService struct {
	server *http.Server
	L      logger.Logger // 用于打印日志
	R      gin.IRouter
	C      *conf.Config // 用于获取项目名称

	before StartFuncAop // 执行start前的逻辑
}

func NewHttpService(f StartFuncAop) *HttpService {

	service := &HttpService{
		L:      zap.L().Named("Http Service"),
		C:      conf.C(),
		before: f,
	}

	// New了一个Gin 的Router 实例， 并未加载路由
	r := gin.Default()
	service.R = r

	// HTTP服务对象
	service.server = &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,                   // 1M
		Addr:              service.C.App.HTTP.Addr(), // 获取IP和端口
		Handler:           r,
	}

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
	if err := s.server.ListenAndServe(); err != nil {
		// 如果错误是正常关闭，则不报错
		if err == http.ErrServerClosed {
			s.L.Infof("服务 stop 成功")
			return nil
		}
		return fmt.Errorf("开启 [HTTP] 服务异常：%s", err.Error())
	}

	return nil
}

// Stop 停止服务
func (s *HttpService) Stop() error {
	s.L.Infof("服务开始 stop")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.L.Warnf("关闭服务异常：%s", err)
		return err
	}
	return nil
}
