// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apps/video/pb/video.proto

package video

import (
	user "github.com/Go-To-Byte/DouSheng/apps/user"
	common "github.com/Go-To-Byte/DouSheng/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 视频唯一标识
	// @gotags: json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 视频作者信息
	// @gotags: json:"author"
	Author *user.User `protobuf:"bytes,2,opt,name=author,proto3" json:"author"`
	// 视频播放地址
	// @gotags: json:"play_url"
	PlayUrl string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url"`
	// 视频封面地址
	// @gotags: json:"cover_url"
	CoverUrl string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url"`
	// 视频的点赞总数
	// @gotags: json:"favorite_count"
	FavoriteCount int64 `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count"`
	// 视频的评论总数
	// @gotags: json:"comment_count"
	CommentCount int64 `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count"`
	// true-已点赞,false-未点赞
	// @gotags: json:"is_favorite"
	IsFavorite bool `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite"`
	// 视频标题
	// @gotags: json:"title"
	Title string `protobuf:"bytes,8,opt,name=title,proto3" json:"title"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *user.User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FeedVideosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *common.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	// @gotags: json:"latest_time"
	LatestTime *int64 `protobuf:"varint,2,opt,name=latest_time,json=latestTime,proto3,oneof" json:"latest_time"`
	// 可选参数，登录用户设置
	// @gotags: json:"token"
	Token *string `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token"`
}

func (x *FeedVideosRequest) Reset() {
	*x = FeedVideosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedVideosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedVideosRequest) ProtoMessage() {}

func (x *FeedVideosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedVideosRequest.ProtoReflect.Descriptor instead.
func (*FeedVideosRequest) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{1}
}

func (x *FeedVideosRequest) GetPage() *common.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FeedVideosRequest) GetLatestTime() int64 {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return 0
}

func (x *FeedVideosRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type PublishVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户鉴权token
	// @gotags: json:"token" validate:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token" validate:"token"`
	// 视频数据
	// @gotags: json:"data" validate:"data"
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data" validate:"data"`
	// 视频标题
	// @gotags: json:"title" validate:"title"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title" validate:"title"`
}

func (x *PublishVideoRequest) Reset() {
	*x = PublishVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoRequest) ProtoMessage() {}

func (x *PublishVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoRequest.ProtoReflect.Descriptor instead.
func (*PublishVideoRequest) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{2}
}

func (x *PublishVideoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PublishVideoRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PublishVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FeedSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码，0-成功， 其他值-失败
	// @gotags: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// 返回状态描述信息
	// @gotags: json:"status_msg"
	StatusMsg *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg"`
	// 视频列表
	// @gotags: json:"video_list"
	VideoList []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list"`
	// 本次返回的视频中，发布最早的时间，作为下次请求的latest_time
	// @gotags: json:"next_time"
	NextTime *int64 `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time"`
}

func (x *FeedSetResponse) Reset() {
	*x = FeedSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedSetResponse) ProtoMessage() {}

func (x *FeedSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedSetResponse.ProtoReflect.Descriptor instead.
func (*FeedSetResponse) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{3}
}

func (x *FeedSetResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FeedSetResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *FeedSetResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *FeedSetResponse) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

type PublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码,0-成功,其其他值-失败
	// @gotags: json:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	// 返回状态描述
	// @gotags: json:"status_msg"
	StatusMsg *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg"`
	// 用户发布的视频列表
	// @gotags: json:"video_list"
	VideoList []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list"`
}

func (x *PublishListResponse) Reset() {
	*x = PublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListResponse) ProtoMessage() {}

func (x *PublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListResponse.ProtoReflect.Descriptor instead.
func (*PublishListResponse) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{4}
}

func (x *PublishListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PublishListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *PublishListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_apps_video_pb_video_proto protoreflect.FileDescriptor

var file_apps_video_pb_video_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x5f,
	0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x36, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x46, 0x65, 0x65, 0x64, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x55, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x0f, 0x46, 0x65,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12,
	0x3f, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x32,
	0xd3, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0a, 0x46,
	0x65, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x5f, 0x74,
	0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x2d, 0x54, 0x6f, 0x2d, 0x42, 0x79, 0x74, 0x65, 0x2f, 0x44,
	0x6f, 0x75, 0x53, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_video_pb_video_proto_rawDescOnce sync.Once
	file_apps_video_pb_video_proto_rawDescData = file_apps_video_pb_video_proto_rawDesc
)

func file_apps_video_pb_video_proto_rawDescGZIP() []byte {
	file_apps_video_pb_video_proto_rawDescOnce.Do(func() {
		file_apps_video_pb_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_video_pb_video_proto_rawDescData)
	})
	return file_apps_video_pb_video_proto_rawDescData
}

var file_apps_video_pb_video_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_video_pb_video_proto_goTypes = []interface{}{
	(*Video)(nil),                        // 0: go_to_byte.dousheng.video.Video
	(*FeedVideosRequest)(nil),            // 1: go_to_byte.dousheng.video.FeedVideosRequest
	(*PublishVideoRequest)(nil),          // 2: go_to_byte.dousheng.video.PublishVideoRequest
	(*FeedSetResponse)(nil),              // 3: go_to_byte.dousheng.video.FeedSetResponse
	(*PublishListResponse)(nil),          // 4: go_to_byte.dousheng.video.PublishListResponse
	(*user.User)(nil),                    // 5: go_to_byte.dousheng.user.User
	(*common.PageRequest)(nil),           // 6: go_to_byte.dousheng.common.PageRequest
	(*common.UserIDAndTokenRequest)(nil), // 7: go_to_byte.dousheng.common.UserIDAndTokenRequest
	(*common.CodeAndMsgResponse)(nil),    // 8: go_to_byte.dousheng.common.CodeAndMsgResponse
}
var file_apps_video_pb_video_proto_depIdxs = []int32{
	5, // 0: go_to_byte.dousheng.video.Video.author:type_name -> go_to_byte.dousheng.user.User
	6, // 1: go_to_byte.dousheng.video.FeedVideosRequest.page:type_name -> go_to_byte.dousheng.common.PageRequest
	0, // 2: go_to_byte.dousheng.video.FeedSetResponse.video_list:type_name -> go_to_byte.dousheng.video.Video
	0, // 3: go_to_byte.dousheng.video.PublishListResponse.video_list:type_name -> go_to_byte.dousheng.video.Video
	1, // 4: go_to_byte.dousheng.video.Service.FeedVideos:input_type -> go_to_byte.dousheng.video.FeedVideosRequest
	2, // 5: go_to_byte.dousheng.video.Service.PublishVideo:input_type -> go_to_byte.dousheng.video.PublishVideoRequest
	7, // 6: go_to_byte.dousheng.video.Service.PublishList:input_type -> go_to_byte.dousheng.common.UserIDAndTokenRequest
	3, // 7: go_to_byte.dousheng.video.Service.FeedVideos:output_type -> go_to_byte.dousheng.video.FeedSetResponse
	8, // 8: go_to_byte.dousheng.video.Service.PublishVideo:output_type -> go_to_byte.dousheng.common.CodeAndMsgResponse
	4, // 9: go_to_byte.dousheng.video.Service.PublishList:output_type -> go_to_byte.dousheng.video.PublishListResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_video_pb_video_proto_init() }
func file_apps_video_pb_video_proto_init() {
	if File_apps_video_pb_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_video_pb_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_video_pb_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedVideosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_video_pb_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_video_pb_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedSetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_video_pb_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_apps_video_pb_video_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_apps_video_pb_video_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_apps_video_pb_video_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_video_pb_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_video_pb_video_proto_goTypes,
		DependencyIndexes: file_apps_video_pb_video_proto_depIdxs,
		MessageInfos:      file_apps_video_pb_video_proto_msgTypes,
	}.Build()
	File_apps_video_pb_video_proto = out.File
	file_apps_video_pb_video_proto_rawDesc = nil
	file_apps_video_pb_video_proto_goTypes = nil
	file_apps_video_pb_video_proto_depIdxs = nil
}