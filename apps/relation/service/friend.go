// Author: BeYoung
// Date: 2023/1/30 18:20
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
	"go.uber.org/zap"
)

func (r *Relation) FriendList(ctx context.Context, req *proto.FriendListRequest) (*proto.FriendListResponse, error) {
	relation := model.Relation{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: req.UserId,
		Flag:     0,
	}
	result := dao.RelationFindByUserIDAndToUserID(relation)

	list := make([]int64, len(result))
	for i := range result {
		list = append(list, result[i].UserID)
	}
	zap.S().Debugf("id:%v ==> len(friend):%v", req.UserId, len(list))

	return &proto.FriendListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   list,
	}, nil
}
