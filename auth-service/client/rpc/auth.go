// @Author: Ciusyan 2023/2/7
package rpc

import (
	"github.com/Go-To-Byte/DouSheng/auth-service/apps/token"
	"github.com/Go-To-Byte/DouSheng/auth-service/common/utils"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"
	"github.com/gin-gonic/gin"
)

// GinAuthHandlerFunc HTTP auth中间件
func (a *AuthServiceClient) GinAuthHandlerFunc() exception.AppHandler {
	return func(ctx *gin.Context) error {

		// 从请求中解析出Token
		ak := utils.GetToken(ctx)

		// 验证Token
		req := token.NewValidateTokenRequest(ak)
		tk, err := a.TokenService().ValidateToken(ctx.Request.Context(), req)

		if err != nil {
			a.l.Errorf("Token认证失败：%s", err.Error())
			// 有错误、直接终止传递
			ctx.Abort()
			return exception.WithStatusCode(constant.ERROR_TOKEN_VALIDATE)
		} else {
			a.l.Infof("Token认证成功")
		}

		// 把Token传递给下一个链路
		ctx.Set(constant.REQUEST_TOKEN, tk)
		// 把请求传递下去
		ctx.Next()

		return nil
	}
}
