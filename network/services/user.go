// Author: BeYoung
// Date: 2023/2/1 0:42
// Software: GoLand

package services

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/network/milddlewares"
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
	c := proto.NewUserClient(models.Dials["user"])

	// TODO: md5.Sum(password)
	request := proto.RegisterRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}

	var err error
	var response *proto.RegisterResponse
	if response, err = c.Register(ctx, &request); err != nil {
		zap.S().Errorf("Failed to register: %v", &request)
		ctx.JSON(http.StatusBadRequest, models.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "failed to register",
			Token:      "",
			UserID:     0,
		})
		ctx.Abort()
		return
	}

	zap.S().Debugf("Registered: %+v", response)
	jwt := milddlewares.NewJWT()
	token, err := jwt.CreateToken(response.UserId)
	if err != nil {
		zap.S().Panicf("Failed to register: %v", &request)
		ctx.JSON(http.StatusBadRequest, models.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "failed to register",
			Token:      "",
			UserID:     0,
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, models.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      token,
		UserID:     response.UserId,
	})
}

func Login(ctx *gin.Context) {
	zap.S().Debugf("Register")
	c := proto.NewUserClient(models.Dials["user"])

	// TODO: md5.Sum(password)
	request := proto.LoginRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}

	var err error
	var response *proto.LoginResponse
	if response, err = c.Login(ctx, &request); err != nil {
		zap.S().Panicf("Failed to login: %v", &request)
		ctx.JSON(http.StatusBadRequest, models.LoginResponse{
			StatusCode: 1,
			StatusMsg:  "failed to login",
		})
		ctx.Abort()
		return
	}

	zap.S().Debugf("login: %+v", response)
	jwt := milddlewares.NewJWT()
	token, err := jwt.CreateToken(response.UserId)
	if err != nil {
		zap.S().Panicf("Failed to login: %v", &request)
		ctx.JSON(http.StatusBadRequest, models.LoginResponse{
			StatusCode: 1,
			StatusMsg:  "failed to login",
		})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, models.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      token,
		UserID:     response.UserId,
	})
}

func Info(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	zap.S().Debugf("Info: %v", userID)
	c := proto.NewUserClient(models.Dials["user"])

	// 解析查询用户的id
	var err error
	var toUserID int64
	if toUserID, err = strconv.ParseInt(ctx.Query("user_id"), 10, 64); err != nil {
		zap.S().Panicf("Failed to parse user_id(%v): %v", ctx.Query("user_id"), err)
		ctx.JSON(http.StatusBadRequest, models.InfoResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Failed to parse user_id: %v", ctx.Query("user_id")),
		})
		ctx.Abort()
		return
	}

	var response *proto.InfoResponse
	request := proto.InfoRequest{UserId: toUserID}
	if response, err = c.Info(ctx, &request); err != nil {
		zap.S().Errorf("Failed to get user info(%v): %v", toUserID, err)
		ctx.JSON(http.StatusBadRequest, models.InfoResponse{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Failed to get user info: %v", toUserID),
		})
		ctx.Abort()
		return
	}

	zap.S().Debugf("Get user info(%v): %v", toUserID, response)

	// 调用 relation 模块填充数据
	user, err := getUserInfo(userID.(int64), toUserID)
	if err != nil {
		zap.S().Errorf("Failed to get user info")
		ctx.JSON(http.StatusBadRequest, models.InfoResponse{
			StatusCode: 1,
			StatusMsg:  "failed to get user info",
		})
	}
	ctx.JSON(http.StatusOK, models.InfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User:       user,
	})
}
