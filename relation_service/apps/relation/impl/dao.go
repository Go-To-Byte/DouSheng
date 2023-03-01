// @Author: Hexiaoming 2023/2/18
package impl

import (
	"context"
	// "reflect"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"

	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
)

// 根据userid获取关注列表(DAO层)
func (s *relationServiceImpl) getFollowListByUserId(ctx context.Context, userId int64) (
	[]*relation.UserFollowPo, error) {

	db := s.db.WithContext(ctx)
	set := make([]*relation.UserFollowPo, 50)

	// 构建条件、排序、 查询
	// s.db.Where("id = ?", userId).Order("created_at desc").Find(&set)
	s.db.Where("user_id = ? AND follow_flag = ?", userId, constant.FOLLOW_ACTION).Find(&set)
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

	// 构建条件、排序、 查询
	// s.db.Where("id = ?", userId).Order("created_at desc").Find(&set)
	s.db.Where("user_id = ? AND follower_flag = ?", userId, constant.FOLLOW_ACTION).Find(&set)
	if db.Error != nil {
		s.l.Errorf("relation: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
}

func (s *relationServiceImpl) update(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.UserFollowPo, error) {

	// 获取 UserFollow po 对象
	userFollowPo, err := s.getUserFollowPo(ctx, req)
	if err != nil {
		return nil, err
	}

	// 更新关注数据
	tx := s.db.WithContext(ctx).Where("user_id=?", userFollowPo.UserId).Where("follow_id=?", userFollowPo.FollowId).Save(userFollowPo)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}

	// 获取 UserFollower po 对象
	userFollowerPo, err := s.getUserFollowerPo(ctx, req)
	if err != nil {
		return nil, err
	}

	// 更新粉丝数据
	tx = s.db.WithContext(ctx).Where("user_id=?", userFollowerPo.UserId).Where("follower_id=?", userFollowerPo.FollowerId).Save(userFollowerPo)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}

	return nil, err
}

func (s *relationServiceImpl) insert(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.UserFollowPo, error) {

	// 虽然这里存在数据冗余, 但是在做检索关注/粉丝数时更快
	// TODO 这里可能存在数据不一致性, 需要保证两次事务要么都执行要么都回退, 待处理

	// 写入关注数据
	// 获取 UserFollow po 对象
	userFollowPo, err := s.getUserFollowPo(ctx, req)
	if err != nil {
		return nil, err
	}

	tx := s.db.WithContext(ctx).Create(userFollowPo)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}

	// 写入粉丝数据
	userFollowerPo, err := s.getUserFollowerPo(ctx, req)
	if err != nil {
		return nil, err
	}

	// 关注写入数据库
	tx = s.db.WithContext(ctx).Create(userFollowerPo)
	if tx.Error != nil {
		s.l.Errorf(tx.Error.Error())
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}

	return nil, err
}

func (s *relationServiceImpl) getUserFollowPo(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.UserFollowPo, error) {

	tokenReq := token.NewValidateTokenRequest(req.Token)

	// 获取用户ID
	validatedToken, err := s.tokenService.ValidateToken(ctx, tokenReq)

	if err != nil {
		s.l.Errorf(err.Error())
		return nil, err
	}

	userFollowPo := relation.NewUserFollowPo(req)
	userFollowPo.UserId = validatedToken.GetUserId()
	userFollowPo.FollowFlag = req.ActionType
	return userFollowPo, nil
}

func (s *relationServiceImpl) getUserFollowerPo(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.UserFollowerPo, error) {

	tokenReq := token.NewValidateTokenRequest(req.Token)

	// 获取用户ID
	validatedToken, err := s.tokenService.ValidateToken(ctx, tokenReq)

	if err != nil {
		s.l.Errorf(err.Error())
		return nil, err
	}

	userFollowerPo := relation.NewUserFollowerPo(req)
	userFollowerPo.FollowerId = validatedToken.GetUserId()
	userFollowerPo.FollowerFlag = req.ActionType
	return userFollowerPo, nil
}

func (s *relationServiceImpl) getFavoriteCount(ctx context.Context, req *relation.ListCountRequest) (
	*relation.ListCountResponse, error) {

	// 拼接SQL
	db := s.db.WithContext(ctx)
	db1 := db.Model(&relation.UserFollowPo{}).
		Where("user_id = ? AND follow_flag = ?", req.UserId, constant.FOLLOW_ACTION)
	db2 := db.Model(&relation.UserFollowerPo{}).
		Where("user_id = ? AND follower_flag = ?", req.UserId, constant.FOLLOW_ACTION)

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
