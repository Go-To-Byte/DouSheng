// @Author: Ciusyan 2023/2/7
package impl

import (
	"context"

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
	s.db.Where("user_id=?", userId).Find(&set)
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
	s.db.Where("user_id=?", userId).Find(&set)
	if db.Error != nil {
		s.l.Errorf("relation: query 查询错误: %s", db.Error.Error())
		return set, db.Error
	}

	return set, nil
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
	return userFollowerPo, nil
}
