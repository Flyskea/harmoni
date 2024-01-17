// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/like/grpc/v1/like.proto

package v1

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

const (
	Like_Like_FullMethodName         = "/like.Like/Like"
	Like_UserLikeList_FullMethodName = "/like.Like/UserLikeList"
)

// LikeClient is the client API for Like service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeClient interface {
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeReply, error)
	UserLikeList(ctx context.Context, in *UserLikeListRequest, opts ...grpc.CallOption) (*UserLikeListReply, error)
}

type likeClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeClient(cc grpc.ClientConnInterface) LikeClient {
	return &likeClient{cc}
}

func (c *likeClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeReply, error) {
	out := new(LikeReply)
	err := c.cc.Invoke(ctx, Like_Like_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) UserLikeList(ctx context.Context, in *UserLikeListRequest, opts ...grpc.CallOption) (*UserLikeListReply, error) {
	out := new(UserLikeListReply)
	err := c.cc.Invoke(ctx, Like_UserLikeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServer is the server API for Like service.
// All implementations must embed UnimplementedLikeServer
// for forward compatibility
type LikeServer interface {
	Like(context.Context, *LikeRequest) (*LikeReply, error)
	UserLikeList(context.Context, *UserLikeListRequest) (*UserLikeListReply, error)
	mustEmbedUnimplementedLikeServer()
}

// UnimplementedLikeServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServer struct {
}

func (UnimplementedLikeServer) Like(context.Context, *LikeRequest) (*LikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedLikeServer) UserLikeList(context.Context, *UserLikeListRequest) (*UserLikeListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLikeList not implemented")
}
func (UnimplementedLikeServer) mustEmbedUnimplementedLikeServer() {}

// UnsafeLikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServer will
// result in compilation errors.
type UnsafeLikeServer interface {
	mustEmbedUnimplementedLikeServer()
}

func RegisterLikeServer(s grpc.ServiceRegistrar, srv LikeServer) {
	s.RegisterService(&Like_ServiceDesc, srv)
}

func _Like_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_UserLikeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLikeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).UserLikeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_UserLikeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).UserLikeList(ctx, req.(*UserLikeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Like_ServiceDesc is the grpc.ServiceDesc for Like service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Like_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "like.Like",
	HandlerType: (*LikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Like",
			Handler:    _Like_Like_Handler,
		},
		{
			MethodName: "UserLikeList",
			Handler:    _Like_UserLikeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/like/grpc/v1/like.proto",
}
