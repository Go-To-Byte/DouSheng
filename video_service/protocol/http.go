// @Author: Ciusyan 2023/2/7
package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc/middlerware"

	"github.com/Go-To-Byte/DouSheng/video_service/apps/rpcservice"
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

	// 1、获取中间件
	mids, err := s.getMiddle()
	if err != nil {
		return err
	}

	// 2、将所有的Gin服务对象注册到IOC中
	option := ioc.NewGinOption(s.r, "/"+s.c.App.Name, mids...)
	option.NotVersion = true
	option.NotName = true
	ioc.RegistryGin(option)

	// 3、监听 TCP（HTTP）地址
	if err := s.server.ListenAndServe(); err != nil {
		// 如果错误是正常关闭，则不报错
		if err == http.ErrServerClosed {
			s.l.Infof("服务 stop 成功")
			return nil
		}
		return fmt.Errorf("开启 [HTTP] 服务异常：%s", err.Error())
	}
	s.l.Infof("[HTTP] 服务监听地址：%s", s.c.App.HTTP.Addr())
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

// 获取路由中间件
func (s *HttpService) getMiddle() ([]gin.HandlerFunc, error) {
	// 从配置文件中获取 user_center 的Client
	center, err := rpcservice.NewUserCenterFromCfg()
	if err != nil {
		return nil, err
	}

	// Token认证中间件
	auther := middlerware.NewHttpAuther(center.TokenService())
	middles := []gin.HandlerFunc{auther.GinAuthHandlerFunc()}

	return middles, nil
}
