//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 12:14:01

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/proto/p2p_chat/message_reject.proto

package p2p_chat

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

// 消息拒绝
type MessageReject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceId string `protobuf:"bytes,1,opt,name=sequenceId,proto3" json:"sequenceId,omitempty"` // 序号
}

func (x *MessageReject) Reset() {
	*x = MessageReject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_p2p_chat_message_reject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReject) ProtoMessage() {}

func (x *MessageReject) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_p2p_chat_message_reject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReject.ProtoReflect.Descriptor instead.
func (*MessageReject) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescGZIP(), []int{0}
}

func (x *MessageReject) GetSequenceId() string {
	if x != nil {
		return x.SequenceId
	}
	return ""
}

var File_pkg_rpc_proto_p2p_chat_message_reject_proto protoreflect.FileDescriptor

var file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70,
	0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x22, 0x2f, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescOnce sync.Once
	file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescData = file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDesc
)

func file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescGZIP() []byte {
	file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescData)
	})
	return file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDescData
}

var file_pkg_rpc_proto_p2p_chat_message_reject_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_rpc_proto_p2p_chat_message_reject_proto_goTypes = []interface{}{
	(*MessageReject)(nil), // 0: pkg.rpc.proto.p2p_chat.MessageReject
}
var file_pkg_rpc_proto_p2p_chat_message_reject_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_rpc_proto_p2p_chat_message_reject_proto_init() }
func file_pkg_rpc_proto_p2p_chat_message_reject_proto_init() {
	if File_pkg_rpc_proto_p2p_chat_message_reject_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_proto_p2p_chat_message_reject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReject); i {
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
			RawDescriptor: file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_rpc_proto_p2p_chat_message_reject_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_proto_p2p_chat_message_reject_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_proto_p2p_chat_message_reject_proto_msgTypes,
	}.Build()
	File_pkg_rpc_proto_p2p_chat_message_reject_proto = out.File
	file_pkg_rpc_proto_p2p_chat_message_reject_proto_rawDesc = nil
	file_pkg_rpc_proto_p2p_chat_message_reject_proto_goTypes = nil
	file_pkg_rpc_proto_p2p_chat_message_reject_proto_depIdxs = nil
}
