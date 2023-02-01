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
func getUserInfo(userID int64, to_user_id int64) (response *proto.InfoResponse, err error) {
	zap.S().Debugf("get UserInfo: %d", userID)
	user := proto.NewUserClient(models.GrpcConn)
	relation := proto.NewRelationClient(models.GrpcConn)
	userRequest := proto.InfoRequest{UserId: userID}
	followListRequest := proto.FollowListRequest{UserId: userID}
	followerListRequest := proto.FollowerListRequest{UserId: userID}
	friendRequest := proto.FriendRequest{
		UserId:   userID,
		ToUserId: to_user_id,
	}
	response := models.User{
		FollowCount:   0,
		FollowerCount: 0,
		ID:            0,
		IsFollow:      false,
		Name:          "",
	}

	// 获取用户信息
	if response, err = user.Info(context.Background(), &request); err != nil {
		zap.S().Errorf("error getting user info: (%v) ==> %v", userID, err)
		return response, err
	}

	// 获取关注人数

	// 获取粉丝人数

	// 获取是否关注信息

	return response, err
}
