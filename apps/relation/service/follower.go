// Author: BeYoung
// Date: 2023/1/30 18:17
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
	"go.uber.org/zap"
)

func (r *Relation) FollowerList(ctx context.Context, req *proto.FollowerListRequest) (*proto.FollowerListResponse, error) {
	relation := model.Relation{
		ID:       0,
		UserID:   0,
		ToUserID: req.UserId,
		Flag:     1,
	}
	result := dao.FindByToUserID(relation)
	list := make([]int64, 0)
	for i := range result {
		list = append(list, result[i].UserID)
	}
	zap.S().Debugf("id:%v ==> len(follower):%v", req.UserId, len(list))

	return &proto.FollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   list,
	}, nil
}
