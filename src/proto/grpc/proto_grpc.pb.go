// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// PortDomainServiceClient is the client API for PortDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainServiceClient interface {
	Store(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Port, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Port, error)
}

type portDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainServiceClient(cc grpc.ClientConnInterface) PortDomainServiceClient {
	return &portDomainServiceClient{cc}
}

func (c *portDomainServiceClient) Store(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/grpc.PortDomainService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/grpc.PortDomainService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServiceServer is the server API for PortDomainService service.
// All implementations must embed UnimplementedPortDomainServiceServer
// for forward compatibility
type PortDomainServiceServer interface {
	Store(context.Context, *CreateRequest) (*Port, error)
	Fetch(context.Context, *FetchRequest) (*Port, error)
	mustEmbedUnimplementedPortDomainServiceServer()
}

// UnimplementedPortDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServiceServer struct {
}

func (UnimplementedPortDomainServiceServer) Store(context.Context, *CreateRequest) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedPortDomainServiceServer) Fetch(context.Context, *FetchRequest) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedPortDomainServiceServer) mustEmbedUnimplementedPortDomainServiceServer() {}

// UnsafePortDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServiceServer will
// result in compilation errors.
type UnsafePortDomainServiceServer interface {
	mustEmbedUnimplementedPortDomainServiceServer()
}

func RegisterPortDomainServiceServer(s grpc.ServiceRegistrar, srv PortDomainServiceServer) {
	s.RegisterService(&PortDomainService_ServiceDesc, srv)
}

func _PortDomainService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PortDomainService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).Store(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomainService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PortDomainService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortDomainService_ServiceDesc is the grpc.ServiceDesc for PortDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PortDomainService",
	HandlerType: (*PortDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _PortDomainService_Store_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _PortDomainService_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}