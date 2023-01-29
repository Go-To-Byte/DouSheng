// Author: BeYoung
// Date: 2023/1/26 2:47
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/models"
	"github.com/Go-To-Byte/DouSheng/apps/user/proto"
	"go.uber.org/zap"
)

// // Register Http API
// func Register(context *gin.Context) {
// 	zap.S().Debugf("Register")
// 	user := model.UserInfo{
// 		ID:       models.Node.Generate().Int64(),
// 		Username: context.Query("username"),
// 		Passwd:   context.Query("password"),
// 	}
//
// 	// TODO: md5.Sum(password)
// 	passwd := make([]byte, 128)
// 	copy(passwd, user.Passwd)
//
// 	result := dao.FindByName(user)
// 	if result == nil || len(result) > 0 {
// 		context.JSON(http.StatusForbidden, gin.H{
// 			"user_id":     0,
// 			"status_code": http.StatusForbidden,
// 			"status_msg":  "failed",
// 			"token":       "",
// 		})
// 		return
// 	}
//
// 	zap.S().Debugf("add user: %+v", user)
// 	dao.Add(user)
// 	context.JSON(http.StatusOK, gin.H{
// 		"user_id":     user.ID,
// 		"status_code": http.StatusOK,
// 		"status_msg":  "success",
// 		"token":       "token",
// 	})
// }

type Register struct {
	proto.UnimplementedUserServer
}

func (r *Register) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	u := model.UserInfo{
		ID:       models.Node.Generate().Int64(),
		Username: req.Username,
		Passwd:   req.Password,
	}

	// 查询用户是否存在
	result := dao.FindByName(u)
	if result == nil || len(result) > 0 {
		return &proto.RegisterResponse{
			StatusCode: 6,
			StatusMsg:  "registered",
			UserId:     result[0].ID,
			Token:      "",
		}, nil
	}

	// 添加用户, TODO: 密码加密
	zap.S().Debugf("add user: %+v", u)
	dao.Add(u)

	// TODO: 生成token
	return &proto.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserId:     u.ID,
		Token:      string(u.ID),
	}, nil
}
