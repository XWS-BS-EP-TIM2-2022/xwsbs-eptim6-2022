// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: posts_service/posts_service.proto

package posts_service

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

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetAllPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PostsResponse, error)
	GetAllPostsByUser(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*PostsResponse, error)
	GetPostById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*PostResponse, error)
	LikePost(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*PostResponse, error)
	DislikePost(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*PostResponse, error)
	AddNewPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	AddNewComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*PostResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetAllPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/GetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetAllPostsByUser(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/GetAllPostsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetPostById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/GetPostById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) LikePost(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/LikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) DislikePost(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/DislikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddNewPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/AddNewPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddNewComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/posts_service.ProfileService/AddNewComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	GetAllPosts(context.Context, *EmptyRequest) (*PostsResponse, error)
	GetAllPostsByUser(context.Context, *GetByUsernameRequest) (*PostsResponse, error)
	GetPostById(context.Context, *GetByIdRequest) (*PostResponse, error)
	LikePost(context.Context, *GetByIdRequest) (*PostResponse, error)
	DislikePost(context.Context, *GetByIdRequest) (*PostResponse, error)
	AddNewPost(context.Context, *PostRequest) (*PostResponse, error)
	AddNewComment(context.Context, *CommentRequest) (*PostResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) GetAllPosts(context.Context, *EmptyRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedProfileServiceServer) GetAllPostsByUser(context.Context, *GetByUsernameRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPostsByUser not implemented")
}
func (UnimplementedProfileServiceServer) GetPostById(context.Context, *GetByIdRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedProfileServiceServer) LikePost(context.Context, *GetByIdRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedProfileServiceServer) DislikePost(context.Context, *GetByIdRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DislikePost not implemented")
}
func (UnimplementedProfileServiceServer) AddNewPost(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPost not implemented")
}
func (UnimplementedProfileServiceServer) AddNewComment(context.Context, *CommentRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewComment not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/GetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetAllPosts(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetAllPostsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetAllPostsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/GetAllPostsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetAllPostsByUser(ctx, req.(*GetByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/GetPostById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetPostById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/LikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).LikePost(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_DislikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).DislikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/DislikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).DislikePost(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddNewPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddNewPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/AddNewPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddNewPost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddNewComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddNewComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_service.ProfileService/AddNewComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddNewComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posts_service.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPosts",
			Handler:    _ProfileService_GetAllPosts_Handler,
		},
		{
			MethodName: "GetAllPostsByUser",
			Handler:    _ProfileService_GetAllPostsByUser_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _ProfileService_GetPostById_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _ProfileService_LikePost_Handler,
		},
		{
			MethodName: "DislikePost",
			Handler:    _ProfileService_DislikePost_Handler,
		},
		{
			MethodName: "AddNewPost",
			Handler:    _ProfileService_AddNewPost_Handler,
		},
		{
			MethodName: "AddNewComment",
			Handler:    _ProfileService_AddNewComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts_service/posts_service.proto",
}
