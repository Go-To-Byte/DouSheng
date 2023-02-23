// Author: BeYoung
// Date: 2023/1/30 15:50
// Software: GoLand

package service

import (
	"context"
	"errors"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/proto"
	"go.uber.org/zap"
)

func (u *User) Info(ctx context.Context, req *proto.InfoRequest) (*proto.InfoResponse, error) {
	user := model.User{ID: req.UserId}

	// 查询用户
	zap.S().Debugf("user: %v", user.ID)
	r := dao.UserFindById(user)
	if len(r) > 0 {
		userInfo := proto.User{
			Id:              r[0].ID,
			Name:            r[0].Username,
			Avatar:          "https://cyan-1257348513.cos.ap-shanghai.myqcloud.com/avatar/7.jpg",
			BackgroundImage: "https://cyan-1257348513.cos.ap-shanghai.myqcloud.com/background/62.jpg",
			Signature:       "hello world",
		}
		return &proto.InfoResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
			User:       &userInfo,
		}, nil
	}

	return &proto.InfoResponse{
		StatusCode: 1,
		StatusMsg:  "failed get user info",
		User:       nil,
	}, errors.New("failed get user info")
}
