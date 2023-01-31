// Author: BeYoung
// Date: 2023/2/1 0:42
// Software: GoLand

package services

import (
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Register Http API
func Register(ctx *gin.Context) {
	zap.S().Debugf("Register")
	c := proto.NewUserClient(models.GrpcConn)

	// TODO: md5.Sum(password)
	request := proto.RegisterRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}

	if response, err := c.Register(ctx, &request); err != nil {
		zap.S().Panicf("Failed to register: %v", &request)
		ctx.JSON(http.StatusBadRequest, models.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "failed to register",
			Token:      "",
			UserID:     0,
		})
		ctx.Abort()
	} else {
		zap.S().Debugf("Registered: %+v", response)
		ctx.JSON(http.StatusOK, models.RegisterResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			Token:      strconv.FormatInt(response.UserId, 10),
			UserID:     response.UserId,
		})
	}
}
