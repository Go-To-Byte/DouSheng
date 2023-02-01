// Author: BeYoung
// Date: 2023/2/1 0:42
// Software: GoLand

package services

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"go.uber.org/zap"
)

// 建立 user grpc 连接, 处理用户信息
func getUserInfo(userID int64, toUserId int64) (response *proto.InfoResponse, err error) {
	zap.S().Debugf("get UserInfo: %d", userID)
	user := proto.NewUserClient(models.GrpcConn)
	relation := proto.NewRelationClient(models.GrpcConn)
	userRequest := proto.InfoRequest{UserId: userID}
	followListRequest := proto.FollowListRequest{UserId: userID}
	followerListRequest := proto.FollowerListRequest{UserId: userID}
	followJudgeRequest := proto.FollowJudgeRequest{
		UserId:   userID,
		ToUserId: toUserId,
	}

	// 获取用户信息
	if r, e := user.Info(context.Background(), &userRequest); err != nil {
		zap.S().Errorf("error getting user info: (%v) ==> %v", userID, e)
	} else {
		response.User.Name = r.User.Name
		response.User.Id = r.User.Id
	}

	// 获取关注人数
	if r, e := relation.FollowList(context.Background(), &followListRequest); err != nil {
		zap.S().Errorf("error getting relation followList: (%v) ==> %v", userID, e)
	} else {
		response.User.FollowCount = int64(len(r.UserList))
	}

	// 获取粉丝人数
	if r, e := relation.FollowerList(context.Background(), &followerListRequest); err != nil {
		zap.S().Errorf("error getting relation followerList: (%v) ==> %v", userID, e)
	} else {
		response.User.FollowerCount = int64(len(r.UserList))
	}

	// 获取是否关注信息
	if r, e := relation.FollowJudge(context.Background(), &followJudgeRequest); err != nil {
		zap.S().Errorf("error getting relation followJudge: (%v) ==> %v", userID, e)
	} else {
		response.User.IsFollow = r.IsFollow == 1
	}

	return response, err
}