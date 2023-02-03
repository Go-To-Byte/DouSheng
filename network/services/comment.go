// Author: BeYoung
// Date: 2023/2/3 2:02
// Software: GoLand

package services

import (
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Comment(ctx *gin.Context) {
	zap.S().Debugf("Comment")
	c := proto.NewRelationClient(models.Dials["comment"])

	// JWT Authorization
	var err error
	jwt := milddlewares.NewJWT()
	token := &models.TokenClaims{}
	if token, err = jwt.ParseToken(ctx.Query("token")); err != nil {
		zap.S().Panicf("Invalid token value (token: %v): %v", token, err)
		ctx.JSON(http.StatusForbidden, models.CommentResponse{
			Comment:    models.Comment{},
			StatusCode: 1,
			StatusMsg:  "",
		})
		ctx.Abort()
		return
	}
}

func CommentList(ctx *gin.Context) {

}
