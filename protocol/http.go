// @Author: Ciusyan 2023/1/25
package protocol

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"
)

// 使用HTTP对外暴露服务

func NewHttpService() *HttpService {

	// New了一个Gin 的Router 实例， 并未加载路由
	r := gin.Default()

	// HTTP服务对象
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}

	return &HttpService{
		server: server,
		l:      zap.L().Named("Http Service"),
		r:      r,
	}
}

type HttpService struct {
	server *http.Server
	l      logger.Logger
	r      gin.IRouter
}

// Start 开启服务
func (s *HttpService) Start() error {
	// 1、将所有的Gin服务对象注册到IOC中
	apps.InitGin(s.r)
	// 打印已加载的Gin 服务
	s.l.Infof("已加载的Gin Apps：%v", apps.LoadedGinApps())

	if err := s.server.ListenAndServe(); err != nil {

		// 如果错误是正常关闭，则不报错
		if err == http.ErrServerClosed {
			s.l.Infof("服务 stop 成功")
			return nil
		}
		return fmt.Errorf("开启服务异常：%s", err.Error())
	}

	return nil
}

// Stop 停止服务
func (s *HttpService) Stop() {
	s.l.Infof("服务开始 stop")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("关闭服务异常：%s", err)
	}
}
