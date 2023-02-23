// Author: BeYoung
// Date: 2023/2/1 0:42
// Software: GoLand

package services

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/network/models"
	proto "github.com/Go-To-Byte/DouSheng/network/protos"
	"go.uber.org/zap"
	"sync"
	"time"
)

// 建立 user grpc 连接, 处理用户信息
func getUserInfo(userID int64, toUserId int64) (response models.User, err error) {
	zap.S().Debugf("get UserInfo: %d", userID)
	relation := proto.NewRelationClient(models.Dials["relation"])
	userRequest := proto.InfoRequest{UserId: toUserId}
	followListRequest := proto.FollowListRequest{UserId: toUserId}
	followerListRequest := proto.FollowerListRequest{UserId: toUserId}
	followJudgeRequest := proto.FollowJudgeRequest{
		UserId:   userID,
		ToUserId: toUserId,
	}

	wait := sync.WaitGroup{}
	wait.Add(4)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// 获取用户信息
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get user info timeout: %v", ctx.Err())
				return

			default:
				user := proto.NewUserClient(models.Dials["user"])
				if r, e := user.Info(context.Background(), &userRequest); err != nil {
					zap.S().Errorf("error getting user info: (%v) ==> %v", userID, e)
				} else {
					response.Name = r.User.Name
					response.ID = r.User.Id
					response.Avatar = r.User.Avatar
					response.BackgroundImage = r.User.BackgroundImage
					response.Signature = r.User.Signature
				}
			}
		}
	}(ctx)

	// 获取关注人数
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get follow timeout: %v", ctx.Err())
				return

			default:
				if r, e := relation.FollowList(context.Background(), &followListRequest); err != nil {
					zap.S().Errorf("error getting relation followList: (%v) ==> %v", userID, e)
				} else {
					response.FollowCount = int64(len(r.UserList))
				}
			}
		}
	}(ctx)

	// 获取粉丝人数
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get follow list timeout: %v", ctx.Err())
				return

			default:
				if r, e := relation.FollowerList(context.Background(), &followerListRequest); err != nil {
					zap.S().Errorf("error getting relation followerList: (%v) ==> %v", userID, e)
				} else {
					response.FollowerCount = int64(len(r.UserList))
				}
			}
		}
	}(ctx)

	// 获取是否关注信息
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("follow judge timeout: %v", ctx.Err())
				return

			default:
				if r, e := relation.FollowJudge(context.Background(), &followJudgeRequest); err != nil {
					zap.S().Errorf("error getting relation followJudge: (%v) ==> %v", userID, e)
				} else {
					response.IsFollow = r.IsFollow == 1
				}
			}
		}
	}(ctx)
	wait.Wait()
	return response, err
}

func getVideoInfo(userID int64, videoID int64) (response models.Video, err error) {
	zap.S().Debugf("get videoInfo: %d", videoID)

	// TODO: goroutine
	var authorID int64 // 作者id
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

	wait := sync.WaitGroup{}
	wait.Add(4)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 获取video info
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get follow list timeout: %v", ctx.Err())
				return

			default:
				if r, e := video.Info(context.Background(), &videoRequest); err != nil {
					zap.S().Errorf("error getting video info: (%v) ==> %v", videoID, e)
				} else {
					response.ID = r.Video.Id
					authorID = r.Video.Author
					response.Title = r.Video.Title
					response.PlayURL = r.Video.PlayUrl
					response.CoverURL = r.Video.CoverUrl
				}
			}
		}
	}(ctx)

	// 获取author info
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get author info timeout: %v", ctx.Err())
				return

			default:
				if r, e := getUserInfo(userID, authorID); err != nil {
					zap.S().Errorf("error getting author info: (%v) ==> %v", userID, e)
				} else {
					response.Author = r
				}
			}
		}
	}(ctx)

	// 获取 CommentCount
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get CommentCount timeout: %v", ctx.Err())
				return

			default:
				if r, e := comment.CommentList(context.Background(), &commentListRequest); err != nil {
					zap.S().Errorf("error getting comment list: (%v) ==> %v", videoID, e)
				} else {
					response.CommentCount = int64(len(r.CommentList))
				}
			}
		}
	}(ctx)

	// 获取 FavoriteCount && IsFavorite
	go func(c context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				zap.S().Errorf("get FavoriteCount && IsFavorite timeout: %v", ctx.Err())
				return

			default:
				if r, e := favorite.FavoredList(context.Background(), &favoredListRequest); err != nil {
					zap.S().Errorf("error getting favored list: (%v) ==> %v", videoID, e)
				} else {
					response.FavoriteCount = int64(len(r.UserList))
				}
				if r, e := favorite.FavoriteJudge(context.Background(), &favoriteJudgeRequest); err != nil {
					zap.S().Errorf("error getting favorite judge: (%v:%v) ==> %v", userID, videoID, e)
				} else {
					response.IsFavorite = r.IsFavorite >= 1
				}
			}
		}
	}(ctx)

	wait.Wait()
	return
}
