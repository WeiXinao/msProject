// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: api/proto/file/v1/file.proto

package v1

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

type TaskFileRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	PathName         string                 `protobuf:"bytes,1,opt,name=pathName,proto3" json:"pathName,omitempty"`
	FileName         string                 `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Extension        string                 `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	Size             int64                  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	ProjectCode      string                 `protobuf:"bytes,5,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	TaskCode         string                 `protobuf:"bytes,6,opt,name=taskCode,proto3" json:"taskCode,omitempty"`
	OrganizationCode string                 `protobuf:"bytes,7,opt,name=organizationCode,proto3" json:"organizationCode,omitempty"`
	FileUrl          string                 `protobuf:"bytes,8,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
	FileType         string                 `protobuf:"bytes,9,opt,name=fileType,proto3" json:"fileType,omitempty"`
	MemberId         int64                  `protobuf:"varint,10,opt,name=memberId,proto3" json:"memberId,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TaskFileRequest) Reset() {
	*x = TaskFileRequest{}
	mi := &file_api_proto_file_v1_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFileRequest) ProtoMessage() {}

func (x *TaskFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_v1_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFileRequest.ProtoReflect.Descriptor instead.
func (*TaskFileRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *TaskFileRequest) GetPathName() string {
	if x != nil {
		return x.PathName
	}
	return ""
}

func (x *TaskFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *TaskFileRequest) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *TaskFileRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TaskFileRequest) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *TaskFileRequest) GetTaskCode() string {
	if x != nil {
		return x.TaskCode
	}
	return ""
}

func (x *TaskFileRequest) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *TaskFileRequest) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *TaskFileRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *TaskFileRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type TaskFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskFileResponse) Reset() {
	*x = TaskFileResponse{}
	mi := &file_api_proto_file_v1_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFileResponse) ProtoMessage() {}

func (x *TaskFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_v1_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFileResponse.ProtoReflect.Descriptor instead.
func (*TaskFileResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_file_v1_file_proto_rawDescGZIP(), []int{1}
}

type TaskSourcesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskCode      string                 `protobuf:"bytes,1,opt,name=taskCode,proto3" json:"taskCode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskSourcesRequest) Reset() {
	*x = TaskSourcesRequest{}
	mi := &file_api_proto_file_v1_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskSourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSourcesRequest) ProtoMessage() {}

func (x *TaskSourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_v1_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSourcesRequest.ProtoReflect.Descriptor instead.
func (*TaskSourcesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_v1_file_proto_rawDescGZIP(), []int{2}
}

func (x *TaskSourcesRequest) GetTaskCode() string {
	if x != nil {
		return x.TaskCode
	}
	return ""
}

type TaskSourceMessage struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	SourceType       string                 `protobuf:"bytes,3,opt,name=sourceType,proto3" json:"sourceType,omitempty"`
	SourceCode       string                 `protobuf:"bytes,4,opt,name=sourceCode,proto3" json:"sourceCode,omitempty"`
	LinkType         string                 `protobuf:"bytes,5,opt,name=linkType,proto3" json:"linkType,omitempty"`
	LinkCode         string                 `protobuf:"bytes,6,opt,name=linkCode,proto3" json:"linkCode,omitempty"`
	OrganizationCode string                 `protobuf:"bytes,7,opt,name=OrganizationCode,proto3" json:"OrganizationCode,omitempty"`
	CreateBy         string                 `protobuf:"bytes,8,opt,name=createBy,proto3" json:"createBy,omitempty"`
	CreateTime       string                 `protobuf:"bytes,9,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Sort             int32                  `protobuf:"varint,10,opt,name=sort,proto3" json:"sort,omitempty"`
	Title            string                 `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	SourceDetail     *SourceDetail          `protobuf:"bytes,12,opt,name=sourceDetail,proto3" json:"sourceDetail,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TaskSourceMessage) Reset() {
	*x = TaskSourceMessage{}
	mi := &file_api_proto_file_v1_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskSourceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSourceMessage) ProtoMessage() {}

func (x *TaskSourceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_v1_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSourceMessage.ProtoReflect.Descriptor instead.
func (*TaskSourceMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_file_v1_file_proto_rawDescGZIP(), []int{3}
}

func (x *TaskSourceMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskSourceMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TaskSourceMessage) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *TaskSourceMessage) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

func (x *TaskSourceMessage) GetLinkType() string {
	if x != nil {
		return x.LinkType
	}
	return ""
}

func (x *TaskSourceMessage) GetLinkCode() string {
	if x != nil {
		return x.LinkCode
	}
	return ""
}

func (x *TaskSourceMessage) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *TaskSourceMessage) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *TaskSourceMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *TaskSourceMessage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *TaskSourceMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskSourceMessage) GetSourceDetail() *SourceDetail {
	if x != nil {
		return x.SourceDetail
	}
	return nil
}

type SourceDetail struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	PathName         string                 `protobuf:"bytes,3,opt,name=pathName,proto3" json:"pathName,omitempty"`
	Title            string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Extension        string                 `protobuf:"bytes,5,opt,name=Extension,proto3" json:"Extension,omitempty"`
	Size             int32                  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	ObjectType       string                 `protobuf:"bytes,7,opt,name=ObjectType,proto3" json:"ObjectType,omitempty"`
	OrganizationCode string                 `protobuf:"bytes,8,opt,name=OrganizationCode,proto3" json:"OrganizationCode,omitempty"`
	TaskCode         string                 `protobuf:"bytes,9,opt,name=TaskCode,proto3" json:"TaskCode,omitempty"`
	ProjectCode      string                 `protobuf:"bytes,10,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	CreateBy         string                 `protobuf:"bytes,11,opt,name=createBy,proto3" json:"createBy,omitempty"`
	CreateTime       string                 `protobuf:"bytes,12,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Downloads        int32                  `protobuf:"varint,13,opt,name=downloads,proto3" json:"downloads,omitempty"`
	Extra            string                 `protobuf:"bytes,14,opt,name=Extra,proto3" json:"Extra,omitempty"`
	Deleted          int32                  `protobuf:"varint,15,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	FileUrl          string                 `protobuf:"bytes,16,opt,name=FileUrl,proto3" json:"FileUrl,omitempty"`
	FileType         string                 `protobuf:"bytes,17,opt,name=FileType,proto3" json:"FileType,omitempty"`
	DeletedTime      string                 `protobuf:"bytes,18,opt,name=deletedTime,proto3" json:"deletedTime,omitempty"`
	ProjectName      string                 `protobuf:"bytes,19,opt,name=ProjectName,proto3" json:"ProjectName,omitempty"`
	FullName         string                 `protobuf:"bytes,20,opt,name=FullName,proto3" json:"FullName,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *SourceDetail) Reset() {
	*x = SourceDetail{}
	mi := &file_api_proto_file_v1_file_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceDetail) ProtoMessage() {}

func (x *SourceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_v1_file_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceDetail.ProtoReflect.Descriptor instead.
func (*SourceDetail) Descriptor() ([]byte, []int) {
	return file_api_proto_file_v1_file_proto_rawDescGZIP(), []int{4}
}

func (x *SourceDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SourceDetail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SourceDetail) GetPathName() string {
	if x != nil {
		return x.PathName
	}
	return ""
}

func (x *SourceDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SourceDetail) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *SourceDetail) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SourceDetail) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *SourceDetail) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *SourceDetail) GetTaskCode() string {
	if x != nil {
		return x.TaskCode
	}
	return ""
}

func (x *SourceDetail) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *SourceDetail) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *SourceDetail) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *SourceDetail) GetDownloads() int32 {
	if x != nil {
		return x.Downloads
	}
	return 0
}

func (x *SourceDetail) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *SourceDetail) GetDeleted() int32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *SourceDetail) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *SourceDetail) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *SourceDetail) GetDeletedTime() string {
	if x != nil {
		return x.DeletedTime
	}
	return ""
}

func (x *SourceDetail) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *SourceDetail) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type TaskSourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*TaskSourceMessage   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskSourceResponse) Reset() {
	*x = TaskSourceResponse{}
	mi := &file_api_proto_file_v1_file_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSourceResponse) ProtoMessage() {}

func (x *TaskSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_v1_file_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSourceResponse.ProtoReflect.Descriptor instead.
func (*TaskSourceResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_file_v1_file_proto_rawDescGZIP(), []int{5}
}

func (x *TaskSourceResponse) GetList() []*TaskSourceMessage {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_proto_file_v1_file_proto protoreflect.FileDescriptor

var file_api_proto_file_v1_file_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0xb7, 0x02, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x68, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x74, 0x68, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x54,
	0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x30, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x86, 0x03, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xc0, 0x04, 0x0a, 0x0c, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a,
	0x12, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xbc, 0x01,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5b, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x25,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_file_v1_file_proto_rawDescOnce sync.Once
	file_api_proto_file_v1_file_proto_rawDescData = file_api_proto_file_v1_file_proto_rawDesc
)

func file_api_proto_file_v1_file_proto_rawDescGZIP() []byte {
	file_api_proto_file_v1_file_proto_rawDescOnce.Do(func() {
		file_api_proto_file_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_file_v1_file_proto_rawDescData)
	})
	return file_api_proto_file_v1_file_proto_rawDescData
}

var file_api_proto_file_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_file_v1_file_proto_goTypes = []any{
	(*TaskFileRequest)(nil),    // 0: api.proto.file.v1.TaskFileRequest
	(*TaskFileResponse)(nil),   // 1: api.proto.file.v1.TaskFileResponse
	(*TaskSourcesRequest)(nil), // 2: api.proto.file.v1.TaskSourcesRequest
	(*TaskSourceMessage)(nil),  // 3: api.proto.file.v1.TaskSourceMessage
	(*SourceDetail)(nil),       // 4: api.proto.file.v1.SourceDetail
	(*TaskSourceResponse)(nil), // 5: api.proto.file.v1.TaskSourceResponse
}
var file_api_proto_file_v1_file_proto_depIdxs = []int32{
	4, // 0: api.proto.file.v1.TaskSourceMessage.sourceDetail:type_name -> api.proto.file.v1.SourceDetail
	3, // 1: api.proto.file.v1.TaskSourceResponse.list:type_name -> api.proto.file.v1.TaskSourceMessage
	0, // 2: api.proto.file.v1.File.SaveTaskFile:input_type -> api.proto.file.v1.TaskFileRequest
	2, // 3: api.proto.file.v1.File.TaskSources:input_type -> api.proto.file.v1.TaskSourcesRequest
	1, // 4: api.proto.file.v1.File.SaveTaskFile:output_type -> api.proto.file.v1.TaskFileResponse
	5, // 5: api.proto.file.v1.File.TaskSources:output_type -> api.proto.file.v1.TaskSourceResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_file_v1_file_proto_init() }
func file_api_proto_file_v1_file_proto_init() {
	if File_api_proto_file_v1_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_file_v1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_file_v1_file_proto_goTypes,
		DependencyIndexes: file_api_proto_file_v1_file_proto_depIdxs,
		MessageInfos:      file_api_proto_file_v1_file_proto_msgTypes,
	}.Build()
	File_api_proto_file_v1_file_proto = out.File
	file_api_proto_file_v1_file_proto_rawDesc = nil
	file_api_proto_file_v1_file_proto_goTypes = nil
	file_api_proto_file_v1_file_proto_depIdxs = nil
}
