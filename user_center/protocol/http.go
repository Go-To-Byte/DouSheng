// @Author: Ciusyan 2023/1/25
package protocol

import (
	"Go-To-Byte/DouSheng/user_center/conf"
	"Go-To-Byte/DouSheng/user_center/ioc"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"
)

// =====
// 使用HTTP对外暴露服务
// =====

func NewHttpService() *HttpService {

	service := &HttpService{
		l: zap.L().Named("Http Service"),
		c: conf.C(),
	}

	// New了一个Gin 的Router 实例， 并未加载路由
	r := gin.Default()
	service.r = r

	// HTTP服务对象
	service.server = &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,                   // 1M
		Addr:              service.c.App.HTTP.Addr(), // 获取IP和端口
		Handler:           r,
	}

	return service
}

type HttpService struct {
	server *http.Server
	l      logger.Logger // 用于打印日志
	r      gin.IRouter
	c      *conf.Config // 用于获取项目名称
}

// Start 开启服务
func (s *HttpService) Start() error {
	// 1、将所有的Gin服务对象注册到IOC中

	// 拼接好前缀再注册："/douying"
	ioc.RegistryGin("/"+s.c.App.Name, s.r)

	s.l.Infof("[HTTP] 服务监听地址：%s", s.c.App.HTTP.Addr())
	if err := s.server.ListenAndServe(); err != nil {
		// 如果错误是正常关闭，则不报错
		if err == http.ErrServerClosed {
			s.l.Infof("服务 stop 成功")
			return nil
		}
		return fmt.Errorf("开启 [HTTP] 服务异常：%s", err.Error())
	}

	return nil
}

// Stop 停止服务
func (s *HttpService) Stop() error {
	s.l.Infof("服务开始 stop")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("关闭服务异常：%s", err)
		return err
	}
	return nil
}
