//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 13:59:16

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internal/rpc/message/bottle_chat/bottle_chat.proto

package bottle_chat

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

// 订阅请求
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescGZIP(), []int{0}
}

// 订阅回复
type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"` // 用户 id
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReply.ProtoReflect.Descriptor instead.
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeReply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_internal_rpc_message_bottle_chat_bottle_chat_proto protoreflect.FileDescriptor

var file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x62, 0x6f, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x32, 0x81, 0x01, 0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x73, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x62, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescOnce sync.Once
	file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescData = file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDesc
)

func file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescGZIP() []byte {
	file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescOnce.Do(func() {
		file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescData)
	})
	return file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDescData
}

var file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_rpc_message_bottle_chat_bottle_chat_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil), // 0: internal.rpc.message.bottle_chat.SubscribeRequest
	(*SubscribeReply)(nil),   // 1: internal.rpc.message.bottle_chat.SubscribeReply
}
var file_internal_rpc_message_bottle_chat_bottle_chat_proto_depIdxs = []int32{
	0, // 0: internal.rpc.message.bottle_chat.BottleChat.Subscribe:input_type -> internal.rpc.message.bottle_chat.SubscribeRequest
	1, // 1: internal.rpc.message.bottle_chat.BottleChat.Subscribe:output_type -> internal.rpc.message.bottle_chat.SubscribeReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_rpc_message_bottle_chat_bottle_chat_proto_init() }
func file_internal_rpc_message_bottle_chat_bottle_chat_proto_init() {
	if File_internal_rpc_message_bottle_chat_bottle_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReply); i {
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
			RawDescriptor: file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_rpc_message_bottle_chat_bottle_chat_proto_goTypes,
		DependencyIndexes: file_internal_rpc_message_bottle_chat_bottle_chat_proto_depIdxs,
		MessageInfos:      file_internal_rpc_message_bottle_chat_bottle_chat_proto_msgTypes,
	}.Build()
	File_internal_rpc_message_bottle_chat_bottle_chat_proto = out.File
	file_internal_rpc_message_bottle_chat_bottle_chat_proto_rawDesc = nil
	file_internal_rpc_message_bottle_chat_bottle_chat_proto_goTypes = nil
	file_internal_rpc_message_bottle_chat_bottle_chat_proto_depIdxs = nil
}
