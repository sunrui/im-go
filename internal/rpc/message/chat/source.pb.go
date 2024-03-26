//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 09:55:35

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internal/rpc/message/chat/source.proto

package chat

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

// 来源
type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceId string `protobuf:"bytes,1,opt,name=sequenceId,proto3" json:"sequenceId,omitempty"` // 序号
	// Types that are assignable to Id:
	//
	//	*Source_UserId
	//	*Source_GroupId
	//	*Source_RoomId
	//	*Source_BottleId
	Id        isSource_Id `protobuf_oneof:"Id"`
	Timestamp int64       `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // 时间戳
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_chat_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_chat_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_source_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetSequenceId() string {
	if x != nil {
		return x.SequenceId
	}
	return ""
}

func (m *Source) GetId() isSource_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *Source) GetUserId() string {
	if x, ok := x.GetId().(*Source_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *Source) GetGroupId() string {
	if x, ok := x.GetId().(*Source_GroupId); ok {
		return x.GroupId
	}
	return ""
}

func (x *Source) GetRoomId() string {
	if x, ok := x.GetId().(*Source_RoomId); ok {
		return x.RoomId
	}
	return ""
}

func (x *Source) GetBottleId() string {
	if x, ok := x.GetId().(*Source_BottleId); ok {
		return x.BottleId
	}
	return ""
}

func (x *Source) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type isSource_Id interface {
	isSource_Id()
}

type Source_UserId struct {
	UserId string `protobuf:"bytes,10,opt,name=userId,proto3,oneof"` // 用户 id
}

type Source_GroupId struct {
	GroupId string `protobuf:"bytes,11,opt,name=groupId,proto3,oneof"` // 群组 id
}

type Source_RoomId struct {
	RoomId string `protobuf:"bytes,12,opt,name=roomId,proto3,oneof"` // 房间 id
}

type Source_BottleId struct {
	BottleId string `protobuf:"bytes,13,opt,name=bottleId,proto3,oneof"` // 漂流瓶 id
}

func (*Source_UserId) isSource_Id() {}

func (*Source_GroupId) isSource_Id() {}

func (*Source_RoomId) isSource_Id() {}

func (*Source_BottleId) isSource_Id() {}

var File_internal_rpc_message_chat_source_proto protoreflect.FileDescriptor

var file_internal_rpc_message_chat_source_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x08, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x0a, 0x02, 0x49, 0x64,
	0x42, 0x1b, 0x5a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_message_chat_source_proto_rawDescOnce sync.Once
	file_internal_rpc_message_chat_source_proto_rawDescData = file_internal_rpc_message_chat_source_proto_rawDesc
)

func file_internal_rpc_message_chat_source_proto_rawDescGZIP() []byte {
	file_internal_rpc_message_chat_source_proto_rawDescOnce.Do(func() {
		file_internal_rpc_message_chat_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_message_chat_source_proto_rawDescData)
	})
	return file_internal_rpc_message_chat_source_proto_rawDescData
}

var file_internal_rpc_message_chat_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_rpc_message_chat_source_proto_goTypes = []interface{}{
	(*Source)(nil), // 0: internal.rpc.message.chat.Source
}
var file_internal_rpc_message_chat_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_rpc_message_chat_source_proto_init() }
func file_internal_rpc_message_chat_source_proto_init() {
	if File_internal_rpc_message_chat_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_rpc_message_chat_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
	file_internal_rpc_message_chat_source_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Source_UserId)(nil),
		(*Source_GroupId)(nil),
		(*Source_RoomId)(nil),
		(*Source_BottleId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_rpc_message_chat_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_rpc_message_chat_source_proto_goTypes,
		DependencyIndexes: file_internal_rpc_message_chat_source_proto_depIdxs,
		MessageInfos:      file_internal_rpc_message_chat_source_proto_msgTypes,
	}.Build()
	File_internal_rpc_message_chat_source_proto = out.File
	file_internal_rpc_message_chat_source_proto_rawDesc = nil
	file_internal_rpc_message_chat_source_proto_goTypes = nil
	file_internal_rpc_message_chat_source_proto_depIdxs = nil
}
