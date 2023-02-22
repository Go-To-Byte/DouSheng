// Author: BeYoung
// Date: 2023/2/22 21:43
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/relation/apps/relation"
	"github.com/Go-To-Byte/DouSheng/relation/apps/relation/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/relation/apps/relation/impl/models"
	"go.uber.org/zap"
)

func (r *RelationServiceImpl) Follow(ctx context.Context, req *relation.FollowRequest) (*relation.FollowResponse, error) {
	re := model.Relation{
		ID:       models.Node.Generate().Int64(),
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Flag:     1,
	}
	if req.ActionType == 1 {
		if err := r.Add(re); err != nil {
			return &relation.FollowResponse{
				StatusCode: 1,
				StatusMsg:  "failed add follow",
			}, err
		}
	} else if req.ActionType == 2 {
		if err := r.Delete(re); err != nil {
			return &relation.FollowResponse{
				StatusCode: 1,
				StatusMsg:  "failed delete follow",
			}, err
		}
	}
	return &relation.FollowResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (r *RelationServiceImpl) FollowJudge(ctx context.Context, req *relation.FollowJudgeRequest) (*relation.FollowJudgeResponse, error) {
	re := model.Relation{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Flag:     1,
	}

	if result := r.RelationJudge(re); result == true {
		return &relation.FollowJudgeResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			IsFollow:   1,
		}, nil
	}
	return &relation.FollowJudgeResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		IsFollow:   0,
	}, nil
}

func (r *RelationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListRequest) (*relation.FollowListResponse, error) {
	re := model.Relation{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: 0,
		Flag:     1,
	}
	result := r.FindByUserID(re)
	list := make([]int64, 0)
	for i := range result {
		list = append(list, result[i].ToUserID)
	}
	zap.S().Debugf("id:%v ==> len(follow):%v", req.UserId, len(list))

	return &relation.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil
}

func (r *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (*relation.FollowerListResponse, error) {
	re := model.Relation{
		ID:       0,
		UserID:   0,
		ToUserID: req.UserId,
		Flag:     1,
	}
	result := r.FindByToUserID(re)
	list := make([]int64, 0)
	for i := range result {
		list = append(list, result[i].UserID)
	}
	zap.S().Debugf("id:%v ==> len(follower):%v", req.UserId, len(list))

	return &relation.FollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   list,
	}, nil
}

func (r *RelationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListRequest) (*relation.FriendListResponse, error) {
	re := model.Relation{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: req.UserId,
		Flag:     0,
	}
	result := r.FindByUserIDWithToUserID(re)

	list := make([]int64, 0)
	for i := range result {
		list = append(list, result[i].UserID)
	}
	zap.S().Debugf("id:%v ==> len(friend):%v", req.UserId, len(list))

	return &relation.FriendListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   list,
	}, nil
}
