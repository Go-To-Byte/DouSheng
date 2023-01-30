// Author: BeYoung
// Date: 2023/1/30 18:17
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao"
	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
	"go.uber.org/zap"
)

func (r *Relation) FollowerList(ctx context.Context, req *proto.FollowerListRequest) (*proto.FollowerListResponse, error) {
	result := dao.FollowersFindByID(req.UserId)
	list := make([]int64, len(result))
	for i := range result {
		list = append(list, result[i].ToUserID)
	}
	zap.S().Debugf("id:%v ==> len(follower):%v", req.UserId, len(list))

	if len(list) == 0 {
		return &proto.FollowerListResponse{
			StatusCode: 1,
			StatusMsg:  "failed to follow list",
			UserList:   nil,
		}, nil
	}

	return &proto.FollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil
}
