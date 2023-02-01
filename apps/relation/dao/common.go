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

func Delete(relation model.Relation) (err error) {
	q := query.Use(models.DB)
	r := q.Relation
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.WithContext(context.Background()).Relation.
		Where(r.UserID.Eq(relation.UserID), r.ToUserID.Eq(relation.ToUserID)).
		Update(r.Flag, 0); err != nil {
		zap.S().Panicf("Failed delete relation: %+v", relation)
		return err
	}

	if err = tx.Commit(); err != nil {
		zap.S().Panicf("Failed commit: %v", err)
		return err
	}
	return nil
}

func RelationFindByUserID(relation model.Relation) []*model.Relation {
	q := query.Use(models.DB)
	f := q.Relation

	r, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(relation.UserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find follows: %v", relation.UserID)
	}
	return r
}

func RelationFindByToUserID(relation model.Relation) []*model.Relation {
	q := query.Use(models.DB)
	f := q.Relation

	r, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(relation.ToUserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find followers: %+v", relation)
	}
	return r
}

func RelationFindByUserIDAndToUserID(relation model.Relation) []*model.Relation {
	q := query.Use(models.DB)
	f := q.Relation

	r, err := f.WithContext(context.Background()).
		Where(f.UserID.Eq(relation.UserID), f.ToUserID.Eq(relation.ToUserID), f.Flag.Eq(1)).
		Find()
	if err != nil {
		zap.S().Errorf("Failed find followers: %+v", relation)
	}
	return r
}
