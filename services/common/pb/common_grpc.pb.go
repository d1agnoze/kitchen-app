// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services/common/pb/common.proto

package pb

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

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonServiceClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Puck(ctx context.Context, in *MessageDto, opts ...grpc.CallOption) (*PuckResponse, error)
}

type commonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonServiceClient(cc grpc.ClientConnInterface) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/CommonService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) Puck(ctx context.Context, in *MessageDto, opts ...grpc.CallOption) (*PuckResponse, error) {
	out := new(PuckResponse)
	err := c.cc.Invoke(ctx, "/CommonService/Puck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
// All implementations must embed UnimplementedCommonServiceServer
// for forward compatibility
type CommonServiceServer interface {
	Ping(context.Context, *empty.Empty) (*PingResponse, error)
	Puck(context.Context, *MessageDto) (*PuckResponse, error)
	mustEmbedUnimplementedCommonServiceServer()
}

// UnimplementedCommonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (UnimplementedCommonServiceServer) Ping(context.Context, *empty.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCommonServiceServer) Puck(context.Context, *MessageDto) (*PuckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Puck not implemented")
}
func (UnimplementedCommonServiceServer) mustEmbedUnimplementedCommonServiceServer() {}

// UnsafeCommonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServiceServer will
// result in compilation errors.
type UnsafeCommonServiceServer interface {
	mustEmbedUnimplementedCommonServiceServer()
}

func RegisterCommonServiceServer(s grpc.ServiceRegistrar, srv CommonServiceServer) {
	s.RegisterService(&CommonService_ServiceDesc, srv)
}

func _CommonService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_Puck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Puck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Puck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Puck(ctx, req.(*MessageDto))
	}
	return interceptor(ctx, in, info, handler)
}

// CommonService_ServiceDesc is the grpc.ServiceDesc for CommonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CommonService_Ping_Handler,
		},
		{
			MethodName: "Puck",
			Handler:    _CommonService_Puck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/common/pb/common.proto",
}