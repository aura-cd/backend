// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/k8s_controller_service_v1.proto

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

// K8SCustomControllerClient is the client API for K8SCustomController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type K8SCustomControllerClient interface {
	// Namespace Control
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceReply, error)
	ListNamespaces(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListNamespacesReply, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deployment Control
	ApplyDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*CreateDeploymentReply, error)
	DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Resource Control
	GetAlls(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllsReply, error)
	GetOrgRepos(ctx context.Context, in *GetOrgRepoRequest, opts ...grpc.CallOption) (*GetOrgReposReply, error)
	CreatePreview(ctx context.Context, in *CreatePreviewRequest, opts ...grpc.CallOption) (*CreatePreviewReply, error)
	UpdatePreview(ctx context.Context, in *CreatePreviewRequest, opts ...grpc.CallOption) (*CreatePreviewReply, error)
	DeletePreview(ctx context.Context, in *DeletePreviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildImage(ctx context.Context, in *BuildImageRequest, opts ...grpc.CallOption) (*BuildImageReply, error)
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceReply, error)
}

type k8SCustomControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SCustomControllerClient(cc grpc.ClientConnInterface) K8SCustomControllerClient {
	return &k8SCustomControllerClient{cc}
}

func (c *k8SCustomControllerClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceReply, error) {
	out := new(CreateNamespaceReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) ListNamespaces(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListNamespacesReply, error) {
	out := new(ListNamespacesReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) ApplyDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*CreateDeploymentReply, error) {
	out := new(CreateDeploymentReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/ApplyDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) GetAlls(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllsReply, error) {
	out := new(GetAllsReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/GetAlls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) GetOrgRepos(ctx context.Context, in *GetOrgRepoRequest, opts ...grpc.CallOption) (*GetOrgReposReply, error) {
	out := new(GetOrgReposReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/GetOrgRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) CreatePreview(ctx context.Context, in *CreatePreviewRequest, opts ...grpc.CallOption) (*CreatePreviewReply, error) {
	out := new(CreatePreviewReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/CreatePreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) UpdatePreview(ctx context.Context, in *CreatePreviewRequest, opts ...grpc.CallOption) (*CreatePreviewReply, error) {
	out := new(CreatePreviewReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/UpdatePreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) DeletePreview(ctx context.Context, in *DeletePreviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/DeletePreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) BuildImage(ctx context.Context, in *BuildImageRequest, opts ...grpc.CallOption) (*BuildImageReply, error) {
	out := new(BuildImageReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/BuildImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SCustomControllerClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceReply, error) {
	out := new(GetResourceReply)
	err := c.cc.Invoke(ctx, "/gantrycd.k8s_controller.v1.K8sCustomController/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SCustomControllerServer is the server API for K8SCustomController service.
// All implementations must embed UnimplementedK8SCustomControllerServer
// for forward compatibility
type K8SCustomControllerServer interface {
	// Namespace Control
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceReply, error)
	ListNamespaces(context.Context, *emptypb.Empty) (*ListNamespacesReply, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error)
	// Deployment Control
	ApplyDeployment(context.Context, *CreateDeploymentRequest) (*CreateDeploymentReply, error)
	DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*emptypb.Empty, error)
	// Resource Control
	GetAlls(context.Context, *emptypb.Empty) (*GetAllsReply, error)
	GetOrgRepos(context.Context, *GetOrgRepoRequest) (*GetOrgReposReply, error)
	CreatePreview(context.Context, *CreatePreviewRequest) (*CreatePreviewReply, error)
	UpdatePreview(context.Context, *CreatePreviewRequest) (*CreatePreviewReply, error)
	DeletePreview(context.Context, *DeletePreviewRequest) (*emptypb.Empty, error)
	BuildImage(context.Context, *BuildImageRequest) (*BuildImageReply, error)
	GetResource(context.Context, *GetResourceRequest) (*GetResourceReply, error)
	mustEmbedUnimplementedK8SCustomControllerServer()
}

// UnimplementedK8SCustomControllerServer must be embedded to have forward compatible implementations.
type UnimplementedK8SCustomControllerServer struct {
}

func (UnimplementedK8SCustomControllerServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedK8SCustomControllerServer) ListNamespaces(context.Context, *emptypb.Empty) (*ListNamespacesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedK8SCustomControllerServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedK8SCustomControllerServer) ApplyDeployment(context.Context, *CreateDeploymentRequest) (*CreateDeploymentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyDeployment not implemented")
}
func (UnimplementedK8SCustomControllerServer) DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeployment not implemented")
}
func (UnimplementedK8SCustomControllerServer) GetAlls(context.Context, *emptypb.Empty) (*GetAllsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlls not implemented")
}
func (UnimplementedK8SCustomControllerServer) GetOrgRepos(context.Context, *GetOrgRepoRequest) (*GetOrgReposReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgRepos not implemented")
}
func (UnimplementedK8SCustomControllerServer) CreatePreview(context.Context, *CreatePreviewRequest) (*CreatePreviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePreview not implemented")
}
func (UnimplementedK8SCustomControllerServer) UpdatePreview(context.Context, *CreatePreviewRequest) (*CreatePreviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreview not implemented")
}
func (UnimplementedK8SCustomControllerServer) DeletePreview(context.Context, *DeletePreviewRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePreview not implemented")
}
func (UnimplementedK8SCustomControllerServer) BuildImage(context.Context, *BuildImageRequest) (*BuildImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildImage not implemented")
}
func (UnimplementedK8SCustomControllerServer) GetResource(context.Context, *GetResourceRequest) (*GetResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedK8SCustomControllerServer) mustEmbedUnimplementedK8SCustomControllerServer() {}

// UnsafeK8SCustomControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to K8SCustomControllerServer will
// result in compilation errors.
type UnsafeK8SCustomControllerServer interface {
	mustEmbedUnimplementedK8SCustomControllerServer()
}

func RegisterK8SCustomControllerServer(s grpc.ServiceRegistrar, srv K8SCustomControllerServer) {
	s.RegisterService(&K8SCustomController_ServiceDesc, srv)
}

func _K8SCustomController_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).ListNamespaces(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_ApplyDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).ApplyDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/ApplyDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).ApplyDeployment(ctx, req.(*CreateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_DeleteDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).DeleteDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/DeleteDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).DeleteDeployment(ctx, req.(*DeleteDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_GetAlls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).GetAlls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/GetAlls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).GetAlls(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_GetOrgRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).GetOrgRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/GetOrgRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).GetOrgRepos(ctx, req.(*GetOrgRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_CreatePreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePreviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).CreatePreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/CreatePreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).CreatePreview(ctx, req.(*CreatePreviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_UpdatePreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePreviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).UpdatePreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/UpdatePreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).UpdatePreview(ctx, req.(*CreatePreviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_DeletePreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePreviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).DeletePreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/DeletePreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).DeletePreview(ctx, req.(*DeletePreviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_BuildImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).BuildImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/BuildImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).BuildImage(ctx, req.(*BuildImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SCustomController_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCustomControllerServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gantrycd.k8s_controller.v1.K8sCustomController/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCustomControllerServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// K8SCustomController_ServiceDesc is the grpc.ServiceDesc for K8SCustomController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var K8SCustomController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gantrycd.k8s_controller.v1.K8sCustomController",
	HandlerType: (*K8SCustomControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _K8SCustomController_CreateNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _K8SCustomController_ListNamespaces_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _K8SCustomController_DeleteNamespace_Handler,
		},
		{
			MethodName: "ApplyDeployment",
			Handler:    _K8SCustomController_ApplyDeployment_Handler,
		},
		{
			MethodName: "DeleteDeployment",
			Handler:    _K8SCustomController_DeleteDeployment_Handler,
		},
		{
			MethodName: "GetAlls",
			Handler:    _K8SCustomController_GetAlls_Handler,
		},
		{
			MethodName: "GetOrgRepos",
			Handler:    _K8SCustomController_GetOrgRepos_Handler,
		},
		{
			MethodName: "CreatePreview",
			Handler:    _K8SCustomController_CreatePreview_Handler,
		},
		{
			MethodName: "UpdatePreview",
			Handler:    _K8SCustomController_UpdatePreview_Handler,
		},
		{
			MethodName: "DeletePreview",
			Handler:    _K8SCustomController_DeletePreview_Handler,
		},
		{
			MethodName: "BuildImage",
			Handler:    _K8SCustomController_BuildImage_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _K8SCustomController_GetResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/k8s_controller_service_v1.proto",
}
