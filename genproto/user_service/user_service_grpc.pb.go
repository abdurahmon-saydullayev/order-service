// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*UserPrimaryKey, error)
	GetByID(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*User, error)
	GetList(ctx context.Context, in *GetListUserRequest, opts ...grpc.CallOption) (*GetListUserResponse, error)
	Update(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*User, error)
	UpdatePatch(ctx context.Context, in *UpdatePatchUser, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateUserOTP(ctx context.Context, in *CreateOTP, opts ...grpc.CallOption) (*empty.Empty, error)
	VerifyUserOTP(ctx context.Context, in *VerifyOTP, opts ...grpc.CallOption) (*empty.Empty, error)
	Check(ctx context.Context, in *UserPhoneNumberReq, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*UserPrimaryKey, error) {
	out := new(UserPrimaryKey)
	err := c.cc.Invoke(ctx, "/user_service.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByID(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetList(ctx context.Context, in *GetListUserRequest, opts ...grpc.CallOption) (*GetListUserResponse, error) {
	out := new(GetListUserResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePatch(ctx context.Context, in *UpdatePatchUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.UserService/UpdatePatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_service.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUserOTP(ctx context.Context, in *CreateOTP, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_service.UserService/CreateUserOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyUserOTP(ctx context.Context, in *VerifyOTP, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_service.UserService/VerifyUserOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Check(ctx context.Context, in *UserPhoneNumberReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.UserService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Create(context.Context, *CreateUser) (*UserPrimaryKey, error)
	GetByID(context.Context, *UserPrimaryKey) (*User, error)
	GetList(context.Context, *GetListUserRequest) (*GetListUserResponse, error)
	Update(context.Context, *UpdateUser) (*User, error)
	UpdatePatch(context.Context, *UpdatePatchUser) (*User, error)
	Delete(context.Context, *UserPrimaryKey) (*empty.Empty, error)
	CreateUserOTP(context.Context, *CreateOTP) (*empty.Empty, error)
	VerifyUserOTP(context.Context, *VerifyOTP) (*empty.Empty, error)
	Check(context.Context, *UserPhoneNumberReq) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Create(context.Context, *CreateUser) (*UserPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) GetByID(context.Context, *UserPrimaryKey) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServiceServer) GetList(context.Context, *GetListUserRequest) (*GetListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) UpdatePatch(context.Context, *UpdatePatchUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatch not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *UserPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) CreateUserOTP(context.Context, *CreateOTP) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserOTP not implemented")
}
func (UnimplementedUserServiceServer) VerifyUserOTP(context.Context, *VerifyOTP) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserOTP not implemented")
}
func (UnimplementedUserServiceServer) Check(context.Context, *UserPhoneNumberReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByID(ctx, req.(*UserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetList(ctx, req.(*GetListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatchUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/UpdatePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePatch(ctx, req.(*UpdatePatchUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUserOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOTP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/CreateUserOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserOTP(ctx, req.(*CreateOTP))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyUserOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyUserOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/VerifyUserOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyUserOTP(ctx, req.(*VerifyOTP))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPhoneNumberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Check(ctx, req.(*UserPhoneNumberReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _UserService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _UserService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "UpdatePatch",
			Handler:    _UserService_UpdatePatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "CreateUserOTP",
			Handler:    _UserService_CreateUserOTP_Handler,
		},
		{
			MethodName: "VerifyUserOTP",
			Handler:    _UserService_VerifyUserOTP_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _UserService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}