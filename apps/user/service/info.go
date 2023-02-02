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
	user := model.User{
		ID:       req.UserId,
		Username: "",
		Passwd:   "",
	}

	// 查询用户
	zap.S().Debugf("user: %v", user.ID)
	r := dao.UserFindById(user)
	if len(r) > 0 {
		userInfo := proto.User{
			Id:            r[0].ID,
			Name:          r[0].Username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
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
