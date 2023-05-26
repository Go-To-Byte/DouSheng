// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"context"
	"time"

	"github.com/Go-To-Byte/DouSheng/api-rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/dou-kit/exception"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
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

func (s *commentServiceImpl) getCommentPoList(ctx context.Context, req *comment.GetCommentListRequest) ([]*comment.CommentPo, error) {
	db := s.db.WithContext(ctx)

	pos := make([]*comment.CommentPo, 10)
	db.Where("video_id = ?", req.VideoId).Find(&pos)
	if db.Error != nil {
		s.l.Errorf("comment getCommentPoList：出现错误，%s", db.Error)
		return pos, db.Error
	}

	return pos, nil
}

func (s *commentServiceImpl) getCommentCount(ctx context.Context, videoId int64) (int64, error) {
	db := s.db.WithContext(ctx)
	var count int64
	db.Model(&comment.CommentPo{}).Where("video_id = ?", videoId).Count(&count)
	if db.Error != nil {
		s.l.Errorf("喜欢视频总数查询失败:%s", db.Error.Error())
		return 0, db.Error
	}

	return count, nil
}
