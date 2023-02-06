// @Author: Ciusyan 2023/2/7
package middlerware

import (
	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	constant2 "github.com/Go-To-Byte/DouSheng/user_center/common/constant"
	"github.com/Go-To-Byte/DouSheng/user_center/utils"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
)

func NewHttpAuther(auther token.ServiceClient) *httpAuther {
	return &httpAuther{
		authenticator: auther,
		l:             zap.L().Named("AUTH"),
	}
}

// HTTP认证器
type httpAuther struct {
	// token 认证器
	authenticator token.ServiceClient

	l logger.Logger
}

// GinAuthHandlerFunc HTTP auth中间件
func (a *httpAuther) GinAuthHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 从请求中解析出Token
		ak := utils.GetToken(ctx)

		// 验证Token
		req := token.NewValidateTokenRequest(ak)
		tk, err := a.authenticator.ValidateToken(ctx.Request.Context(), req)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, constant2.BAD_TOKEN_ERROR)
			return
		}

		a.l.Infof("Token认证成功")
		// 把Token传递给下一个链路
		ctx.Set(constant2.REQUEST_TOKEN, tk)
		// 把请求传递下去
		ctx.Next()
	}
}
