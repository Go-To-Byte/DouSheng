// Author: BeYoung
// Date: 2023/1/30 2:46
// Software: GoLand

package dao

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao/dal/query"
	"github.com/Go-To-Byte/DouSheng/apps/relation/models"

	"go.uber.org/zap"
)

func Add(relation model.Relation) (err error) {
	q := query.Use(models.DB)
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = tx.Relation.Create(&relation); err != nil {
		zap.S().Panicf("Failed create relation: %+v", relation)
		return err
	}

	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func FollowsFindByID(userID int64) []*model.Follow {
	q := query.Use(models.DB)
	f := q.Follow

	r, err := f.WithContext(context.Background()).Where(f.UserID.Eq(userID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find follows: %v", userID)
	}
	return r
}

func FollowersFindByID(userID int64) []*model.Follower {
	q := query.Use(models.DB)
	f := q.Follower

	r, err := f.WithContext(context.Background()).Where(f.UserID.Eq(userID)).Find()
	if err != nil {
		zap.S().Panicf("Failed find follows: %v", userID)
	}
	return r
}

func FriendsFindByID(userID int64) []int64 {
	follows := FollowsFindByID(userID)
	followers := FollowersFindByID(userID)
	var hash = map[int64]bool{}
	list := make([]int64, 5)
	for i := range follows {
		if follows[i].Flag != 0 {
			hash[follows[i].ToUserID] = true
		}
	}
	for i := range followers {
		if hash[followers[i].ToUserID] && followers[i].Flag != 0 {
			list = append(list, followers[i].ToUserID)
		}
	}
	return list
}
