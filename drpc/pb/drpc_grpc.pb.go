// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// DrpcServiceClient is the client API for DrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DrpcServiceClient interface {
	DrpcFunc(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type drpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDrpcServiceClient(cc grpc.ClientConnInterface) DrpcServiceClient {
	return &drpcServiceClient{cc}
}

func (c *drpcServiceClient) DrpcFunc(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/drpc.DrpcService/DrpcFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DrpcServiceServer is the server API for DrpcService service.
// All implementations must embed UnimplementedDrpcServiceServer
// for forward compatibility
type DrpcServiceServer interface {
	DrpcFunc(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedDrpcServiceServer()
}

// UnimplementedDrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDrpcServiceServer struct {
}

func (UnimplementedDrpcServiceServer) DrpcFunc(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrpcFunc not implemented")
}
func (UnimplementedDrpcServiceServer) mustEmbedUnimplementedDrpcServiceServer() {}

// UnsafeDrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DrpcServiceServer will
// result in compilation errors.
type UnsafeDrpcServiceServer interface {
	mustEmbedUnimplementedDrpcServiceServer()
}

func RegisterDrpcServiceServer(s grpc.ServiceRegistrar, srv DrpcServiceServer) {
	s.RegisterService(&DrpcService_ServiceDesc, srv)
}

func _DrpcService_DrpcFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DrpcServiceServer).DrpcFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drpc.DrpcService/DrpcFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DrpcServiceServer).DrpcFunc(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// DrpcService_ServiceDesc is the grpc.ServiceDesc for DrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "drpc.DrpcService",
	HandlerType: (*DrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DrpcFunc",
			Handler:    _DrpcService_DrpcFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drpc.proto",
}