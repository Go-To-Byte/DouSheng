// @Author: Hexiaoming 2023/2/18
package impl

import (
	"context"
	"errors"
	"gorm.io/gorm"

	// "reflect"

	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/relation-service/apps/relation"
)

// 根据userid获取关注列表(DAO层) TODO：这个可以和 getFollowerListByUserId 的合并成一个方法
func (s *relationServiceImpl) getFollowListByUserId(ctx context.Context, userId int64) (
	[]*relation.UserFollowPo, error) {

	db := s.db.WithContext(ctx)
	set := make([]*relation.UserFollowPo, 50)

	// 构建条件、排序、 查询
	// s.db.Where("id = ?", userId).Find(&set)
	db.Where("user_id = ? AND follow_flag = ?", userId, relation.ActionType_FOLLOW_ACTION).Find(&set)

	if db.Error != nil {
		s.l.Errorf("relation: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}

// 根据userid获取关注列表(DAO层)
func (s *relationServiceImpl) getFollowerListByUserId(ctx context.Context, userId int64) (
	[]*relation.UserFollowerPo, error) {

	db := s.db.WithContext(ctx)
	set := make([]*relation.UserFollowerPo, 50)

	// 构建条件、、 查询
	s.db.Where("user_id = ? AND follower_flag = ?", userId, relation.ActionType_FOLLOW_ACTION).Find(&set)

	if db.Error != nil {
		s.l.Errorf("relation query：查询错误: %s", db.Error.Error())

		return set, db.Error
	}

	return set, nil
}

func (s *relationServiceImpl) getFriendListByUId(ctx context.Context, userId int64) (
	[]*relation.UserFollowerPo, error) {

	db := s.db.WithContext(ctx)
	set := make([]*relation.UserFollowerPo, 0)

	// 构建条件、查询
	db.Raw(friendSql, relation.ActionType_FOLLOW_ACTION, relation.ActionType_FOLLOW_ACTION, userId).Scan(&set)

	if db.Error != nil {
		s.l.Errorf("relation query：查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}

type savePo struct {
	// 关注 po
	follow *relation.UserFollowPo
	// 粉丝 po
	follower *relation.UserFollowerPo
	// 操作类型
	action relation.ActionType
}

func newSavePo(follow *relation.UserFollowPo, follower *relation.UserFollowerPo) *savePo {
	return &savePo{
		follow:   follow,
		follower: follower,
	}
}

func (s *savePo) validate() error {
	if s.follow == nil || s.follower == nil {
		return errors.New("请给属性赋值")
	}
	return nil
}

// 更新、插入操作
func (s *relationServiceImpl) save(ctx context.Context, req *savePo) error {

	err := req.validate()
	if err != nil {
		s.l.Errorf("请给属性赋值")
		return err
	}

	db := s.db.WithContext(ctx)

	switch req.action {
	case relation.ActionType_FOLLOW_ACTION:
		// insert 操作
		return db.Transaction(func(tx *gorm.DB) error {
			// 写入关注数据
			tx.Create(req.follow)
			if tx.Error != nil {
				s.l.Errorf(tx.Error.Error())
				return exception.WithStatusCode(constant.ERROR_SAVE)
			}

			// 写入粉丝数据
			tx.Create(req.follower)
			if tx.Error != nil {
				s.l.Errorf(tx.Error.Error())
				return exception.WithStatusCode(constant.ERROR_SAVE)
			}

			return nil
		})

	case relation.ActionType_UN_FOLLOW_ACTION:
		// unfollow 操作
		return db.Transaction(func(tx *gorm.DB) error {
			// 更新关注数据
			tx.Where("user_id = ? AND follow_id = ?",
				req.follow.UserId, req.follow.FollowId).Save(req.follow)
			if tx.Error != nil {
				s.l.Errorf(tx.Error.Error())
				return exception.WithStatusCode(constant.ERROR_SAVE)
			}

			// 更新粉丝数据
			tx.Where("user_id = ? AND follower_id = ?",
				req.follower.UserId, req.follower.FollowerId).Save(req.follower)

			if tx.Error != nil {
				s.l.Errorf(tx.Error.Error())
				return exception.WithStatusCode(constant.ERROR_SAVE)
			}
			return nil
		})

	case relation.ActionType_AGAIN_FOLLOW:
		// 变成去更新操作
		req.follower.FollowerFlag = relation.ActionType_FOLLOW_ACTION
		req.follow.FollowFlag = relation.ActionType_FOLLOW_ACTION
		req.action = relation.ActionType_UN_FOLLOW_ACTION // 去更新

		// 递归去调用，但是操作已经变成更新了
		return s.save(ctx, req)

	default:
		s.l.Errorf("relation save：只支持更新和插入操作")

		return errors.New("不支持此类型")
	}
}

func (s *relationServiceImpl) getFavoriteCount(ctx context.Context, req *relation.ListCountRequest) (
	*relation.ListCountResponse, error) {

	// 拼接SQL
	db := s.db.WithContext(ctx)
	db1 := db.Model(&relation.UserFollowPo{}).
		Where("user_id = ? AND follow_flag = ?", req.UserId, relation.ActionType_FOLLOW_ACTION)
	db2 := db.Model(&relation.UserFollowerPo{}).
		Where("user_id = ? AND follower_flag = ?", req.UserId, relation.ActionType_FOLLOW_ACTION)
	resp := relation.NewListCountResponse()

	errs := make([]error, 0)

	switch req.Type {
	case relation.CountType_FOLLOW:
		// 查询关注数的条件
		db1 = db1.Count(&resp.FollowCount)
		errs = append(errs, db1.Error)
	case relation.CountType_FOLLOWER:
		// 查询粉丝数的条件
		db2 = db2.Count(&resp.FollowerCount)
		errs = append(errs, db2.Error)
	case relation.CountType_ALL:
		// 两个表都要查
		db1 = db1.Count(&resp.FollowCount)
		errs = append(errs, db1.Error)

		db2 = db2.Count(&resp.FollowerCount)
		errs = append(errs, db2.Error)
	default:
		s.l.Error("relation getFavoriteCount：暂不支持此类型")
	}

	for _, err := range errs {
		if err != nil {
			s.l.Errorf("relation getFavoriteCount：获取总数出现错误，%s", err.Error())
			return resp, err
		}
	}

	return resp, nil
}

// 查询是否存在此记录了
func (s *relationServiceImpl) exist(ctx context.Context, req *relation.UserFollowPo) (bool, error) {

	// 只是查询，看看是否有条记录
	db := s.db.WithContext(ctx).
		Where("user_id = ? AND follow_id = ?", req.UserId, req.FollowId).
		Find(relation.NewUserFollowPo())

	if db.Error != nil {
		s.l.Errorf("relation isFollowerByUId：查询错误，%s", db.Error.Error())
		return false, db.Error
	}

	return db.RowsAffected == 1, nil
}
