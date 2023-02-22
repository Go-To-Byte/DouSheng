// Package impl
// Author: BeYoung
// Date: 2023/2/22 21:43
// Software: GoLand
package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/relation/apps/relation/impl/dal/model"
	"github.com/Go-To-Byte/DouSheng/relation/apps/relation/impl/dal/query"
	"go.uber.org/zap"
)

func (r *RelationServiceImpl) Add(relation model.Relation) (err error) {
	q := query.Use(r.db)
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

func (r *RelationServiceImpl) Delete(relation model.Relation) (err error) {
	q := query.Use(r.db)
	re := q.Relation
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.WithContext(context.Background()).Relation.
		Where(re.UserID.Eq(relation.UserID), re.ToUserID.Eq(relation.ToUserID)).
		Update(re.Flag, 0); err != nil {
		zap.S().Panicf("Failed delete relation: %+v", relation)
		return err
	}

	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func (r *RelationServiceImpl) FindByUserID(relation model.Relation) []*model.Relation {
	q := query.Use(r.db)
	f := q.Relation

	re, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(relation.UserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find follows: %v", relation.UserID)
	}
	return re
}

func (r *RelationServiceImpl) FindByToUserID(relation model.Relation) []*model.Relation {
	q := query.Use(r.db)
	f := q.Relation

	re, err := f.WithContext(context.Background()).
		Where(f.ToUserID.Eq(relation.ToUserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find followers: %+v", relation)
	}
	return re
}

func (r *RelationServiceImpl) FindByUserIDWithToUserID(relation model.Relation) []*model.Relation {
	q := query.Use(r.db)
	f := q.Relation

	re, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(relation.UserID), f.ToUserID.Eq(relation.ToUserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find followers: %+v", relation)
	}
	return re
}

func (r *RelationServiceImpl) RelationJudge(relation model.Relation) bool {
	// 如果 id 一致，直接返回 true
	if relation.UserID == relation.ToUserID {
		return true
	}

	q := query.Use(r.db)
	f := q.Relation

	re, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(relation.UserID), f.ToUserID.Eq(relation.ToUserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find followers: %+v", relation)
	}

	return len(re) >= 1
}
