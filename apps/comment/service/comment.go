// Author: BeYoung
// Date: 2023/1/30 23:28
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/comment/dao"
	"github.com/Go-To-Byte/DouSheng/apps/comment/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/comment/models"
	"github.com/Go-To-Byte/DouSheng/apps/comment/proto"
	"go.uber.org/zap"
	"time"
)

func (c *Comment) Comment(ctx context.Context, req *proto.CommentRequest) (*proto.CommentResponse, error) {
	// 从 request 获取评论信息
	comment := model.Comment{
		ID:      models.Node.Generate().Int64(),
		VideoID: req.VideoId,
		UserID:  req.UserId,
		Content: req.Content,
	}

	// 添加评论
	if err := dao.Add(comment); err != nil {
		zap.S().Errorf("failed to add comment: %+v", comment)

		return &proto.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "failed",
			Comment: &proto.Comment{
				Id:         0,
				User:       0,
				Content:    "",
				CreateDate: "",
			},
		}, err
	}

	zap.S().Debugf("success to add comment: %+v", comment)
	return &proto.CommentResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Comment: &proto.Comment{
			Id:         comment.ID,
			User:       comment.UserID,
			Content:    comment.Content,
			CreateDate: time.Unix(comment.ID>>22, 0).Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (c *Comment) CommentList(ctx context.Context, req *proto.CommentListRequest) (*proto.CommentListResponse, error) {
	r := dao.CommentFindByVideoID(req.VideoId)
	list := make([]*proto.Comment, 0)

	for i := range r {
		comment := &proto.Comment{
			Id:         r[i].ID,
			User:       r[i].UserID,
			Content:    r[i].Content,
			CreateDate: time.Unix(r[i].ID>>22, 0).Format("2006-01-02 15:04:05"),
		}
		list = append(list, comment)
	}

	zap.S().Debugf("success to get comment list ==> len(comment_list): %v", len(list))
	return &proto.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: list,
	}, nil
}