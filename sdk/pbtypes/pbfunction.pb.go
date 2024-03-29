// Copyright 2022 Listware

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: pbtypes/pbfunction.proto

package pbtypes

import (
	any "github.com/golang/protobuf/ptypes/any"
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

type FunctionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *FunctionType) Reset() {
	*x = FunctionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbtypes_pbfunction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionType) ProtoMessage() {}

func (x *FunctionType) ProtoReflect() protoreflect.Message {
	mi := &file_pbtypes_pbfunction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionType.ProtoReflect.Descriptor instead.
func (*FunctionType) Descriptor() ([]byte, []int) {
	return file_pbtypes_pbfunction_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionType) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FunctionType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type FunctionParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Default *any.Any `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *FunctionParameter) Reset() {
	*x = FunctionParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbtypes_pbfunction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionParameter) ProtoMessage() {}

func (x *FunctionParameter) ProtoReflect() protoreflect.Message {
	mi := &file_pbtypes_pbfunction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionParameter.ProtoReflect.Descriptor instead.
func (*FunctionParameter) Descriptor() ([]byte, []int) {
	return file_pbtypes_pbfunction_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionParameter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FunctionParameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FunctionParameter) GetDefault() *any.Any {
	if x != nil {
		return x.Default
	}
	return nil
}

// Function - is a CMDB object which represent Flink's stateful function
type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        *FunctionType        `protobuf:"bytes,1,opt,name=type,json=function_type,proto3" json:"type,omitempty"`
	Description string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Params      []*FunctionParameter `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbtypes_pbfunction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_pbtypes_pbfunction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_pbtypes_pbfunction_proto_rawDescGZIP(), []int{2}
}

func (x *Function) GetType() *FunctionType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Function) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Function) GetParams() []*FunctionParameter {
	if x != nil {
		return x.Params
	}
	return nil
}

type FunctionContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionType *FunctionType `protobuf:"bytes,1,opt,name=function_type,proto3" json:"function_type,omitempty"`
	Id           string        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Value        []byte        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	ReplyResult  *ReplyResult  `protobuf:"bytes,4,opt,name=reply_result,proto3" json:"reply_result,omitempty"`
}

func (x *FunctionContext) Reset() {
	*x = FunctionContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbtypes_pbfunction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionContext) ProtoMessage() {}

func (x *FunctionContext) ProtoReflect() protoreflect.Message {
	mi := &file_pbtypes_pbfunction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionContext.ProtoReflect.Descriptor instead.
func (*FunctionContext) Descriptor() ([]byte, []int) {
	return file_pbtypes_pbfunction_proto_rawDescGZIP(), []int{3}
}

func (x *FunctionContext) GetFunctionType() *FunctionType {
	if x != nil {
		return x.FunctionType
	}
	return nil
}

func (x *FunctionContext) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FunctionContext) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *FunctionContext) GetReplyResult() *ReplyResult {
	if x != nil {
		return x.ReplyResult
	}
	return nil
}

// FunctionRoute - is a LinkMessage from cmdb.Object to Function
type FunctionRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FunctionRoute) Reset() {
	*x = FunctionRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbtypes_pbfunction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionRoute) ProtoMessage() {}

func (x *FunctionRoute) ProtoReflect() protoreflect.Message {
	mi := &file_pbtypes_pbfunction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionRoute.ProtoReflect.Descriptor instead.
func (*FunctionRoute) Descriptor() ([]byte, []int) {
	return file_pbtypes_pbfunction_proto_rawDescGZIP(), []int{4}
}

func (x *FunctionRoute) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FunctionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Mountpoint string         `protobuf:"bytes,2,opt,name=mountpoint,proto3" json:"mountpoint,omitempty"`
	Route      *FunctionRoute `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *FunctionMessage) Reset() {
	*x = FunctionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbtypes_pbfunction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionMessage) ProtoMessage() {}

func (x *FunctionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pbtypes_pbfunction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionMessage.ProtoReflect.Descriptor instead.
func (*FunctionMessage) Descriptor() ([]byte, []int) {
	return file_pbtypes_pbfunction_proto_rawDescGZIP(), []int{5}
}

func (x *FunctionMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FunctionMessage) GetMountpoint() string {
	if x != nil {
		return x.Mountpoint
	}
	return ""
}

func (x *FunctionMessage) GetRoute() *FunctionRoute {
	if x != nil {
		return x.Route
	}
	return nil
}

var File_pbtypes_pbfunction_proto protoreflect.FileDescriptor

var file_pbtypes_pbfunction_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x70, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x11, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x62,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0xd0, 0x01, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x62,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x70, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x21, 0x0a, 0x0d, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3d, 0x0a,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x70, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x48, 0x0a, 0x10,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x64, 0x6b,
	0x42, 0x09, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x2e, 0x66, 0x67, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x72, 0x75, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbtypes_pbfunction_proto_rawDescOnce sync.Once
	file_pbtypes_pbfunction_proto_rawDescData = file_pbtypes_pbfunction_proto_rawDesc
)

func file_pbtypes_pbfunction_proto_rawDescGZIP() []byte {
	file_pbtypes_pbfunction_proto_rawDescOnce.Do(func() {
		file_pbtypes_pbfunction_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbtypes_pbfunction_proto_rawDescData)
	})
	return file_pbtypes_pbfunction_proto_rawDescData
}

var file_pbtypes_pbfunction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pbtypes_pbfunction_proto_goTypes = []interface{}{
	(*FunctionType)(nil),      // 0: org.listware.sdk.pbtypes.FunctionType
	(*FunctionParameter)(nil), // 1: org.listware.sdk.pbtypes.FunctionParameter
	(*Function)(nil),          // 2: org.listware.sdk.pbtypes.Function
	(*FunctionContext)(nil),   // 3: org.listware.sdk.pbtypes.FunctionContext
	(*FunctionRoute)(nil),     // 4: org.listware.sdk.pbtypes.FunctionRoute
	(*FunctionMessage)(nil),   // 5: org.listware.sdk.pbtypes.FunctionMessage
	(*any.Any)(nil),           // 6: google.protobuf.Any
	(*ReplyResult)(nil),       // 7: org.listware.sdk.pbtypes.ReplyResult
}
var file_pbtypes_pbfunction_proto_depIdxs = []int32{
	6, // 0: org.listware.sdk.pbtypes.FunctionParameter.default:type_name -> google.protobuf.Any
	0, // 1: org.listware.sdk.pbtypes.Function.type:type_name -> org.listware.sdk.pbtypes.FunctionType
	1, // 2: org.listware.sdk.pbtypes.Function.params:type_name -> org.listware.sdk.pbtypes.FunctionParameter
	0, // 3: org.listware.sdk.pbtypes.FunctionContext.function_type:type_name -> org.listware.sdk.pbtypes.FunctionType
	7, // 4: org.listware.sdk.pbtypes.FunctionContext.reply_result:type_name -> org.listware.sdk.pbtypes.ReplyResult
	4, // 5: org.listware.sdk.pbtypes.FunctionMessage.route:type_name -> org.listware.sdk.pbtypes.FunctionRoute
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pbtypes_pbfunction_proto_init() }
func file_pbtypes_pbfunction_proto_init() {
	if File_pbtypes_pbfunction_proto != nil {
		return
	}
	file_pbtypes_pbresult_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbtypes_pbfunction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionType); i {
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
		file_pbtypes_pbfunction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionParameter); i {
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
		file_pbtypes_pbfunction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
		file_pbtypes_pbfunction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionContext); i {
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
		file_pbtypes_pbfunction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionRoute); i {
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
		file_pbtypes_pbfunction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionMessage); i {
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
			RawDescriptor: file_pbtypes_pbfunction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbtypes_pbfunction_proto_goTypes,
		DependencyIndexes: file_pbtypes_pbfunction_proto_depIdxs,
		MessageInfos:      file_pbtypes_pbfunction_proto_msgTypes,
	}.Build()
	File_pbtypes_pbfunction_proto = out.File
	file_pbtypes_pbfunction_proto_rawDesc = nil
	file_pbtypes_pbfunction_proto_goTypes = nil
	file_pbtypes_pbfunction_proto_depIdxs = nil
}
