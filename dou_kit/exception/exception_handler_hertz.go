// @Author: Ciusyan 2023/2/13
package exception

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/exception/custom"
)

// 编写 hertz 的中间件

type HertzHandler func(c context.Context, ctx *app.RequestContext) error

// HertzErrWrapper 用于统一处理控制层 error
func HertzErrWrapper(handler HertzHandler) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		log := zap.L().Named("HertzErrWrapper")
		err := handler(c, ctx)

		defer func() {
			if r := recover(); r != nil {
				log.Infof("Panic: %v", r)
				ctx.JSON(http.StatusInternalServerError, custom.Bad(constant.INTERNAL))
			}
		}()

		if err == nil {
			return
		}

		log.Errorf("拦截到异常：%s", err.Error())

		switch e := err.(type) {
		case *custom.Exception:
			ctx.JSON(http.StatusBadRequest, e.CodeMsg())

			// 还可进行其他 Case 因为我给grpc调用的err用方法GrpcErrWrapper包装了
			// 所以从 控制层的error，基本上都是 custom.Exception
		default:
			ctx.JSON(http.StatusInternalServerError, custom.Bad(constant.INTERNAL))
		}
	}
}
