// Author: BeYoung
// Date: 2023/2/1 0:42
// Software: GoLand

package services

import (
	"fmt"
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

func Login(ctx *gin.Context) {
	zap.S().Debugf("Register")
	c := proto.NewUserClient(models.GrpcConn)

	// TODO: md5.Sum(password)
	request := proto.LoginRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}

	if response, err := c.Login(ctx, &request); err != nil {
		zap.S().Panicf("Failed to login: %v", &request)
		ctx.JSON(http.StatusBadRequest, models.LoginResponse{
			StatusCode: 1,
			StatusMsg:  "failed to login",
			Token:      "",
			UserID:     0,
		})
		ctx.Abort()
	} else {
		zap.S().Debugf("login: %+v", response)
		ctx.JSON(http.StatusOK, models.LoginResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			Token:      strconv.FormatInt(response.UserId, 10),
			UserID:     response.UserId,
		})
	}
}

func Info(ctx *gin.Context) {
	zap.S().Debugf("Register")
	c := proto.NewUserClient(models.GrpcConn)

	if userID, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64); err != nil {
		zap.S().Panicf("Failed to parse user_id(%v): %v", ctx.Query("user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.InfoResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Failed to parse user_id: %v", ctx.Query("user_id")),
			User:       models.User{},
		})
		ctx.Abort()
		return
	} else {
		request := proto.InfoRequest{UserId: userID}
		if response, err := c.Info(ctx, &request); err != nil {
			zap.S().Panicf("Failed to get user info(%v): %v", userID, err)
			ctx.JSON(http.StatusBadRequest, models.InfoResponse{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("Failed to get user info: %v", userID),
				User:       models.User{},
			})
			ctx.Abort()
			return
		} else {
			zap.S().Debugf("Get user info(%v): %v", userID, response)

			// TODO: 调用 relation模块填充数据
			user := getUserInfo()
			ctx.JSON(http.StatusOK, models.InfoResponse{
				StatusCode: 0,
				StatusMsg:  "success",
				User: models.User{
					FollowCount:   0,
					FollowerCount: 0,
					ID:            response.User.Id,
					IsFollow:      false,
					Name:          response.User.Name,
				},
			})
		}
	}
}