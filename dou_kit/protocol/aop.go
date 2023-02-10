// @Author: Ciusyan 2023/2/10
package protocol

import "github.com/Go-To-Byte/DouSheng/dou_kit/ioc"

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
		option := ioc.NewGinOption(s.R, "/"+s.C.App.Name)
		ioc.RegistryGin(option)
		return nil
	}
}
