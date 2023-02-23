// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"time"
)

func (c *commentServiceImpl) NewCommentPo(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentPo, error) {
	//根据Token获取User
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := c.tokenService.ValidateToken(ctx, tokenReq)
	if err != nil {
		c.l.Errorf(err.Error())
		return nil, err
	}
	po := comment.NewDefaultCommentPo()
	//  TODO ID算法
	po.Id = time.Now().UnixNano()
	po.Content = req.CommentText
	po.CreateDate = time.Now().Format("01-02")
	po.UserId = validatedToken.GetUserId()
	po.VideoId = req.VideoId
	return po, nil
}

// 评论功能实现 成功返回Po 失败返回nil
func (c *commentServiceImpl) InsertCommentRecord(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentPo, error) {
	db := c.db.WithContext(ctx)
	//根据Token获取User
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := c.tokenService.ValidateToken(ctx, tokenReq)
	if err != nil {
		c.l.Errorf(err.Error())
		return nil, err
	}
	//检查是否已经存在此纪录
	db = db.Where("user_id = ? AND video_id = ?", validatedToken.GetUserId(), req.VideoId).Find(&comment.CommentPo{})
	//记录已存在
	if db.RowsAffected != 0 {
		return nil, exception.WithStatusCode(constant.WRONG_EXIST_USERS)
	}
	//	保存新记录
	po, err := c.NewCommentPo(ctx, req)
	if err != nil {
		return nil, err
	}
	tx := c.db.WithContext(ctx).Create(po)
	if tx.Error != nil {
		return nil, exception.WithStatusCode(constant.ERROR_SAVE)
	}
	return po, err
}

// 删除评论，成功返回po，失败返回nil
func (c *commentServiceImpl) DeleteCommentById(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentPo, error) {
	//根据Token获取User
	tokenReq := token.NewValidateTokenRequest(req.Token)
	validatedToken, err := c.tokenService.ValidateToken(ctx, tokenReq)
	if err != nil {
		c.l.Errorf(err.Error())
		return nil, err
	}
	//检查是否已经存在此纪录
	db := c.db.WithContext(ctx)
	commentPo := comment.NewDefaultCommentPo()
	//通过comment_id获取评论
	db = c.db.WithContext(ctx).Where("id = ?", req.CommentId).First(&commentPo)
	if db.Error != nil {
		return nil, err
	}
	fmt.Println(db.RowsAffected)
	//评论必然存在，找到之后判断是否是本人的评论
	if commentPo.UserId != validatedToken.GetUserId() {
		//返回错误，没有权限删除
		return nil, exception.WithStatusCode(constant.WRONG_NO_PERMISSION)
	}
	//po := comment.NewDefaultCommentPo()
	db.Delete(&commentPo)
	fmt.Println(db.RowsAffected)
	//删除失败
	if db.RowsAffected == 0 {
		return nil, exception.WithStatusCode(constant.WRONG_USER_NOT_EXIST)
	}
	//删除成功
	return commentPo, nil
}

func (c *commentServiceImpl) GetCommentPoList(ctx context.Context, req *comment.GetCommentListRequest) ([]*comment.CommentPo, error) {
	db := c.db.Where(ctx)
	var count int64 = 0
	//查找评论数量
	db.Where("video_id = ?", req.VideoId).Count(&count)
	list := make([]*comment.CommentPo, count)
	if count == 0 {
		return list, nil
	}
	db.Where("video_id = ?", req.VideoId).Find(&list)
	return list, nil
}
