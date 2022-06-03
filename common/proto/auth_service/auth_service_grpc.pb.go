// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: auth_service/auth_service.proto

package auth_service

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	AddNewUser(ctx context.Context, in *CreateNewUser, opts ...grpc.CallOption) (*CreateNewUser, error)
	UpdateUserPassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*CreateNewUser, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	LoginUser(ctx context.Context, in *CreateNewUser, opts ...grpc.CallOption) (*Token, error)
	AuthorizeJWT(ctx context.Context, in *ValidateToken, opts ...grpc.CallOption) (*CreateNewUser, error)
	GetUserPermissions(ctx context.Context, in *ValidateToken, opts ...grpc.CallOption) (*UserPermissions, error)
	ActivateUserAccount(ctx context.Context, in *ActivationToken, opts ...grpc.CallOption) (*ActivationResponse, error)
	ForgottenPassword(ctx context.Context, in *UserEmailMessage, opts ...grpc.CallOption) (*ActivationResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordWithTokenMessage, opts ...grpc.CallOption) (*ActivationResponse, error)
	GeneratePasswordlessLoginToken(ctx context.Context, in *UserEmailMessage, opts ...grpc.CallOption) (*ActivationResponse, error)
	PasswordlessLogin(ctx context.Context, in *ActivationTokenMessage, opts ...grpc.CallOption) (*ActivationResponse, error)
	GenerateApiKey(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*Token, error)
	GetUserApiKey(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*Token, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AddNewUser(ctx context.Context, in *CreateNewUser, opts ...grpc.CallOption) (*CreateNewUser, error) {
	out := new(CreateNewUser)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/AddNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserPassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*CreateNewUser, error) {
	out := new(CreateNewUser)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/UpdateUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LoginUser(ctx context.Context, in *CreateNewUser, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthorizeJWT(ctx context.Context, in *ValidateToken, opts ...grpc.CallOption) (*CreateNewUser, error) {
	out := new(CreateNewUser)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/AuthorizeJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserPermissions(ctx context.Context, in *ValidateToken, opts ...grpc.CallOption) (*UserPermissions, error) {
	out := new(UserPermissions)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/GetUserPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ActivateUserAccount(ctx context.Context, in *ActivationToken, opts ...grpc.CallOption) (*ActivationResponse, error) {
	out := new(ActivationResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/ActivateUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgottenPassword(ctx context.Context, in *UserEmailMessage, opts ...grpc.CallOption) (*ActivationResponse, error) {
	out := new(ActivationResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/ForgottenPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordWithTokenMessage, opts ...grpc.CallOption) (*ActivationResponse, error) {
	out := new(ActivationResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GeneratePasswordlessLoginToken(ctx context.Context, in *UserEmailMessage, opts ...grpc.CallOption) (*ActivationResponse, error) {
	out := new(ActivationResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/GeneratePasswordlessLoginToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PasswordlessLogin(ctx context.Context, in *ActivationTokenMessage, opts ...grpc.CallOption) (*ActivationResponse, error) {
	out := new(ActivationResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/PasswordlessLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateApiKey(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/GenerateApiKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserApiKey(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/GetUserApiKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	AddNewUser(context.Context, *CreateNewUser) (*CreateNewUser, error)
	UpdateUserPassword(context.Context, *ChangePasswordRequest) (*CreateNewUser, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	LoginUser(context.Context, *CreateNewUser) (*Token, error)
	AuthorizeJWT(context.Context, *ValidateToken) (*CreateNewUser, error)
	GetUserPermissions(context.Context, *ValidateToken) (*UserPermissions, error)
	ActivateUserAccount(context.Context, *ActivationToken) (*ActivationResponse, error)
	ForgottenPassword(context.Context, *UserEmailMessage) (*ActivationResponse, error)
	ResetPassword(context.Context, *ResetPasswordWithTokenMessage) (*ActivationResponse, error)
	GeneratePasswordlessLoginToken(context.Context, *UserEmailMessage) (*ActivationResponse, error)
	PasswordlessLogin(context.Context, *ActivationTokenMessage) (*ActivationResponse, error)
	GenerateApiKey(context.Context, *GetAllRequest) (*Token, error)
	GetUserApiKey(context.Context, *GetAllRequest) (*Token, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AddNewUser(context.Context, *CreateNewUser) (*CreateNewUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserPassword(context.Context, *ChangePasswordRequest) (*CreateNewUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
func (UnimplementedAuthServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAuthServiceServer) LoginUser(context.Context, *CreateNewUser) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthServiceServer) AuthorizeJWT(context.Context, *ValidateToken) (*CreateNewUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeJWT not implemented")
}
func (UnimplementedAuthServiceServer) GetUserPermissions(context.Context, *ValidateToken) (*UserPermissions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPermissions not implemented")
}
func (UnimplementedAuthServiceServer) ActivateUserAccount(context.Context, *ActivationToken) (*ActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUserAccount not implemented")
}
func (UnimplementedAuthServiceServer) ForgottenPassword(context.Context, *UserEmailMessage) (*ActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgottenPassword not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *ResetPasswordWithTokenMessage) (*ActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) GeneratePasswordlessLoginToken(context.Context, *UserEmailMessage) (*ActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePasswordlessLoginToken not implemented")
}
func (UnimplementedAuthServiceServer) PasswordlessLogin(context.Context, *ActivationTokenMessage) (*ActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordlessLogin not implemented")
}
func (UnimplementedAuthServiceServer) GenerateApiKey(context.Context, *GetAllRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateApiKey not implemented")
}
func (UnimplementedAuthServiceServer) GetUserApiKey(context.Context, *GetAllRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserApiKey not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AddNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/AddNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddNewUser(ctx, req.(*CreateNewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/UpdateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserPassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoginUser(ctx, req.(*CreateNewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthorizeJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthorizeJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/AuthorizeJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthorizeJWT(ctx, req.(*ValidateToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/GetUserPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserPermissions(ctx, req.(*ValidateToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ActivateUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ActivateUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/ActivateUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ActivateUserAccount(ctx, req.(*ActivationToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgottenPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmailMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgottenPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/ForgottenPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgottenPassword(ctx, req.(*UserEmailMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordWithTokenMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*ResetPasswordWithTokenMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GeneratePasswordlessLoginToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmailMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GeneratePasswordlessLoginToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/GeneratePasswordlessLoginToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GeneratePasswordlessLoginToken(ctx, req.(*UserEmailMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PasswordlessLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationTokenMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PasswordlessLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/PasswordlessLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PasswordlessLogin(ctx, req.(*ActivationTokenMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/GenerateApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateApiKey(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/GetUserApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserApiKey(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewUser",
			Handler:    _AuthService_AddNewUser_Handler,
		},
		{
			MethodName: "UpdateUserPassword",
			Handler:    _AuthService_UpdateUserPassword_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AuthService_GetAll_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AuthService_LoginUser_Handler,
		},
		{
			MethodName: "AuthorizeJWT",
			Handler:    _AuthService_AuthorizeJWT_Handler,
		},
		{
			MethodName: "GetUserPermissions",
			Handler:    _AuthService_GetUserPermissions_Handler,
		},
		{
			MethodName: "ActivateUserAccount",
			Handler:    _AuthService_ActivateUserAccount_Handler,
		},
		{
			MethodName: "ForgottenPassword",
			Handler:    _AuthService_ForgottenPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "GeneratePasswordlessLoginToken",
			Handler:    _AuthService_GeneratePasswordlessLoginToken_Handler,
		},
		{
			MethodName: "PasswordlessLogin",
			Handler:    _AuthService_PasswordlessLogin_Handler,
		},
		{
			MethodName: "GenerateApiKey",
			Handler:    _AuthService_GenerateApiKey_Handler,
		},
		{
			MethodName: "GetUserApiKey",
			Handler:    _AuthService_GetUserApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service/auth_service.proto",
}
