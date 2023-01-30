// Author: BeYoung
// Date: 2023/1/30 17:30
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao"
	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
	"go.uber.org/zap"
)

func (r *Relation) Follow(ctx context.Context, req *proto.FollowRequest) (*proto.FollowResponse, error) {
	if err := dao.Add(req.UserId, req.ToUserId); err != nil {
		return &proto.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed to follow",
		}, err
	}
	return &proto.FollowResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
	}, nil
}

func (r *Relation) FollowList(ctx context.Context, req *proto.FollowListRequest) (*proto.FollowListResponse, error) {
	result := dao.FollowsFindByID(req.UserId)
	list := make([]int64, len(result))
	for i := range result {
		list = append(list, result[i].ToUserID)
	}
	zap.S().Debugf("id:%v ==> len(follow):%v", req.UserId, len(list))

	if list == nil || len(list) == 0 {
		return &proto.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "failed to follow list",
			UserList:   nil,
		}, nil
	}
	return &proto.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil
}
