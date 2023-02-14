// @Author: Ciusyan 2023/2/13
package exception

import (
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
)

// 编写 gin 的中间件

type AppHandler func(c *gin.Context) error

// GinErrWrapper 用于统一处理控制层 error
func GinErrWrapper(handler AppHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		log := zap.L().Named("GinErrWrapper")
		err := handler(c)

		defer func() {
			if r := recover(); r != nil {
				log.Infof("Panic: %v", r)
				c.JSON(http.StatusInternalServerError, custom.Bad(constant.INTERNAL))
			}
		}()

		if err == nil {
			return
		}

		log.Errorf("拦截到异常：%s", err.Error())

		switch e := err.(type) {
		case *custom.Exception:
			c.JSON(http.StatusBadRequest, e.CodeMsg())

			// 还可进行其他 Case 因为我给grpc调用的err用方法GrpcErrWrapper包装了
			// 所以从 控制层的error，基本上都是 custom.Exception
		default:
			c.JSON(http.StatusInternalServerError, custom.Bad(constant.INTERNAL))
		}
	}
}

// GrpcErrWrapper 用于包装 GPRC 调用产生的 err
func GrpcErrWrapper(err error) *custom.Exception {
	if err == nil {
		return nil
	}
	// 通过此方法，一定能够将err转换为 grpc 调用产生的异常
	s := status.Convert(err)

	log := zap.L().Named("GrpcErrWrapper")
	log.Errorf("拦截到异常：%s", s.String())

	// 在通过包装成 自定义异常，统一在 gin 中进行拦截处理
	return WithStatusMsg(s.Message())
}
