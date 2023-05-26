package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Go-To-Byte/DouSheng/user-service/apps/relation"
)

// TODO：实现关系接口

func (s *relationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListRequest) (
	*relation.FollowListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}

func (s *relationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (
	*relation.FollowerListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}

func (s *relationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListRequest) (
	*relation.FriendListResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}

func (s *relationServiceImpl) FollowAction(ctx context.Context, req *relation.FollowActionRequest) (
	*relation.FollowActionResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}

func (s *relationServiceImpl) IsFollower(ctx context.Context, req *relation.UserFollowPo) (
	*relation.IsFollowerResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}
