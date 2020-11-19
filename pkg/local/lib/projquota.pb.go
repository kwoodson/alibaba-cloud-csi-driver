// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: projquota.proto

package lib

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProjQuotaSubpathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PvName    string `protobuf:"bytes,1,opt,name=pvName,proto3" json:"pvName,omitempty"`
	QuotaSize string `protobuf:"bytes,2,opt,name=quotaSize,proto3" json:"quotaSize,omitempty"`
}

func (x *CreateProjQuotaSubpathRequest) Reset() {
	*x = CreateProjQuotaSubpathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projquota_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjQuotaSubpathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjQuotaSubpathRequest) ProtoMessage() {}

func (x *CreateProjQuotaSubpathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projquota_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjQuotaSubpathRequest.ProtoReflect.Descriptor instead.
func (*CreateProjQuotaSubpathRequest) Descriptor() ([]byte, []int) {
	return file_projquota_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjQuotaSubpathRequest) GetPvName() string {
	if x != nil {
		return x.PvName
	}
	return ""
}

func (x *CreateProjQuotaSubpathRequest) GetQuotaSize() string {
	if x != nil {
		return x.QuotaSize
	}
	return ""
}

type CreateProjQuotaSubpathReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjQuotaSubpath string `protobuf:"bytes,1,opt,name=proj_quota_subpath,json=projQuotaSubpath,proto3" json:"proj_quota_subpath,omitempty"`
	CommandOutput    string `protobuf:"bytes,2,opt,name=command_output,json=commandOutput,proto3" json:"command_output,omitempty"`
	ProjectId        string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateProjQuotaSubpathReply) Reset() {
	*x = CreateProjQuotaSubpathReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projquota_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjQuotaSubpathReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjQuotaSubpathReply) ProtoMessage() {}

func (x *CreateProjQuotaSubpathReply) ProtoReflect() protoreflect.Message {
	mi := &file_projquota_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjQuotaSubpathReply.ProtoReflect.Descriptor instead.
func (*CreateProjQuotaSubpathReply) Descriptor() ([]byte, []int) {
	return file_projquota_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjQuotaSubpathReply) GetProjQuotaSubpath() string {
	if x != nil {
		return x.ProjQuotaSubpath
	}
	return ""
}

func (x *CreateProjQuotaSubpathReply) GetCommandOutput() string {
	if x != nil {
		return x.CommandOutput
	}
	return ""
}

func (x *CreateProjQuotaSubpathReply) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type RemoveProjQuotaSubpathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuotaSubpath string `protobuf:"bytes,1,opt,name=quota_subpath,json=quotaSubpath,proto3" json:"quota_subpath,omitempty"`
	ProjectId    string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *RemoveProjQuotaSubpathRequest) Reset() {
	*x = RemoveProjQuotaSubpathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projquota_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProjQuotaSubpathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProjQuotaSubpathRequest) ProtoMessage() {}

func (x *RemoveProjQuotaSubpathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projquota_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProjQuotaSubpathRequest.ProtoReflect.Descriptor instead.
func (*RemoveProjQuotaSubpathRequest) Descriptor() ([]byte, []int) {
	return file_projquota_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveProjQuotaSubpathRequest) GetQuotaSubpath() string {
	if x != nil {
		return x.QuotaSubpath
	}
	return ""
}

func (x *RemoveProjQuotaSubpathRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type RemoveProjQuotaSubpathReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandOutput string `protobuf:"bytes,1,opt,name=command_output,json=commandOutput,proto3" json:"command_output,omitempty"`
}

func (x *RemoveProjQuotaSubpathReply) Reset() {
	*x = RemoveProjQuotaSubpathReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projquota_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProjQuotaSubpathReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProjQuotaSubpathReply) ProtoMessage() {}

func (x *RemoveProjQuotaSubpathReply) ProtoReflect() protoreflect.Message {
	mi := &file_projquota_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProjQuotaSubpathReply.ProtoReflect.Descriptor instead.
func (*RemoveProjQuotaSubpathReply) Descriptor() ([]byte, []int) {
	return file_projquota_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveProjQuotaSubpathReply) GetCommandOutput() string {
	if x != nil {
		return x.CommandOutput
	}
	return ""
}

type SetSubpathProjQuotaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjQuotaSubpath string `protobuf:"bytes,1,opt,name=proj_quota_subpath,json=projQuotaSubpath,proto3" json:"proj_quota_subpath,omitempty"`
	BlockSoftlimit   string `protobuf:"bytes,2,opt,name=block_softlimit,json=blockSoftlimit,proto3" json:"block_softlimit,omitempty"`
	BlockHardlimit   string `protobuf:"bytes,3,opt,name=block_hardlimit,json=blockHardlimit,proto3" json:"block_hardlimit,omitempty"`
	InodeSoftlimit   string `protobuf:"bytes,4,opt,name=inode_softlimit,json=inodeSoftlimit,proto3" json:"inode_softlimit,omitempty"`
	InodeHardlimit   string `protobuf:"bytes,5,opt,name=inode_hardlimit,json=inodeHardlimit,proto3" json:"inode_hardlimit,omitempty"`
}

func (x *SetSubpathProjQuotaRequest) Reset() {
	*x = SetSubpathProjQuotaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projquota_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSubpathProjQuotaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSubpathProjQuotaRequest) ProtoMessage() {}

func (x *SetSubpathProjQuotaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projquota_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSubpathProjQuotaRequest.ProtoReflect.Descriptor instead.
func (*SetSubpathProjQuotaRequest) Descriptor() ([]byte, []int) {
	return file_projquota_proto_rawDescGZIP(), []int{4}
}

func (x *SetSubpathProjQuotaRequest) GetProjQuotaSubpath() string {
	if x != nil {
		return x.ProjQuotaSubpath
	}
	return ""
}

func (x *SetSubpathProjQuotaRequest) GetBlockSoftlimit() string {
	if x != nil {
		return x.BlockSoftlimit
	}
	return ""
}

func (x *SetSubpathProjQuotaRequest) GetBlockHardlimit() string {
	if x != nil {
		return x.BlockHardlimit
	}
	return ""
}

func (x *SetSubpathProjQuotaRequest) GetInodeSoftlimit() string {
	if x != nil {
		return x.InodeSoftlimit
	}
	return ""
}

func (x *SetSubpathProjQuotaRequest) GetInodeHardlimit() string {
	if x != nil {
		return x.InodeHardlimit
	}
	return ""
}

type SetSubpathProjQuotaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandOutput string `protobuf:"bytes,1,opt,name=command_output,json=commandOutput,proto3" json:"command_output,omitempty"`
}

func (x *SetSubpathProjQuotaReply) Reset() {
	*x = SetSubpathProjQuotaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projquota_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSubpathProjQuotaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSubpathProjQuotaReply) ProtoMessage() {}

func (x *SetSubpathProjQuotaReply) ProtoReflect() protoreflect.Message {
	mi := &file_projquota_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSubpathProjQuotaReply.ProtoReflect.Descriptor instead.
func (*SetSubpathProjQuotaReply) Descriptor() ([]byte, []int) {
	return file_projquota_proto_rawDescGZIP(), []int{5}
}

func (x *SetSubpathProjQuotaReply) GetCommandOutput() string {
	if x != nil {
		return x.CommandOutput
	}
	return ""
}

var File_projquota_proto protoreflect.FileDescriptor

var file_projquota_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x76, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x76, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x91, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x75,
	0x62, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f,
	0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x1d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x75,
	0x62, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xee,
	0x01, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x70, 0x72, 0x6f, 0x6a, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x75, 0x62, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x6f, 0x66, 0x74, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61,
	0x72, 0x64, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x72, 0x64, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x6f, 0x66,
	0x74, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x68, 0x61, 0x72, 0x64, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x72, 0x64, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x41, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x32, 0xb4, 0x02, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x12, 0x64, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x70,
	0x61, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x12, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x75, 0x62, 0x70, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_projquota_proto_rawDescOnce sync.Once
	file_projquota_proto_rawDescData = file_projquota_proto_rawDesc
)

func file_projquota_proto_rawDescGZIP() []byte {
	file_projquota_proto_rawDescOnce.Do(func() {
		file_projquota_proto_rawDescData = protoimpl.X.CompressGZIP(file_projquota_proto_rawDescData)
	})
	return file_projquota_proto_rawDescData
}

var file_projquota_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_projquota_proto_goTypes = []interface{}{
	(*CreateProjQuotaSubpathRequest)(nil), // 0: proto.CreateProjQuotaSubpathRequest
	(*CreateProjQuotaSubpathReply)(nil),   // 1: proto.CreateProjQuotaSubpathReply
	(*RemoveProjQuotaSubpathRequest)(nil), // 2: proto.RemoveProjQuotaSubpathRequest
	(*RemoveProjQuotaSubpathReply)(nil),   // 3: proto.RemoveProjQuotaSubpathReply
	(*SetSubpathProjQuotaRequest)(nil),    // 4: proto.SetSubpathProjQuotaRequest
	(*SetSubpathProjQuotaReply)(nil),      // 5: proto.SetSubpathProjQuotaReply
}
var file_projquota_proto_depIdxs = []int32{
	0, // 0: proto.ProjQuota.CreateProjQuotaSubpath:input_type -> proto.CreateProjQuotaSubpathRequest
	4, // 1: proto.ProjQuota.SetSubpathProjQuota:input_type -> proto.SetSubpathProjQuotaRequest
	2, // 2: proto.ProjQuota.RemoveProjQuotaSubpath:input_type -> proto.RemoveProjQuotaSubpathRequest
	1, // 3: proto.ProjQuota.CreateProjQuotaSubpath:output_type -> proto.CreateProjQuotaSubpathReply
	5, // 4: proto.ProjQuota.SetSubpathProjQuota:output_type -> proto.SetSubpathProjQuotaReply
	3, // 5: proto.ProjQuota.RemoveProjQuotaSubpath:output_type -> proto.RemoveProjQuotaSubpathReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_projquota_proto_init() }
func file_projquota_proto_init() {
	if File_projquota_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_projquota_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjQuotaSubpathRequest); i {
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
		file_projquota_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjQuotaSubpathReply); i {
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
		file_projquota_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProjQuotaSubpathRequest); i {
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
		file_projquota_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProjQuotaSubpathReply); i {
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
		file_projquota_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSubpathProjQuotaRequest); i {
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
		file_projquota_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSubpathProjQuotaReply); i {
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
			RawDescriptor: file_projquota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_projquota_proto_goTypes,
		DependencyIndexes: file_projquota_proto_depIdxs,
		MessageInfos:      file_projquota_proto_msgTypes,
	}.Build()
	File_projquota_proto = out.File
	file_projquota_proto_rawDesc = nil
	file_projquota_proto_goTypes = nil
	file_projquota_proto_depIdxs = nil
}
