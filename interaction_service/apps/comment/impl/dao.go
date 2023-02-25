// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou_kit/exception"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"time"
)

func (c *commentServiceImpl) NewCommentPo(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentPo, error) {
	//根据Token获取User
	userReq := token.NewValidateTokenRequest(req.Token)
	tkRsp, err := c.tokenService.ValidateToken(ctx, userReq)
	if err != nil {
		c.l.Errorf(err.Error())
		return nil, err
	}
	userId := tkRsp.GetUserId()
	po := comment.NewDefaultCommentPo()
	//  TODO ID算法
	po.Id = time.Now().UnixNano()
	po.Content = req.CommentText
	po.CreateDate = time.Now().Format("01-02")
	po.UserId = userId
	po.VideoId = req.VideoId
	return po, nil
}

// 评论功能实现 成功返回Po 失败返回nil
func (c *commentServiceImpl) InsertCommentRecord(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentPo, error) {
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
	db = c.db.WithContext(ctx)
	db = db.Where(" id = ? ", req.CommentId).Find(&commentPo)
	if db.Error != nil {
		c.l.Errorf("评论失败: %s", err.Error())
		return nil, err
	}
	//评论必然存在，找到之后判断是否是本人的评论
	if commentPo.UserId != validatedToken.GetUserId() {
		//返回错误，没有权限删除
		return nil, exception.WithStatusCode(constant.WRONG_NO_PERMISSION)
	}
	//删除失败
	db.Delete(&commentPo)
	if db.Error != nil || db.RowsAffected == 0 {
		c.l.Errorf("删除数据失败:%s", err.Error())
		return nil, exception.WithStatusCode(constant.ERROR_REMOVE)
	}
	//删除成功
	return commentPo, nil
}

func (c *commentServiceImpl) GetCommentPoList(ctx context.Context, req *comment.GetCommentListRequest) ([]*comment.CommentPo, error) {
	db := c.db.WithContext(ctx)
	var count int64 = 0
	//查找评论数量
	db.Table("comment").Where("video_id = ?", req.VideoId).Count(&count)
	pos := make([]*comment.CommentPo, count)
	if count == 0 {
		return pos, nil
	}
	db.Table("comment").Where("video_id = ?", req.VideoId).Find(&pos)
	return pos, nil
}

func (c *commentServiceImpl) GetCommentCount(ctx context.Context, req *comment.GetCommentCountByIdRequest) (*int64, error) {
	db := c.db.WithContext(ctx)
	var count int64
	db.Table("comment").Where(" video_id = ?", req.VideoId).Count(&count)
	if db.Error != nil {
		c.l.Errorf("喜欢视频总数查询失败:%s", db.Error.Error())
		return nil, db.Error
	}
	return &count, nil
}
