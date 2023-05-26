// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
	"github.com/Go-To-Byte/DouSheng/user-service/apps/user"

	"github.com/Go-To-Byte/DouSheng/interaction-service/apps/comment"
)

// CommentAction 实现评论操作功能
func (c *commentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	//	参数校验
	if err := req.Validate(); err != nil {
		c.l.Errorf("interaction: CommentAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument, constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}
	//发表评论
	switch req.ActionType {
	case 1:
		//发布评论
		po, err := c.InsertCommentRecord(ctx, req)
		if err != nil {
			//c.l.Errorf("视频评论失败：%s", err.Error())
			return nil, status.Error(codes.PermissionDenied,
				constant.Code2Msg(constant.ERROR_SAVE))
		}
		userReq := user.NewUserInfoRequest()
		userReq.UserId = po.UserId
		userReq.Token = req.Token
		userRsp, err := c.userService.UserInfo(ctx, userReq)
		fmt.Println(userRsp)
		if err != nil {
			return nil, err
		}

		//commentVo := c.commentP2commentVo(ctx, po, &user_one)
		resp := comment.NewDefaultCommentActionResponse()
		resp.Comment = c.commentP2commentVo(ctx, po, userRsp.User)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case 2:
		po, err := c.DeleteCommentById(ctx, req)
		if err != nil {
			c.l.Errorf("删除视频评论失败：%s", err.Error())
			return nil, status.Error(codes.PermissionDenied,
				constant.Code2Msg(constant.WRONG_USER_NOT_EXIST))
		}
		userReq := user.NewUserInfoRequest()
		userReq.UserId = po.UserId
		userReq.Token = req.Token
		userRsp, err := c.userService.UserInfo(ctx, userReq)
		if err != nil {
			return nil, err
		}

		resp := comment.NewDefaultCommentActionResponse()
		resp.Comment = c.commentP2commentVo(ctx, po, userRsp.User)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, status.Error(codes.InvalidArgument,
		constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
}

// GetCommentList 获取评论列表，成功返回GetCommentListResponse对象地址，失败返回错误信息
func (c *commentServiceImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (*comment.GetCommentListResponse, error) {
	//	参数校验
	if err := req.Validate(); err != nil {
		c.l.Errorf("interaction: CommentAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument, constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}
	pos, err := c.getCommentPoList(ctx, req)
	fmt.Println()
	if err != nil {
		c.l.Errorf("获取视频评论列表失败：%s", err.Error())
		return nil, status.Error(codes.PermissionDenied,
			constant.Code2Msg(constant.BAD_REQUEST))
	}

	commentList := make([]*comment.Comment, len(pos))
	if len(commentList) == 0 {
		res := comment.NewDefaultGetCommentListResponse()
		res.CommentList = commentList
		return res, nil
	}

	// TODO：批量查询

	userReq := user.NewUserInfoRequest()
	userReq.Token = req.Token
	for index, po := range pos {
		userReq.UserId = po.UserId
		userRsp, err := c.userService.UserInfo(ctx, userReq)
		if err != nil {
			return nil, err
		}

		commentVo := c.commentP2commentVo(ctx, po, userRsp.User)
		commentList[index] = commentVo
	}
	res := comment.GetCommentListResponse{
		CommentList: commentList,
	}
	return &res, nil
}

func (s *commentServiceImpl) CommentCountMap(ctx context.Context, req *comment.CommentMapRequest) (
	*comment.CommentMapResponse, error) {
	// 1、参数校验
	resp := comment.NewCommentMapResponse()
	if req.VideoIds == nil || len(req.VideoIds) <= 0 {
		s.l.Errorf("user userList：你的参数可能有问题哟~")
		return resp, nil
	}

	// 2、获取每个视频的点赞数
	resp.CommentCountMap = make(map[int64]int64, 10)
	for _, v := range req.VideoIds {
		// 获取点赞数

		count, err := s.getCommentCount(ctx, v)

		if err != nil {
			return resp, status.Errorf(codes.Unavailable, constant.Code2Msg(constant.ERROR_ACQUIRE))
		}

		resp.CommentCountMap[v] = count
	}

	return resp, nil
}

func (c *commentServiceImpl) commentP2commentVo(ctx context.Context, po *comment.CommentPo, userInfo *user.User) *comment.Comment {
	//根据id查询用户信息
	commentVo := comment.Comment{}
	commentVo.Id = po.Id
	commentVo.Content = po.Content
	commentVo.User = userInfo
	commentVo.CreateDate = po.CreateDate
	return &commentVo
}

func (c *commentServiceImpl) NewCommentActionResponse(ctx context.Context, vo *comment.Comment) (*comment.CommentActionResponse, error) {
	res := comment.NewDefaultCommentActionResponse()
	res.Comment = vo
	return res, nil
}

func (c *commentServiceImpl) GetCommentCountById(ctx context.Context, req *comment.GetCommentCountByIdRequest) (*comment.GetCommentCountByIdResponse, error) {
	//	参数校验
	if err := req.Validate(); err != nil {
		c.l.Errorf("interaction: CommentAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument, constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	count, err := c.getCommentCount(ctx, req.VideoId)
	if err != nil {
		c.l.Errorf("获取视频评论总数失败:%s", err.Error())
		return nil, err
	}
	countRsp := comment.NewDefaultGetCommentCountByIdResponse()
	countRsp.CommentCount = count
	return countRsp, nil
}
