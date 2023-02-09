// @Author: Ciusyan 2023/2/7
package rpc

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/common/utils"
)

// GinAuthHandlerFunc HTTP auth中间件
func (a *UserCenterClient) GinAuthHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 从请求中解析出Token
		ak := utils.GetToken(ctx)

		// 验证Token
		req := token.NewValidateTokenRequest(ak)
		tk, err := a.tokenService.ValidateToken(ctx.Request.Context(), req)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, constant.ERROR_TOKEN_VALIDATE)
			// 有错误、直接终止传递
			ctx.Abort()
		} else {
			a.l.Infof("Token认证成功")
		}

		// 把Token传递给下一个链路
		ctx.Set(constant.REQUEST_TOKEN, tk)
		// 把请求传递下去
		ctx.Next()
	}
}
