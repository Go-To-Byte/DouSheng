// @Author: Ciusyan 2023/2/10
package protocol

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

// StartFuncAop 用于执行前的逻辑
type StartFuncAop func(*HttpService) error

// Before 用于执行前，切入代码
func (f StartFuncAop) Before(h *HttpService) error {
	return f(h)
}

// After 用于执行后，切入代码
func (f StartFuncAop) After(h *HttpService) error {
	return f(h)
}

// DefaultHttpStartBefore 返回一个默认的切面：HTTPStartBefore
func DefaultHttpStartBefore() StartFuncAop {
	return func(s *HttpService) error {
		// 1、将所有的Gin服务对象注册到IOC中
		option := ioc.NewHertzOption(s.Server, "/"+s.C.App.Name)
		ioc.RegistryHertz(option)
		return nil
	}
}

// AccessLog 日志中间件
func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)

		end := time.Now()
		latency := end.Sub(start).Microseconds
		hlog.CtxTracef(c, "status=%d cost=%d method=[%s] URI=%s",
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.URI().String())
	}
}
