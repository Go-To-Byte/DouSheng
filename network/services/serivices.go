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
func getUserInfo(userID int64, toUserId int64) (response models.User, err error) {
	zap.S().Debugf("get UserInfo: %d", userID)
	user := proto.NewUserClient(models.Dials["user"])
	relation := proto.NewRelationClient(models.Dials["relation"])
	userRequest := proto.InfoRequest{UserId: toUserId}
	followListRequest := proto.FollowListRequest{UserId: toUserId}
	followerListRequest := proto.FollowerListRequest{UserId: toUserId}
	followJudgeRequest := proto.FollowJudgeRequest{
		UserId:   userID,
		ToUserId: toUserId,
	}

	// 获取用户信息
	if r, e := user.Info(context.Background(), &userRequest); err != nil {
		zap.S().Errorf("error getting user info: (%v) ==> %v", userID, e)
	} else {
		response.Name = r.User.Name
		response.ID = r.User.Id
	}

	// 获取关注人数
	if r, e := relation.FollowList(context.Background(), &followListRequest); err != nil {
		zap.S().Errorf("error getting relation followList: (%v) ==> %v", userID, e)
	} else {
		response.FollowCount = int64(len(r.UserList))
	}

	// 获取粉丝人数
	if r, e := relation.FollowerList(context.Background(), &followerListRequest); err != nil {
		zap.S().Errorf("error getting relation followerList: (%v) ==> %v", userID, e)
	} else {
		response.FollowerCount = int64(len(r.UserList))
	}

	// 获取是否关注信息
	if r, e := relation.FollowJudge(context.Background(), &followJudgeRequest); err != nil {
		zap.S().Errorf("error getting relation followJudge: (%v) ==> %v", userID, e)
	} else {
		response.IsFollow = r.IsFollow == 1
	}

	return response, err
}

func getVideoInfo(userID int64, videoID int64) (response models.Video, err error) {
	zap.S().Debugf("get videoInfo: %d", videoID)

	// TODO: goroutine
	video := proto.NewVideoClient(models.Dials["video"])
	comment := proto.NewCommentClient(models.Dials["comment"])
	favorite := proto.NewFavoriteClient(models.Dials["favorite"])
	videoRequest := proto.VideoInfoRequest{VideoId: videoID}
	commentListRequest := proto.CommentListRequest{VideoId: videoID}
	favoredListRequest := proto.FavoredListRequest{VideoId: videoID}
	favoriteJudgeRequest := proto.FavoriteJudgeRequest{
		UserId:  userID,
		VideoId: videoID,
	}

	// 获取video info
	var authorID int64 // 作者id
	if r, e := video.Info(context.Background(), &videoRequest); err != nil {
		zap.S().Errorf("error getting video info: (%v) ==> %v", videoID, e)
	} else {
		response.ID = r.Video.Id
		authorID = r.Video.Author
		response.Title = r.Video.Title
		response.PlayURL = r.Video.PlayUrl
		response.CoverURL = r.Video.CoverUrl
	}

	// 获取author info
	if r, e := getUserInfo(userID, authorID); err != nil {
		zap.S().Errorf("error getting author info: (%v) ==> %v", userID, e)
	} else {
		response.Author = r
	}

	// 获取 CommentCount
	if r, e := comment.CommentList(context.Background(), &commentListRequest); err != nil {
		zap.S().Errorf("error getting comment list: (%v) ==> %v", videoID, e)
	} else {
		response.CommentCount = int64(len(r.CommentList))
	}

	// 获取 FavoriteCount && IsFavorite
	if r, e := favorite.FavoredList(context.Background(), &favoredListRequest); err != nil {
		zap.S().Errorf("error getting favored list: (%v) ==> %v", videoID, e)
	} else {
		response.FavoriteCount = int64(len(r.UserList))
	}
	if r, e := favorite.FavoriteJudge(context.Background(), &favoriteJudgeRequest); err != nil {
		zap.S().Errorf("error getting favorite judge: (%v:%v) ==> %v", userID, videoID, e)
	} else {
		response.IsFavorite = r.IsFavorite == 1
	}

	return
}
