// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: profile_service/profile_service.proto

package profile_service

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
	GetAllUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	AddNewUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AddSkill(ctx context.Context, in *SkillRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AddInterest(ctx context.Context, in *InterestRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AddExperience(ctx context.Context, in *ExperienceRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AddEducation(ctx context.Context, in *EducationRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UnFollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AcceptFollow(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	RejectFollow(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetAllUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddNewUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/AddNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddSkill(ctx context.Context, in *SkillRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/AddSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddInterest(ctx context.Context, in *InterestRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/AddInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddExperience(ctx context.Context, in *ExperienceRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/AddExperience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddEducation(ctx context.Context, in *EducationRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/AddEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UnFollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/UnFollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AcceptFollow(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/AcceptFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) RejectFollow(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/profile_service.ProfileService/RejectFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	GetAllUsers(context.Context, *EmptyRequest) (*UsersResponse, error)
	AddNewUser(context.Context, *UserRequest) (*UserResponse, error)
	AddSkill(context.Context, *SkillRequest) (*UserResponse, error)
	AddInterest(context.Context, *InterestRequest) (*UserResponse, error)
	AddExperience(context.Context, *ExperienceRequest) (*UserResponse, error)
	AddEducation(context.Context, *EducationRequest) (*UserResponse, error)
	FollowUser(context.Context, *FollowUserRequest) (*UserResponse, error)
	UnFollowUser(context.Context, *FollowUserRequest) (*UserResponse, error)
	AcceptFollow(context.Context, *FollowUserRequest) (*UserResponse, error)
	RejectFollow(context.Context, *FollowUserRequest) (*UserResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) GetAllUsers(context.Context, *EmptyRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedProfileServiceServer) AddNewUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewUser not implemented")
}
func (UnimplementedProfileServiceServer) AddSkill(context.Context, *SkillRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSkill not implemented")
}
func (UnimplementedProfileServiceServer) AddInterest(context.Context, *InterestRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInterest not implemented")
}
func (UnimplementedProfileServiceServer) AddExperience(context.Context, *ExperienceRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExperience not implemented")
}
func (UnimplementedProfileServiceServer) AddEducation(context.Context, *EducationRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEducation not implemented")
}
func (UnimplementedProfileServiceServer) FollowUser(context.Context, *FollowUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedProfileServiceServer) UnFollowUser(context.Context, *FollowUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (UnimplementedProfileServiceServer) AcceptFollow(context.Context, *FollowUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFollow not implemented")
}
func (UnimplementedProfileServiceServer) RejectFollow(context.Context, *FollowUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectFollow not implemented")
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

func _ProfileService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetAllUsers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/AddNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddNewUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/AddSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddSkill(ctx, req.(*SkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/AddInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddInterest(ctx, req.(*InterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/AddExperience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddExperience(ctx, req.(*ExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/AddEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddEducation(ctx, req.(*EducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UnFollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UnFollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/UnFollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UnFollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AcceptFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AcceptFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/AcceptFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AcceptFollow(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_RejectFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).RejectFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_service.ProfileService/RejectFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).RejectFollow(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile_service.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsers",
			Handler:    _ProfileService_GetAllUsers_Handler,
		},
		{
			MethodName: "AddNewUser",
			Handler:    _ProfileService_AddNewUser_Handler,
		},
		{
			MethodName: "AddSkill",
			Handler:    _ProfileService_AddSkill_Handler,
		},
		{
			MethodName: "AddInterest",
			Handler:    _ProfileService_AddInterest_Handler,
		},
		{
			MethodName: "AddExperience",
			Handler:    _ProfileService_AddExperience_Handler,
		},
		{
			MethodName: "AddEducation",
			Handler:    _ProfileService_AddEducation_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _ProfileService_FollowUser_Handler,
		},
		{
			MethodName: "UnFollowUser",
			Handler:    _ProfileService_UnFollowUser_Handler,
		},
		{
			MethodName: "AcceptFollow",
			Handler:    _ProfileService_AcceptFollow_Handler,
		},
		{
			MethodName: "RejectFollow",
			Handler:    _ProfileService_RejectFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile_service/profile_service.proto",
}
