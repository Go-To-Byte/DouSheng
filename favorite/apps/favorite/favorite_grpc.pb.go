// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: apps/favorite/pb/favorite.proto

package favorite

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FavoriteClient is the client API for Favorite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteClient interface {
	Favorite(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*FavoriteResponse, error)
	FavoredList(ctx context.Context, in *FavoredListRequest, opts ...grpc.CallOption) (*FavoredListResponse, error)
	FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error)
	FavoriteJudge(ctx context.Context, in *FavoriteJudgeRequest, opts ...grpc.CallOption) (*FavoriteJudgeResponse, error)
}

type favoriteClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteClient(cc grpc.ClientConnInterface) FavoriteClient {
	return &favoriteClient{cc}
}

func (c *favoriteClient) Favorite(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*FavoriteResponse, error) {
	out := new(FavoriteResponse)
	err := c.cc.Invoke(ctx, "/dousheng.favorite.favorite/favorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) FavoredList(ctx context.Context, in *FavoredListRequest, opts ...grpc.CallOption) (*FavoredListResponse, error) {
	out := new(FavoredListResponse)
	err := c.cc.Invoke(ctx, "/dousheng.favorite.favorite/favored_list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error) {
	out := new(FavoriteListResponse)
	err := c.cc.Invoke(ctx, "/dousheng.favorite.favorite/favorite_list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) FavoriteJudge(ctx context.Context, in *FavoriteJudgeRequest, opts ...grpc.CallOption) (*FavoriteJudgeResponse, error) {
	out := new(FavoriteJudgeResponse)
	err := c.cc.Invoke(ctx, "/dousheng.favorite.favorite/favorite_judge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServer is the server API for Favorite service.
// All implementations must embed UnimplementedFavoriteServer
// for forward compatibility
type FavoriteServer interface {
	Favorite(context.Context, *FavoriteRequest) (*FavoriteResponse, error)
	FavoredList(context.Context, *FavoredListRequest) (*FavoredListResponse, error)
	FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error)
	FavoriteJudge(context.Context, *FavoriteJudgeRequest) (*FavoriteJudgeResponse, error)
	mustEmbedUnimplementedFavoriteServer()
}

// UnimplementedFavoriteServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServer struct {
}

func (UnimplementedFavoriteServer) Favorite(context.Context, *FavoriteRequest) (*FavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Favorite not implemented")
}
func (UnimplementedFavoriteServer) FavoredList(context.Context, *FavoredListRequest) (*FavoredListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoredList not implemented")
}
func (UnimplementedFavoriteServer) FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedFavoriteServer) FavoriteJudge(context.Context, *FavoriteJudgeRequest) (*FavoriteJudgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteJudge not implemented")
}
func (UnimplementedFavoriteServer) mustEmbedUnimplementedFavoriteServer() {}

// UnsafeFavoriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServer will
// result in compilation errors.
type UnsafeFavoriteServer interface {
	mustEmbedUnimplementedFavoriteServer()
}

func RegisterFavoriteServer(s grpc.ServiceRegistrar, srv FavoriteServer) {
	s.RegisterService(&Favorite_ServiceDesc, srv)
}

func _Favorite_Favorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).Favorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.favorite.favorite/favorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).Favorite(ctx, req.(*FavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_FavoredList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoredListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoredList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.favorite.favorite/favored_list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoredList(ctx, req.(*FavoredListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.favorite.favorite/favorite_list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoriteList(ctx, req.(*FavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_FavoriteJudge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteJudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoriteJudge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.favorite.favorite/favorite_judge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoriteJudge(ctx, req.(*FavoriteJudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Favorite_ServiceDesc is the grpc.ServiceDesc for Favorite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Favorite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dousheng.favorite.favorite",
	HandlerType: (*FavoriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "favorite",
			Handler:    _Favorite_Favorite_Handler,
		},
		{
			MethodName: "favored_list",
			Handler:    _Favorite_FavoredList_Handler,
		},
		{
			MethodName: "favorite_list",
			Handler:    _Favorite_FavoriteList_Handler,
		},
		{
			MethodName: "favorite_judge",
			Handler:    _Favorite_FavoriteJudge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/favorite/pb/favorite.proto",
}