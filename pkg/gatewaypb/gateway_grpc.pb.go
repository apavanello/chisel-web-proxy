// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: gateway.proto

package gatewaypb

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

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/gateway.ProxyService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/gateway.ProxyService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/gateway.ProxyService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServiceServer is the server API for ProxyService service.
// All implementations must embed UnimplementedProxyServiceServer
// for forward compatibility
type ProxyServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	mustEmbedUnimplementedProxyServiceServer()
}

// UnimplementedProxyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServiceServer struct {
}

func (UnimplementedProxyServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedProxyServiceServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedProxyServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedProxyServiceServer) mustEmbedUnimplementedProxyServiceServer() {}

// UnsafeProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServiceServer will
// result in compilation errors.
type UnsafeProxyServiceServer interface {
	mustEmbedUnimplementedProxyServiceServer()
}

func RegisterProxyServiceServer(s grpc.ServiceRegistrar, srv ProxyServiceServer) {
	s.RegisterService(&ProxyService_ServiceDesc, srv)
}

func _ProxyService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.ProxyService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.ProxyService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.ProxyService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProxyService_ServiceDesc is the grpc.ServiceDesc for ProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _ProxyService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _ProxyService_Disconnect_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ProxyService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}

// PreLoadServiceClient is the client API for PreLoadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreLoadServiceClient interface {
	PreLoad(ctx context.Context, in *PreLoadRequest, opts ...grpc.CallOption) (*PreLoadResponse, error)
	GetHosts(ctx context.Context, in *HostsRequest, opts ...grpc.CallOption) (*HostsResponse, error)
}

type preLoadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPreLoadServiceClient(cc grpc.ClientConnInterface) PreLoadServiceClient {
	return &preLoadServiceClient{cc}
}

func (c *preLoadServiceClient) PreLoad(ctx context.Context, in *PreLoadRequest, opts ...grpc.CallOption) (*PreLoadResponse, error) {
	out := new(PreLoadResponse)
	err := c.cc.Invoke(ctx, "/gateway.PreLoadService/PreLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *preLoadServiceClient) GetHosts(ctx context.Context, in *HostsRequest, opts ...grpc.CallOption) (*HostsResponse, error) {
	out := new(HostsResponse)
	err := c.cc.Invoke(ctx, "/gateway.PreLoadService/GetHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreLoadServiceServer is the server API for PreLoadService service.
// All implementations must embed UnimplementedPreLoadServiceServer
// for forward compatibility
type PreLoadServiceServer interface {
	PreLoad(context.Context, *PreLoadRequest) (*PreLoadResponse, error)
	GetHosts(context.Context, *HostsRequest) (*HostsResponse, error)
	mustEmbedUnimplementedPreLoadServiceServer()
}

// UnimplementedPreLoadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPreLoadServiceServer struct {
}

func (UnimplementedPreLoadServiceServer) PreLoad(context.Context, *PreLoadRequest) (*PreLoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreLoad not implemented")
}
func (UnimplementedPreLoadServiceServer) GetHosts(context.Context, *HostsRequest) (*HostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHosts not implemented")
}
func (UnimplementedPreLoadServiceServer) mustEmbedUnimplementedPreLoadServiceServer() {}

// UnsafePreLoadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreLoadServiceServer will
// result in compilation errors.
type UnsafePreLoadServiceServer interface {
	mustEmbedUnimplementedPreLoadServiceServer()
}

func RegisterPreLoadServiceServer(s grpc.ServiceRegistrar, srv PreLoadServiceServer) {
	s.RegisterService(&PreLoadService_ServiceDesc, srv)
}

func _PreLoadService_PreLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreLoadServiceServer).PreLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.PreLoadService/PreLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreLoadServiceServer).PreLoad(ctx, req.(*PreLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PreLoadService_GetHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreLoadServiceServer).GetHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.PreLoadService/GetHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreLoadServiceServer).GetHosts(ctx, req.(*HostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PreLoadService_ServiceDesc is the grpc.ServiceDesc for PreLoadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PreLoadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.PreLoadService",
	HandlerType: (*PreLoadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreLoad",
			Handler:    _PreLoadService_PreLoad_Handler,
		},
		{
			MethodName: "GetHosts",
			Handler:    _PreLoadService_GetHosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
