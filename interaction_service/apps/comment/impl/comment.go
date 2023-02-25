// Created by yczbest at 2023/02/21 14:58

package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"
	"github.com/Go-To-Byte/DouSheng/interaction_service/apps/comment"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 实现评论操作功能
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
		userRsp, err := c.userService.UserInfo(ctx, userReq)
		fmt.Println(userRsp)
		if err != nil {
			return nil, err
		}
		commentVo := c.commentP2commentVo(ctx, po, userRsp.User)
		//commentVo := c.commentP2commentVo(ctx, po, &user_one)
		rsp, err := c.NewCommentActionResponse(ctx, commentVo)
		if err != nil {
			return nil, err
		}
		return rsp, nil
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
		commentVo := c.commentP2commentVo(ctx, po, userRsp.User)
		rsp, err := c.NewCommentActionResponse(ctx, commentVo)
		if err != nil {
			return nil, err
		}
		return rsp, nil
	}
	return nil, status.Error(codes.InvalidArgument,
		constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
}

// 获取评论列表，成功返回GetCommentListResponse对象地址，失败返回错误信息
func (c *commentServiceImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (*comment.GetCommentListResponse, error) {
	//	参数校验
	if err := req.Validate(); err != nil {
		c.l.Errorf("interaction: CommentAction 参数校验失败：%s", err.Error())
		return nil, status.Error(codes.InvalidArgument, constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}
	pos, err := c.GetCommentPoList(ctx, req)
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
	for index, po := range pos {
		userReq := user.UserInfoRequest{
			UserId: po.UserId,
		}
		userRsp, err := c.userService.UserInfo(ctx, &userReq)
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
	commentReq := comment.NewDefaultGetCommentCountByIdRequest()
	commentReq.VideoId = req.VideoId
	countRsp := comment.NewDefaultGetCommentCountByIdResponse()
	count, err := c.GetCommentCount(ctx, commentReq)
	if err != nil {
		c.l.Errorf("获取视频评论总数失败:%s", err.Error())
		return nil, err
	}
	countRsp.CommentCount = *count
	return countRsp, nil
}
