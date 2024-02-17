// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/metric/metric_v1.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResourceWatcherClient is the client API for ResourceWatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceWatcherClient interface {
	GetResource(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetResourceReply, error)
}

type resourceWatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceWatcherClient(cc grpc.ClientConnInterface) ResourceWatcherClient {
	return &resourceWatcherClient{cc}
}

func (c *resourceWatcherClient) GetResource(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetResourceReply, error) {
	out := new(GetResourceReply)
	err := c.cc.Invoke(ctx, "/gantrycd.resource_v1.ResourceWatcher/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceWatcherServer is the server API for ResourceWatcher service.
// All implementations must embed UnimplementedResourceWatcherServer
// for forward compatibility
type ResourceWatcherServer interface {
	GetResource(context.Context, *emptypb.Empty) (*GetResourceReply, error)
	mustEmbedUnimplementedResourceWatcherServer()
}

// UnimplementedResourceWatcherServer must be embedded to have forward compatible implementations.
type UnimplementedResourceWatcherServer struct {
}

func (UnimplementedResourceWatcherServer) GetResource(context.Context, *emptypb.Empty) (*GetResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedResourceWatcherServer) mustEmbedUnimplementedResourceWatcherServer() {}

// UnsafeResourceWatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceWatcherServer will
// result in compilation errors.
type UnsafeResourceWatcherServer interface {
	mustEmbedUnimplementedResourceWatcherServer()
}

func RegisterResourceWatcherServer(s grpc.ServiceRegistrar, srv ResourceWatcherServer) {
	s.RegisterService(&ResourceWatcher_ServiceDesc, srv)
}

func _ResourceWatcher_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceWatcherServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.resource_v1.ResourceWatcher/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceWatcherServer).GetResource(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceWatcher_ServiceDesc is the grpc.ServiceDesc for ResourceWatcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceWatcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gantrycd.resource_v1.ResourceWatcher",
	HandlerType: (*ResourceWatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _ResourceWatcher_GetResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metric/metric_v1.proto",
}
