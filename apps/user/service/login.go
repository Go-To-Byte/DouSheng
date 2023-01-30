// Author: BeYoung
// Date: 2023/1/30 2:41
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao"
	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/user/proto"
	"strconv"
)

func (r *User) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	u := model.User{
		ID:       0,
		Username: req.Username,
		Passwd:   req.Password,
	}

	// 查询用户是否存在
	results := dao.FindByName(u)
	if results == nil || len(results) > 1 {
		return &proto.LoginResponse{
			StatusCode: 6,
			StatusMsg:  "user not found",
			UserId:     results[0].ID,
			Token:      "",
		}, nil
	}

	// 密码匹配，TODO: 使用加密算法匹配
	result := results[0]
	if u.Passwd != result.Passwd {
		return &proto.LoginResponse{
			StatusCode: 6,
			StatusMsg:  "passwd mismatch",
			UserId:     results[0].ID,
			Token:      "",
		}, nil
	}

	return &proto.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserId:     result.ID,
		Token:      strconv.FormatInt(result.ID, 10),
	}, nil
}
