// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment"
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment/impl/dal/model"
	_ "github.com/Go-To-Byte/DouSheng/comment/apps/comment/impl/init"
	"github.com/Go-To-Byte/DouSheng/comment/apps/comment/impl/models"

	"go.uber.org/zap"
	"time"
)

func (c *CommentServiceImpl) Comment(ctx context.Context, req *comment.CommentRequest) (*comment.CommentResponse, error) {
	// 从 request 获取评论信息
	cmt := model.Comment{
		ID:      models.Node.Generate().Int64(),
		VideoID: req.VideoId,
		UserID:  req.UserId,
		Content: req.Content,
	}

	// 添加评论
	if err := c.Add(cmt); err != nil {
		zap.S().Errorf("failed to add comment: %+v", cmt)

		return &comment.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
			Comment: &comment.Comment{
				Id:         0,
				User:       0,
				Content:    "",
				CreateDate: "",
			},
		}, err
	}

	zap.S().Debugf("success to add comment: %+v", cmt)
	return &comment.CommentResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Comment: &comment.Comment{
			Id:         cmt.ID,
			User:       cmt.UserID,
			Content:    cmt.Content,
			CreateDate: time.Unix((cmt.ID>>22)/1000+1288834974657/1000, 0).Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (c *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	r := c.CommentFindByVideoID(req.VideoId)
	list := make([]*comment.Comment, 0)

	for i := range r {
		cmt := &comment.Comment{
			Id:         r[i].ID,
			User:       r[i].UserID,
			Content:    r[i].Content,
			CreateDate: time.Unix((r[i].ID>>22)/1000+1288834974657/1000, 0).Format("2006-01-02 15:04:05"),
		}
		list = append(list, cmt)
	}

	zap.S().Debugf("success to get comment list ==> len(comment_list): %v", len(list))
	return &comment.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: list,
	}, nil
}
