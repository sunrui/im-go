//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-02-24 23:39:45

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: internal/rpc/common/reply.proto

package common

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

// 回复代码
type ReplyCode int32

const (
	ReplyCode_OK                  ReplyCode = 0 // 成功
	ReplyCode_NO_AUTH             ReplyCode = 1 // 未登录
	ReplyCode_PARAMETER_ERROR     ReplyCode = 2 // 参数错误
	ReplyCode_CHAT_USER_NOT_FOUND ReplyCode = 3 // 用户不存在
	ReplyCode_CHAT_USER_BLOCKED   ReplyCode = 4 // 用户已拉黑
	ReplyCode_RATE_LIMIT          ReplyCode = 5 // 限流
	ReplyCode_INTERNAL_ERROR      ReplyCode = 6 // 内部错误
)

// Enum value maps for ReplyCode.
var (
	ReplyCode_name = map[int32]string{
		0: "OK",
		1: "NO_AUTH",
		2: "PARAMETER_ERROR",
		3: "CHAT_USER_NOT_FOUND",
		4: "CHAT_USER_BLOCKED",
		5: "RATE_LIMIT",
		6: "INTERNAL_ERROR",
	}
	ReplyCode_value = map[string]int32{
		"OK":                  0,
		"NO_AUTH":             1,
		"PARAMETER_ERROR":     2,
		"CHAT_USER_NOT_FOUND": 3,
		"CHAT_USER_BLOCKED":   4,
		"RATE_LIMIT":          5,
		"INTERNAL_ERROR":      6,
	}
)

func (x ReplyCode) Enum() *ReplyCode {
	p := new(ReplyCode)
	*p = x
	return p
}

func (x ReplyCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReplyCode) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_rpc_common_reply_proto_enumTypes[0].Descriptor()
}

func (ReplyCode) Type() protoreflect.EnumType {
	return &file_internal_rpc_common_reply_proto_enumTypes[0]
}

func (x ReplyCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReplyCode.Descriptor instead.
func (ReplyCode) EnumDescriptor() ([]byte, []int) {
	return file_internal_rpc_common_reply_proto_rawDescGZIP(), []int{0}
}

// 回复
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ReplyCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.ReplyCode" json:"code,omitempty"` // 代码
	Message *string   `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`            // 消息
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_common_reply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_common_reply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_internal_rpc_common_reply_proto_rawDescGZIP(), []int{0}
}

func (x *Reply) GetCode() ReplyCode {
	if x != nil {
		return x.Code
	}
	return ReplyCode_OK
}

func (x *Reply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_internal_rpc_common_reply_proto protoreflect.FileDescriptor

var file_internal_rpc_common_reply_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x05, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2a, 0x89, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f,
	0x5f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x45, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x48, 0x41, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06,
	0x42, 0x15, 0x5a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_common_reply_proto_rawDescOnce sync.Once
	file_internal_rpc_common_reply_proto_rawDescData = file_internal_rpc_common_reply_proto_rawDesc
)

func file_internal_rpc_common_reply_proto_rawDescGZIP() []byte {
	file_internal_rpc_common_reply_proto_rawDescOnce.Do(func() {
		file_internal_rpc_common_reply_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_common_reply_proto_rawDescData)
	})
	return file_internal_rpc_common_reply_proto_rawDescData
}

var file_internal_rpc_common_reply_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_rpc_common_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_rpc_common_reply_proto_goTypes = []interface{}{
	(ReplyCode)(0), // 0: common.ReplyCode
	(*Reply)(nil),  // 1: common.Reply
}
var file_internal_rpc_common_reply_proto_depIdxs = []int32{
	0, // 0: common.Reply.code:type_name -> common.ReplyCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_rpc_common_reply_proto_init() }
func file_internal_rpc_common_reply_proto_init() {
	if File_internal_rpc_common_reply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_rpc_common_reply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
	file_internal_rpc_common_reply_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_rpc_common_reply_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_rpc_common_reply_proto_goTypes,
		DependencyIndexes: file_internal_rpc_common_reply_proto_depIdxs,
		EnumInfos:         file_internal_rpc_common_reply_proto_enumTypes,
		MessageInfos:      file_internal_rpc_common_reply_proto_msgTypes,
	}.Build()
	File_internal_rpc_common_reply_proto = out.File
	file_internal_rpc_common_reply_proto_rawDesc = nil
	file_internal_rpc_common_reply_proto_goTypes = nil
	file_internal_rpc_common_reply_proto_depIdxs = nil
}
