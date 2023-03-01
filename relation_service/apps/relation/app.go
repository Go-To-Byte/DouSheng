// @Author: Hexiaoming 2023/2/15
package relation

import (
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "relation"
)

var (
	validate = validator.New()
)

// NewUserFollowPo TODO UserFollowPo 可以不初始化Id吗
func NewUserFollowPo(req *FollowActionRequest) *UserFollowPo {
	return &UserFollowPo{
		UserId:     1,
		FollowId:   req.ToUserId,
		FollowFlag: 1,
	}
}

// NewUserFollowerPo TODO
func NewUserFollowerPo(req *FollowActionRequest) *UserFollowerPo {
	return &UserFollowerPo{
		UserId:       req.ToUserId,
		FollowerId:   1,
		FollowerFlag: 1,
	}
}

// Validate 获取关注列表 相关
func (r *FollowListRequest) Validate() error {
	return validate.Struct(r)
}

func NewFollowListRequest() *FollowListRequest {
	return &FollowListRequest{}
}

func NewFollowListResponse() *FollowListResponse {
	return &FollowListResponse{}
}

// Validate 获取粉丝列表 相关
func (r *FollowerListRequest) Validate() error {
	return validate.Struct(r)
}

func NewFollowerListRequest() *FollowerListRequest {
	return &FollowerListRequest{}
}

func NewFollowerListResponse() *FollowerListResponse {
	return &FollowerListResponse{}
}

// Validate 获取好友列表 相关
func (r *FriendListRequest) Validate() error {
	return validate.Struct(r)
}

func NewFriendListRequest() *FriendListRequest {
	return &FriendListRequest{}
}

func NewFriendListResponse() *FriendListResponse {
	return &FriendListResponse{}
}

// Validate 关注操作 相关
func (r *FollowActionRequest) Validate() error {
	return validate.Struct(r)
}

func NewFollowActionRequest() *FollowActionRequest {
	return &FollowActionRequest{}
}

func NewFollowActionResponse() *FollowActionResponse {
	return &FollowActionResponse{}
}

// TableName 指明表名 -> gorm 参数映射
func (*UserFollowPo) TableName() string {
	return "user_follow"
}

// TableName 指明表名 -> gorm 参数映射
func (*UserFollowerPo) TableName() string {
	return "user_follower"
}

func (r *ListCountRequest) Validate() error {
	return validate.Struct(r)
}

func NewListCountRequest() *ListCountRequest {
	return &ListCountRequest{}
}

func NewListCountResponse() *ListCountResponse {
	return &ListCountResponse{}
}
