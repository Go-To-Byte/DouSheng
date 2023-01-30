// Author: BeYoung
// Date: 2023/1/30 15:50
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/proto"
)

func (u *User) Info(ctx context.Context, req *proto.InfoRequest) (*proto.InfoResponse, error) {
	info := model.Info{
		ID:            req.GetUserId(),
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
	}

	// 查询用户
	results := dao.InfoFindByID(info)
	userInfo := proto.Info{
		Id:            results[0].ID,
		Name:          results[0].Name,
		FollowCount:   results[0].FollowCount,
		FollowerCount: results[0].FollowerCount,
		IsFollow:      false,
	}

	return &proto.InfoResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		Info:       &userInfo,
	}, nil
}
