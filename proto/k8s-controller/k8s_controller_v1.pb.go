// ライセンスはいつか書いておく

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.22.2
// source: proto/k8s-controller/k8s_controller_v1.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateNamespaceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNamespaceReply) Reset() {
	*x = CreateNamespaceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceReply) ProtoMessage() {}

func (x *CreateNamespaceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceReply.ProtoReflect.Descriptor instead.
func (*CreateNamespaceReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNamespaceReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListNamespacesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListNamespacesReply) Reset() {
	*x = ListNamespacesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespacesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespacesReply) ProtoMessage() {}

func (x *ListNamespacesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespacesReply.ProtoReflect.Descriptor instead.
func (*ListNamespacesReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ListNamespacesReply) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace     string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Repository    string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	PrNumber      string `protobuf:"bytes,3,opt,name=pr_number,json=prNumber,proto3" json:"pr_number,omitempty"`
	Image         string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	PodName       string `protobuf:"bytes,5,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Replicas      string `protobuf:"bytes,6,opt,name=replicas,proto3" json:"replicas,omitempty"`
	ContainerPort string `protobuf:"bytes,7,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
	CreatedBy     string `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *CreateDeploymentRequest) Reset() {
	*x = CreateDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeploymentRequest) ProtoMessage() {}

func (x *CreateDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeploymentRequest.ProtoReflect.Descriptor instead.
func (*CreateDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDeploymentRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateDeploymentRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *CreateDeploymentRequest) GetPrNumber() string {
	if x != nil {
		return x.PrNumber
	}
	return ""
}

func (x *CreateDeploymentRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateDeploymentRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *CreateDeploymentRequest) GetReplicas() string {
	if x != nil {
		return x.Replicas
	}
	return ""
}

func (x *CreateDeploymentRequest) GetContainerPort() string {
	if x != nil {
		return x.ContainerPort
	}
	return ""
}

func (x *CreateDeploymentRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type CreateDeploymentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version   string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CreateDeploymentReply) Reset() {
	*x = CreateDeploymentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeploymentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeploymentReply) ProtoMessage() {}

func (x *CreateDeploymentReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeploymentReply.ProtoReflect.Descriptor instead.
func (*CreateDeploymentReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDeploymentReply) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateDeploymentReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDeploymentReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type DeleteDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	PrNumber   string `protobuf:"bytes,3,opt,name=pr_number,json=prNumber,proto3" json:"pr_number,omitempty"`
}

func (x *DeleteDeploymentRequest) Reset() {
	*x = DeleteDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeploymentRequest) ProtoMessage() {}

func (x *DeleteDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeploymentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDeploymentRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteDeploymentRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *DeleteDeploymentRequest) GetPrNumber() string {
	if x != nil {
		return x.PrNumber
	}
	return ""
}

type Repos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repos []string `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
}

func (x *Repos) Reset() {
	*x = Repos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repos) ProtoMessage() {}

func (x *Repos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repos.ProtoReflect.Descriptor instead.
func (*Repos) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{7}
}

func (x *Repos) GetRepos() []string {
	if x != nil {
		return x.Repos
	}
	return nil
}

type OrgsAnsRepos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orgs  string   `protobuf:"bytes,1,opt,name=orgs,proto3" json:"orgs,omitempty"`
	Repos []*Repos `protobuf:"bytes,2,rep,name=repos,proto3" json:"repos,omitempty"`
}

func (x *OrgsAnsRepos) Reset() {
	*x = OrgsAnsRepos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgsAnsRepos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgsAnsRepos) ProtoMessage() {}

func (x *OrgsAnsRepos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgsAnsRepos.ProtoReflect.Descriptor instead.
func (*OrgsAnsRepos) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{8}
}

func (x *OrgsAnsRepos) GetOrgs() string {
	if x != nil {
		return x.Orgs
	}
	return ""
}

func (x *OrgsAnsRepos) GetRepos() []*Repos {
	if x != nil {
		return x.Repos
	}
	return nil
}

type ListOrgsAnsReposReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgsAnsRepos []*OrgsAnsRepos `protobuf:"bytes,1,rep,name=orgs_ans_repos,json=orgsAnsRepos,proto3" json:"orgs_ans_repos,omitempty"`
}

func (x *ListOrgsAnsReposReply) Reset() {
	*x = ListOrgsAnsReposReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrgsAnsReposReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrgsAnsReposReply) ProtoMessage() {}

func (x *ListOrgsAnsReposReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrgsAnsReposReply.ProtoReflect.Descriptor instead.
func (*ListOrgsAnsReposReply) Descriptor() ([]byte, []int) {
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP(), []int{9}
}

func (x *ListOrgsAnsReposReply) GetOrgsAnsRepos() []*OrgsAnsRepos {
	if x != nil {
		return x.OrgsAnsRepos
	}
	return nil
}

var File_proto_k8s_controller_k8s_controller_v1_proto protoreflect.FileDescriptor

var file_proto_k8s_controller_k8s_controller_v1_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x2c,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x87, 0x02, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x63, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x1d, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x22, 0x5b, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x73, 0x41, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6f, 0x72, 0x67, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b,
	0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x22, 0x67, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x73, 0x41, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x73, 0x5f, 0x61,
	0x6e, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x73,
	0x41, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x73, 0x41, 0x6e,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x32, 0x83, 0x05, 0x0a, 0x13, 0x4b, 0x38, 0x73, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x77,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x32, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64,
	0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2f, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64,
	0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x79, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e,
	0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x33, 0x2e, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5d, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x73, 0x41, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31, 0x2e, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x63, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x73, 0x41,
	0x6e, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72,
	0x79, 0x63, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_k8s_controller_k8s_controller_v1_proto_rawDescOnce sync.Once
	file_proto_k8s_controller_k8s_controller_v1_proto_rawDescData = file_proto_k8s_controller_k8s_controller_v1_proto_rawDesc
)

func file_proto_k8s_controller_k8s_controller_v1_proto_rawDescGZIP() []byte {
	file_proto_k8s_controller_k8s_controller_v1_proto_rawDescOnce.Do(func() {
		file_proto_k8s_controller_k8s_controller_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_k8s_controller_k8s_controller_v1_proto_rawDescData)
	})
	return file_proto_k8s_controller_k8s_controller_v1_proto_rawDescData
}

var file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_k8s_controller_k8s_controller_v1_proto_goTypes = []interface{}{
	(*CreateNamespaceRequest)(nil),  // 0: gantrycd.k8s_controller.v1.CreateNamespaceRequest
	(*CreateNamespaceReply)(nil),    // 1: gantrycd.k8s_controller.v1.CreateNamespaceReply
	(*ListNamespacesReply)(nil),     // 2: gantrycd.k8s_controller.v1.ListNamespacesReply
	(*DeleteNamespaceRequest)(nil),  // 3: gantrycd.k8s_controller.v1.DeleteNamespaceRequest
	(*CreateDeploymentRequest)(nil), // 4: gantrycd.k8s_controller.v1.CreateDeploymentRequest
	(*CreateDeploymentReply)(nil),   // 5: gantrycd.k8s_controller.v1.CreateDeploymentReply
	(*DeleteDeploymentRequest)(nil), // 6: gantrycd.k8s_controller.v1.DeleteDeploymentRequest
	(*Repos)(nil),                   // 7: gantrycd.k8s_controller.v1.Repos
	(*OrgsAnsRepos)(nil),            // 8: gantrycd.k8s_controller.v1.OrgsAnsRepos
	(*ListOrgsAnsReposReply)(nil),   // 9: gantrycd.k8s_controller.v1.ListOrgsAnsReposReply
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_proto_k8s_controller_k8s_controller_v1_proto_depIdxs = []int32{
	7,  // 0: gantrycd.k8s_controller.v1.OrgsAnsRepos.repos:type_name -> gantrycd.k8s_controller.v1.Repos
	8,  // 1: gantrycd.k8s_controller.v1.ListOrgsAnsReposReply.orgs_ans_repos:type_name -> gantrycd.k8s_controller.v1.OrgsAnsRepos
	0,  // 2: gantrycd.k8s_controller.v1.K8sCustomController.CreateNamespace:input_type -> gantrycd.k8s_controller.v1.CreateNamespaceRequest
	10, // 3: gantrycd.k8s_controller.v1.K8sCustomController.ListNamespaces:input_type -> google.protobuf.Empty
	3,  // 4: gantrycd.k8s_controller.v1.K8sCustomController.DeleteNamespace:input_type -> gantrycd.k8s_controller.v1.DeleteNamespaceRequest
	4,  // 5: gantrycd.k8s_controller.v1.K8sCustomController.ApplyDeployment:input_type -> gantrycd.k8s_controller.v1.CreateDeploymentRequest
	6,  // 6: gantrycd.k8s_controller.v1.K8sCustomController.DeleteDeployment:input_type -> gantrycd.k8s_controller.v1.DeleteDeploymentRequest
	10, // 7: gantrycd.k8s_controller.v1.K8sCustomController.ListOrgsAnsRepos:input_type -> google.protobuf.Empty
	1,  // 8: gantrycd.k8s_controller.v1.K8sCustomController.CreateNamespace:output_type -> gantrycd.k8s_controller.v1.CreateNamespaceReply
	2,  // 9: gantrycd.k8s_controller.v1.K8sCustomController.ListNamespaces:output_type -> gantrycd.k8s_controller.v1.ListNamespacesReply
	10, // 10: gantrycd.k8s_controller.v1.K8sCustomController.DeleteNamespace:output_type -> google.protobuf.Empty
	5,  // 11: gantrycd.k8s_controller.v1.K8sCustomController.ApplyDeployment:output_type -> gantrycd.k8s_controller.v1.CreateDeploymentReply
	10, // 12: gantrycd.k8s_controller.v1.K8sCustomController.DeleteDeployment:output_type -> google.protobuf.Empty
	9,  // 13: gantrycd.k8s_controller.v1.K8sCustomController.ListOrgsAnsRepos:output_type -> gantrycd.k8s_controller.v1.ListOrgsAnsReposReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_k8s_controller_k8s_controller_v1_proto_init() }
func file_proto_k8s_controller_k8s_controller_v1_proto_init() {
	if File_proto_k8s_controller_k8s_controller_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespacesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeploymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeploymentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeploymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgsAnsRepos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrgsAnsReposReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_k8s_controller_k8s_controller_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_k8s_controller_k8s_controller_v1_proto_goTypes,
		DependencyIndexes: file_proto_k8s_controller_k8s_controller_v1_proto_depIdxs,
		MessageInfos:      file_proto_k8s_controller_k8s_controller_v1_proto_msgTypes,
	}.Build()
	File_proto_k8s_controller_k8s_controller_v1_proto = out.File
	file_proto_k8s_controller_k8s_controller_v1_proto_rawDesc = nil
	file_proto_k8s_controller_k8s_controller_v1_proto_goTypes = nil
	file_proto_k8s_controller_k8s_controller_v1_proto_depIdxs = nil
}
