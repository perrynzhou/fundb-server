//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: dbserver.proto

package pb

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

type CreateSchemaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateSchemaReq) Reset() {
	*x = CreateSchemaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaReq) ProtoMessage() {}

func (x *CreateSchemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchemaReq.ProtoReflect.Descriptor instead.
func (*CreateSchemaReq) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSchemaReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateSchemaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateSchemaResp) Reset() {
	*x = CreateSchemaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaResp) ProtoMessage() {}

func (x *CreateSchemaResp) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchemaResp.ProtoReflect.Descriptor instead.
func (*CreateSchemaResp) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSchemaResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSchemaResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateSchemaResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type QuerySchemaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QuerySchemaReq) Reset() {
	*x = QuerySchemaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySchemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySchemaReq) ProtoMessage() {}

func (x *QuerySchemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySchemaReq.ProtoReflect.Descriptor instead.
func (*QuerySchemaReq) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySchemaReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QuerySchemaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Meta []string `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty"`
}

func (x *QuerySchemaResp) Reset() {
	*x = QuerySchemaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySchemaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySchemaResp) ProtoMessage() {}

func (x *QuerySchemaResp) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySchemaResp.ProtoReflect.Descriptor instead.
func (*QuerySchemaResp) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{3}
}

func (x *QuerySchemaResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QuerySchemaResp) GetMeta() []string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type DropSchemaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DropSchemaReq) Reset() {
	*x = DropSchemaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropSchemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropSchemaReq) ProtoMessage() {}

func (x *DropSchemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropSchemaReq.ProtoReflect.Descriptor instead.
func (*DropSchemaReq) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{4}
}

func (x *DropSchemaReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DropSchemaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DropSchemaResp) Reset() {
	*x = DropSchemaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropSchemaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropSchemaResp) ProtoMessage() {}

func (x *DropSchemaResp) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropSchemaResp.ProtoReflect.Descriptor instead.
func (*DropSchemaResp) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{5}
}

func (x *DropSchemaResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DropSchemaResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DropSchemaResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PutKvReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value      string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutKvReq) Reset() {
	*x = PutKvReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutKvReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutKvReq) ProtoMessage() {}

func (x *PutKvReq) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutKvReq.ProtoReflect.Descriptor instead.
func (*PutKvReq) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{6}
}

func (x *PutKvReq) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *PutKvReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutKvReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutKvResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PutKvResp) Reset() {
	*x = PutKvResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutKvResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutKvResp) ProtoMessage() {}

func (x *PutKvResp) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutKvResp.ProtoReflect.Descriptor instead.
func (*PutKvResp) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{7}
}

func (x *PutKvResp) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *PutKvResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetKvReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKvReq) Reset() {
	*x = GetKvReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKvReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKvReq) ProtoMessage() {}

func (x *GetKvReq) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKvReq.ProtoReflect.Descriptor instead.
func (*GetKvReq) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{8}
}

func (x *GetKvReq) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *GetKvReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetKvResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value      string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetKvResp) Reset() {
	*x = GetKvResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKvResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKvResp) ProtoMessage() {}

func (x *GetKvResp) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKvResp.ProtoReflect.Descriptor instead.
func (*GetKvResp) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{9}
}

func (x *GetKvResp) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *GetKvResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetKvResp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DelKvReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelKvReq) Reset() {
	*x = DelKvReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelKvReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelKvReq) ProtoMessage() {}

func (x *DelKvReq) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelKvReq.ProtoReflect.Descriptor instead.
func (*DelKvReq) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{10}
}

func (x *DelKvReq) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *DelKvReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelKvResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value      string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DelKvResp) Reset() {
	*x = DelKvResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbserver_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelKvResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelKvResp) ProtoMessage() {}

func (x *DelKvResp) ProtoReflect() protoreflect.Message {
	mi := &file_dbserver_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelKvResp.ProtoReflect.Descriptor instead.
func (*DelKvResp) Descriptor() ([]byte, []int) {
	return file_dbserver_proto_rawDescGZIP(), []int{11}
}

func (x *DelKvResp) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *DelKvResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DelKvResp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_dbserver_proto protoreflect.FileDescriptor

var file_dbserver_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x24, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x23, 0x0a, 0x0d, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x53, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x4b, 0x76, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x4b, 0x76, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4b, 0x76, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x54, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4b, 0x76, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x4b, 0x76, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x54, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x4b, 0x76, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dbserver_proto_rawDescOnce sync.Once
	file_dbserver_proto_rawDescData = file_dbserver_proto_rawDesc
)

func file_dbserver_proto_rawDescGZIP() []byte {
	file_dbserver_proto_rawDescOnce.Do(func() {
		file_dbserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbserver_proto_rawDescData)
	})
	return file_dbserver_proto_rawDescData
}

var file_dbserver_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_dbserver_proto_goTypes = []interface{}{
	(*CreateSchemaReq)(nil),  // 0: dbserver.CreateSchemaReq
	(*CreateSchemaResp)(nil), // 1: dbserver.CreateSchemaResp
	(*QuerySchemaReq)(nil),   // 2: dbserver.QuerySchemaReq
	(*QuerySchemaResp)(nil),  // 3: dbserver.QuerySchemaResp
	(*DropSchemaReq)(nil),    // 4: dbserver.DropSchemaReq
	(*DropSchemaResp)(nil),   // 5: dbserver.DropSchemaResp
	(*PutKvReq)(nil),         // 6: dbserver.PutKvReq
	(*PutKvResp)(nil),        // 7: dbserver.PutKvResp
	(*GetKvReq)(nil),         // 8: dbserver.GetKvReq
	(*GetKvResp)(nil),        // 9: dbserver.GetKvResp
	(*DelKvReq)(nil),         // 10: dbserver.DelKvReq
	(*DelKvResp)(nil),        // 11: dbserver.DelKvResp
}
var file_dbserver_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dbserver_proto_init() }
func file_dbserver_proto_init() {
	if File_dbserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dbserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchemaReq); i {
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
		file_dbserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchemaResp); i {
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
		file_dbserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySchemaReq); i {
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
		file_dbserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySchemaResp); i {
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
		file_dbserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropSchemaReq); i {
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
		file_dbserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropSchemaResp); i {
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
		file_dbserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutKvReq); i {
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
		file_dbserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutKvResp); i {
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
		file_dbserver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKvReq); i {
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
		file_dbserver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKvResp); i {
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
		file_dbserver_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelKvReq); i {
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
		file_dbserver_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelKvResp); i {
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
			RawDescriptor: file_dbserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dbserver_proto_goTypes,
		DependencyIndexes: file_dbserver_proto_depIdxs,
		MessageInfos:      file_dbserver_proto_msgTypes,
	}.Build()
	File_dbserver_proto = out.File
	file_dbserver_proto_rawDesc = nil
	file_dbserver_proto_goTypes = nil
	file_dbserver_proto_depIdxs = nil
}
