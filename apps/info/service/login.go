// Author: BeYoung
// Date: 2023/1/30 2:41
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/info/dao"
	"github.com/Go-To-Byte/DouSheng/apps/info/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/info/proto"
)

func (i *Info) Login(ctx context.Context, req *proto.InfoRequest) (*proto.InfoResponse, error) {
	info := model.Info{
		ID:            req.GetUserId(),
		Phone:         "",
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
	}

	// 查询用户
	results := dao.FindById(info)
	user := proto.User{
		Id:            results[0].ID,
		Name:          results[0].Name,
		FollowCount:   results[0].FollowCount,
		FollowerCount: results[0].FollowerCount,
		IsFollow:      false,
	}

	return &proto.InfoResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		User:       &user,
	}, nil
}
