// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/bff_v1.proto

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

// BffServiceClient is the client API for BffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BffServiceClient interface {
	GetOrg(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HomeResponse, error)
	GetRepositoryApps(ctx context.Context, in *GetRepositoryAppsRequest, opts ...grpc.CallOption) (*GetRepositoryAppsResponse, error)
}

type bffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBffServiceClient(cc grpc.ClientConnInterface) BffServiceClient {
	return &bffServiceClient{cc}
}

func (c *bffServiceClient) GetOrg(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HomeResponse, error) {
	out := new(HomeResponse)
	err := c.cc.Invoke(ctx, "/gantrycd.bff.v1.BffService/GetOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffServiceClient) GetRepositoryApps(ctx context.Context, in *GetRepositoryAppsRequest, opts ...grpc.CallOption) (*GetRepositoryAppsResponse, error) {
	out := new(GetRepositoryAppsResponse)
	err := c.cc.Invoke(ctx, "/gantrycd.bff.v1.BffService/GetRepositoryApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BffServiceServer is the server API for BffService service.
// All implementations must embed UnimplementedBffServiceServer
// for forward compatibility
type BffServiceServer interface {
	GetOrg(context.Context, *emptypb.Empty) (*HomeResponse, error)
	GetRepositoryApps(context.Context, *GetRepositoryAppsRequest) (*GetRepositoryAppsResponse, error)
	mustEmbedUnimplementedBffServiceServer()
}

// UnimplementedBffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBffServiceServer struct {
}

func (UnimplementedBffServiceServer) GetOrg(context.Context, *emptypb.Empty) (*HomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedBffServiceServer) GetRepositoryApps(context.Context, *GetRepositoryAppsRequest) (*GetRepositoryAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepositoryApps not implemented")
}
func (UnimplementedBffServiceServer) mustEmbedUnimplementedBffServiceServer() {}

// UnsafeBffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BffServiceServer will
// result in compilation errors.
type UnsafeBffServiceServer interface {
	mustEmbedUnimplementedBffServiceServer()
}

func RegisterBffServiceServer(s grpc.ServiceRegistrar, srv BffServiceServer) {
	s.RegisterService(&BffService_ServiceDesc, srv)
}

func _BffService_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServiceServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.bff.v1.BffService/GetOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServiceServer).GetOrg(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BffService_GetRepositoryApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServiceServer).GetRepositoryApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.bff.v1.BffService/GetRepositoryApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServiceServer).GetRepositoryApps(ctx, req.(*GetRepositoryAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BffService_ServiceDesc is the grpc.ServiceDesc for BffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gantrycd.bff.v1.BffService",
	HandlerType: (*BffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrg",
			Handler:    _BffService_GetOrg_Handler,
		},
		{
			MethodName: "GetRepositoryApps",
			Handler:    _BffService_GetRepositoryApps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bff_v1.proto",
}