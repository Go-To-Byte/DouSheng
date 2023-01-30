// Author: BeYoung
// Date: 2023/1/30 18:20
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao"
	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
)

func (r *Relation) FriendList(ctx context.Context, req *proto.FriendListRequest) (*proto.FriendListResponse, error) {
	result := dao.FriendsFindByID(req.UserId)
	return &proto.FriendListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   result,
	}, nil
}
