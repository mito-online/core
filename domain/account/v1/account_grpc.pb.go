// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: account/v1/account.proto

package accountv1

import (
	context "context"
	v1 "github.com/mito/core/domain/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccountsServiceClient is the client API for AccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsServiceClient interface {
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	GetCurrentAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*v1.BoolResponse, error)
	SetAccountRole(ctx context.Context, in *SetAccountRoleRequest, opts ...grpc.CallOption) (*v1.BoolResponse, error)
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.BoolResponse, error)
}

type accountsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsServiceClient(cc grpc.ClientConnInterface) AccountsServiceClient {
	return &accountsServiceClient{cc}
}

func (c *accountsServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.v1.AccountsService/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/account.v1.AccountsService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetCurrentAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/account.v1.AccountsService/GetCurrentAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*v1.BoolResponse, error) {
	out := new(v1.BoolResponse)
	err := c.cc.Invoke(ctx, "/account.v1.AccountsService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) SetAccountRole(ctx context.Context, in *SetAccountRoleRequest, opts ...grpc.CallOption) (*v1.BoolResponse, error) {
	out := new(v1.BoolResponse)
	err := c.cc.Invoke(ctx, "/account.v1.AccountsService/SetAccountRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.BoolResponse, error) {
	out := new(v1.BoolResponse)
	err := c.cc.Invoke(ctx, "/account.v1.AccountsService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServiceServer is the server API for AccountsService service.
// All implementations must embed UnimplementedAccountsServiceServer
// for forward compatibility
type AccountsServiceServer interface {
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	GetCurrentAccount(context.Context, *emptypb.Empty) (*GetAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*v1.BoolResponse, error)
	SetAccountRole(context.Context, *SetAccountRoleRequest) (*v1.BoolResponse, error)
	Health(context.Context, *emptypb.Empty) (*v1.BoolResponse, error)
	mustEmbedUnimplementedAccountsServiceServer()
}

// UnimplementedAccountsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServiceServer struct {
}

func (UnimplementedAccountsServiceServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountsServiceServer) GetCurrentAccount(context.Context, *emptypb.Empty) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentAccount not implemented")
}
func (UnimplementedAccountsServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*v1.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountsServiceServer) SetAccountRole(context.Context, *SetAccountRoleRequest) (*v1.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccountRole not implemented")
}
func (UnimplementedAccountsServiceServer) Health(context.Context, *emptypb.Empty) (*v1.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedAccountsServiceServer) mustEmbedUnimplementedAccountsServiceServer() {}

// UnsafeAccountsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServiceServer will
// result in compilation errors.
type UnsafeAccountsServiceServer interface {
	mustEmbedUnimplementedAccountsServiceServer()
}

func RegisterAccountsServiceServer(s grpc.ServiceRegistrar, srv AccountsServiceServer) {
	s.RegisterService(&AccountsService_ServiceDesc, srv)
}

func _AccountsService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.AccountsService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.AccountsService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetCurrentAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetCurrentAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.AccountsService/GetCurrentAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetCurrentAccount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.AccountsService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_SetAccountRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAccountRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).SetAccountRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.AccountsService/SetAccountRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).SetAccountRole(ctx, req.(*SetAccountRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.AccountsService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountsService_ServiceDesc is the grpc.ServiceDesc for AccountsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.v1.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccounts",
			Handler:    _AccountsService_ListAccounts_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountsService_GetAccount_Handler,
		},
		{
			MethodName: "GetCurrentAccount",
			Handler:    _AccountsService_GetCurrentAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountsService_DeleteAccount_Handler,
		},
		{
			MethodName: "SetAccountRole",
			Handler:    _AccountsService_SetAccountRole_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _AccountsService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/v1/account.proto",
}
