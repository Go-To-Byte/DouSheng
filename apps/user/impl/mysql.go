// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 检查是否满足接口约束
var _ user.Service = (*UserServiceImpl)(nil)

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		// User模块服务的子Logger
		l: zap.L().Named("User"),
	}
}

// UserServiceImpl 基于Mysql实现的Service
type UserServiceImpl struct {
	// 日志实例
	l logger.Logger
}
