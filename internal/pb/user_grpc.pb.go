// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: user.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserServices_Signup_FullMethodName            = "/pbU.UserServices/Signup"
	UserServices_VerifyOTP_FullMethodName         = "/pbU.UserServices/VerifyOTP"
	UserServices_Login_FullMethodName             = "/pbU.UserServices/Login"
	UserServices_UserProductList_FullMethodName   = "/pbU.UserServices/UserProductList"
	UserServices_UserProductByName_FullMethodName = "/pbU.UserServices/UserProductByName"
	UserServices_UserProductByID_FullMethodName   = "/pbU.UserServices/UserProductByID"
)

// UserServicesClient is the client API for UserServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServicesClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UserProductList(ctx context.Context, in *RNoParam, opts ...grpc.CallOption) (*ProductList, error)
	UserProductByName(ctx context.Context, in *ProductByName, opts ...grpc.CallOption) (*ProductDetails, error)
	UserProductByID(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*ProductDetails, error)
}

type userServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServicesClient(cc grpc.ClientConnInterface) UserServicesClient {
	return &userServicesClient{cc}
}

func (c *userServicesClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, UserServices_Signup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, UserServices_VerifyOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserServices_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UserProductList(ctx context.Context, in *RNoParam, opts ...grpc.CallOption) (*ProductList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductList)
	err := c.cc.Invoke(ctx, UserServices_UserProductList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UserProductByName(ctx context.Context, in *ProductByName, opts ...grpc.CallOption) (*ProductDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductDetails)
	err := c.cc.Invoke(ctx, UserServices_UserProductByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UserProductByID(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*ProductDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductDetails)
	err := c.cc.Invoke(ctx, UserServices_UserProductByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServicesServer is the server API for UserServices service.
// All implementations must embed UnimplementedUserServicesServer
// for forward compatibility.
type UserServicesServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UserProductList(context.Context, *RNoParam) (*ProductList, error)
	UserProductByName(context.Context, *ProductByName) (*ProductDetails, error)
	UserProductByID(context.Context, *ProductID) (*ProductDetails, error)
	mustEmbedUnimplementedUserServicesServer()
}

// UnimplementedUserServicesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServicesServer struct{}

func (UnimplementedUserServicesServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServicesServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedUserServicesServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServicesServer) UserProductList(context.Context, *RNoParam) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserProductList not implemented")
}
func (UnimplementedUserServicesServer) UserProductByName(context.Context, *ProductByName) (*ProductDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserProductByName not implemented")
}
func (UnimplementedUserServicesServer) UserProductByID(context.Context, *ProductID) (*ProductDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserProductByID not implemented")
}
func (UnimplementedUserServicesServer) mustEmbedUnimplementedUserServicesServer() {}
func (UnimplementedUserServicesServer) testEmbeddedByValue()                      {}

// UnsafeUserServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServicesServer will
// result in compilation errors.
type UnsafeUserServicesServer interface {
	mustEmbedUnimplementedUserServicesServer()
}

func RegisterUserServicesServer(s grpc.ServiceRegistrar, srv UserServicesServer) {
	// If the following call pancis, it indicates UnimplementedUserServicesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserServices_ServiceDesc, srv)
}

func _UserServices_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UserProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UserProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UserProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UserProductList(ctx, req.(*RNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UserProductByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UserProductByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UserProductByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UserProductByName(ctx, req.(*ProductByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UserProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UserProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UserProductByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UserProductByID(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServices_ServiceDesc is the grpc.ServiceDesc for UserServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbU.UserServices",
	HandlerType: (*UserServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserServices_Signup_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _UserServices_VerifyOTP_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserServices_Login_Handler,
		},
		{
			MethodName: "UserProductList",
			Handler:    _UserServices_UserProductList_Handler,
		},
		{
			MethodName: "UserProductByName",
			Handler:    _UserServices_UserProductByName_Handler,
		},
		{
			MethodName: "UserProductByID",
			Handler:    _UserServices_UserProductByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
