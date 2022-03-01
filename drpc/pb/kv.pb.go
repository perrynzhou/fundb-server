//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.0
// source: kv.proto

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
		mi := &file_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaReq) ProtoMessage() {}

func (x *CreateSchemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[0]
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
	return file_kv_proto_rawDescGZIP(), []int{0}
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
		mi := &file_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaResp) ProtoMessage() {}

func (x *CreateSchemaResp) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[1]
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
	return file_kv_proto_rawDescGZIP(), []int{1}
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
		mi := &file_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySchemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySchemaReq) ProtoMessage() {}

func (x *QuerySchemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[2]
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
	return file_kv_proto_rawDescGZIP(), []int{2}
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

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Confs []string `protobuf:"bytes,2,rep,name=confs,proto3" json:"confs,omitempty"`
}

func (x *QuerySchemaResp) Reset() {
	*x = QuerySchemaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySchemaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySchemaResp) ProtoMessage() {}

func (x *QuerySchemaResp) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[3]
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
	return file_kv_proto_rawDescGZIP(), []int{3}
}

func (x *QuerySchemaResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QuerySchemaResp) GetConfs() []string {
	if x != nil {
		return x.Confs
	}
	return nil
}

type DeleteSchemaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSchemaReq) Reset() {
	*x = DeleteSchemaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSchemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSchemaReq) ProtoMessage() {}

func (x *DeleteSchemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSchemaReq.ProtoReflect.Descriptor instead.
func (*DeleteSchemaReq) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSchemaReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteSchemaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteSchemaResp) Reset() {
	*x = DeleteSchemaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSchemaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSchemaResp) ProtoMessage() {}

func (x *DeleteSchemaResp) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSchemaResp.ProtoReflect.Descriptor instead.
func (*DeleteSchemaResp) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSchemaResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteSchemaResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PutConfReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	ConfKey    string `protobuf:"bytes,2,opt,name=conf_key,json=confKey,proto3" json:"conf_key,omitempty"`
	ConfVal    string `protobuf:"bytes,3,opt,name=conf_val,json=confVal,proto3" json:"conf_val,omitempty"`
}

func (x *PutConfReq) Reset() {
	*x = PutConfReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutConfReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutConfReq) ProtoMessage() {}

func (x *PutConfReq) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutConfReq.ProtoReflect.Descriptor instead.
func (*PutConfReq) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{6}
}

func (x *PutConfReq) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *PutConfReq) GetConfKey() string {
	if x != nil {
		return x.ConfKey
	}
	return ""
}

func (x *PutConfReq) GetConfVal() string {
	if x != nil {
		return x.ConfVal
	}
	return ""
}

type PutConfResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PutConfResp) Reset() {
	*x = PutConfResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutConfResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutConfResp) ProtoMessage() {}

func (x *PutConfResp) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutConfResp.ProtoReflect.Descriptor instead.
func (*PutConfResp) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{7}
}

func (x *PutConfResp) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *PutConfResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetConfReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	ConfKey    string `protobuf:"bytes,2,opt,name=conf_key,json=confKey,proto3" json:"conf_key,omitempty"`
}

func (x *GetConfReq) Reset() {
	*x = GetConfReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfReq) ProtoMessage() {}

func (x *GetConfReq) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfReq.ProtoReflect.Descriptor instead.
func (*GetConfReq) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{8}
}

func (x *GetConfReq) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *GetConfReq) GetConfKey() string {
	if x != nil {
		return x.ConfKey
	}
	return ""
}

type GetConfResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	ConfKey    string `protobuf:"bytes,2,opt,name=conf_key,json=confKey,proto3" json:"conf_key,omitempty"`
	ConfVal    string `protobuf:"bytes,3,opt,name=conf_val,json=confVal,proto3" json:"conf_val,omitempty"`
}

func (x *GetConfResp) Reset() {
	*x = GetConfResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfResp) ProtoMessage() {}

func (x *GetConfResp) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfResp.ProtoReflect.Descriptor instead.
func (*GetConfResp) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{9}
}

func (x *GetConfResp) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *GetConfResp) GetConfKey() string {
	if x != nil {
		return x.ConfKey
	}
	return ""
}

func (x *GetConfResp) GetConfVal() string {
	if x != nil {
		return x.ConfVal
	}
	return ""
}

type DeleteConfReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	ConfKey    string `protobuf:"bytes,2,opt,name=conf_key,json=confKey,proto3" json:"conf_key,omitempty"`
}

func (x *DeleteConfReq) Reset() {
	*x = DeleteConfReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfReq) ProtoMessage() {}

func (x *DeleteConfReq) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfReq.ProtoReflect.Descriptor instead.
func (*DeleteConfReq) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteConfReq) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *DeleteConfReq) GetConfKey() string {
	if x != nil {
		return x.ConfKey
	}
	return ""
}

type DeleteConfResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	ConfKey    string `protobuf:"bytes,2,opt,name=conf_key,json=confKey,proto3" json:"conf_key,omitempty"`
	ConfVal    string `protobuf:"bytes,3,opt,name=conf_val,json=confVal,proto3" json:"conf_val,omitempty"`
}

func (x *DeleteConfResp) Reset() {
	*x = DeleteConfResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfResp) ProtoMessage() {}

func (x *DeleteConfResp) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfResp.ProtoReflect.Descriptor instead.
func (*DeleteConfResp) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteConfResp) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *DeleteConfResp) GetConfKey() string {
	if x != nil {
		return x.ConfKey
	}
	return ""
}

func (x *DeleteConfResp) GetConfVal() string {
	if x != nil {
		return x.ConfVal
	}
	return ""
}

var File_kv_proto protoreflect.FileDescriptor

var file_kv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x6b, 0x76, 0x22, 0x25,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0f, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6e, 0x66, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x63, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x4b, 0x65,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x56, 0x61, 0x6c, 0x22, 0x40, 0x0a, 0x0b,
	0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x48,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x4b, 0x65, 0x79, 0x22, 0x64, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66,
	0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x56, 0x61, 0x6c, 0x22, 0x4b,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x4b, 0x65, 0x79, 0x22, 0x67, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x66, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x66, 0x56, 0x61, 0x6c, 0x32, 0xd2, 0x02, 0x0a, 0x09, 0x4b, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x13, 0x2e, 0x6b, 0x76, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6b, 0x76, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x12,
	0x2e, 0x6b, 0x76, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6b, 0x76, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x13, 0x2e, 0x6b, 0x76, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x6b, 0x76, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x0e, 0x2e, 0x6b, 0x76, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x6b, 0x76, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x12,
	0x0e, 0x2e, 0x6b, 0x76, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x6b, 0x76, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x11, 0x2e, 0x6b, 0x76, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6b, 0x76, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_proto_rawDescOnce sync.Once
	file_kv_proto_rawDescData = file_kv_proto_rawDesc
)

func file_kv_proto_rawDescGZIP() []byte {
	file_kv_proto_rawDescOnce.Do(func() {
		file_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_proto_rawDescData)
	})
	return file_kv_proto_rawDescData
}

var file_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_kv_proto_goTypes = []interface{}{
	(*CreateSchemaReq)(nil),  // 0: kv.CreateSchemaReq
	(*CreateSchemaResp)(nil), // 1: kv.CreateSchemaResp
	(*QuerySchemaReq)(nil),   // 2: kv.QuerySchemaReq
	(*QuerySchemaResp)(nil),  // 3: kv.QuerySchemaResp
	(*DeleteSchemaReq)(nil),  // 4: kv.DeleteSchemaReq
	(*DeleteSchemaResp)(nil), // 5: kv.DeleteSchemaResp
	(*PutConfReq)(nil),       // 6: kv.PutConfReq
	(*PutConfResp)(nil),      // 7: kv.PutConfResp
	(*GetConfReq)(nil),       // 8: kv.GetConfReq
	(*GetConfResp)(nil),      // 9: kv.GetConfResp
	(*DeleteConfReq)(nil),    // 10: kv.DeleteConfReq
	(*DeleteConfResp)(nil),   // 11: kv.DeleteConfResp
}
var file_kv_proto_depIdxs = []int32{
	0,  // 0: kv.KvService.CreateSchema:input_type -> kv.CreateSchemaReq
	2,  // 1: kv.KvService.QuerySchema:input_type -> kv.QuerySchemaReq
	4,  // 2: kv.KvService.DeleteSchema:input_type -> kv.DeleteSchemaReq
	6,  // 3: kv.KvService.PutConf:input_type -> kv.PutConfReq
	8,  // 4: kv.KvService.GetConf:input_type -> kv.GetConfReq
	10, // 5: kv.KvService.DeleteConf:input_type -> kv.DeleteConfReq
	1,  // 6: kv.KvService.CreateSchema:output_type -> kv.CreateSchemaResp
	3,  // 7: kv.KvService.QuerySchema:output_type -> kv.QuerySchemaResp
	5,  // 8: kv.KvService.DeleteSchema:output_type -> kv.DeleteSchemaResp
	7,  // 9: kv.KvService.PutConf:output_type -> kv.PutConfResp
	9,  // 10: kv.KvService.GetConf:output_type -> kv.GetConfResp
	11, // 11: kv.KvService.DeleteConf:output_type -> kv.DeleteConfResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_kv_proto_init() }
func file_kv_proto_init() {
	if File_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_kv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSchemaReq); i {
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
		file_kv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSchemaResp); i {
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
		file_kv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutConfReq); i {
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
		file_kv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutConfResp); i {
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
		file_kv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfReq); i {
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
		file_kv_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfResp); i {
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
		file_kv_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfReq); i {
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
		file_kv_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfResp); i {
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
			RawDescriptor: file_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_proto_goTypes,
		DependencyIndexes: file_kv_proto_depIdxs,
		MessageInfos:      file_kv_proto_msgTypes,
	}.Build()
	File_kv_proto = out.File
	file_kv_proto_rawDesc = nil
	file_kv_proto_goTypes = nil
	file_kv_proto_depIdxs = nil
}
