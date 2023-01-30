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
)

func (c *Comment) Favorite(ctx context.Context, req *proto.CommentRequest) (*proto.CommentResponse, error) {
	comment := model.Comment{
		ID:      models.Node.Generate().Int64(),
		VideoID: req.VideoId,
		UserID:  req.UserId,
		Comment: req.Content,
	}
	if err := dao.Add(comment); err != nil {
		return &proto.CommentResponse{
			StatusCode: 0,
			StatusMsg:  "",
			Comment:    req.Content,
		}, nil
		zap.S().Errorf("failed to add comment: %+v", comment)
	}
}

func (c *Comment) FavoriteList(ctx context.Context, req *proto.CommentListRequest) (*proto.CommentListResponse, error) {

}
