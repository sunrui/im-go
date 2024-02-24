//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-02-24 22:39:45

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: internal/rpc/push/type.proto

package push

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

// 消息类型
type Type int32

const (
	Type_TEXT                 Type = 0   // 文本
	Type_STICKER              Type = 1   // 表情包
	Type_QUOTED               Type = 2   // 引用
	Type_IMAGE                Type = 3   // 图片
	Type_AUDIO                Type = 4   // 音频
	Type_VIDEO                Type = 5   // 视频
	Type_FILE                 Type = 6   // 文件
	Type_LOCATION             Type = 7   // 定位
	Type_CARD                 Type = 8   // 名片
	Type_VOICE_CALL           Type = 9   // 语音通话
	Type_VIDEO_CALL           Type = 10  // 视频通话
	Type_FRIEND_ONLINE        Type = 101 // 好友上线
	Type_FRIEND_OFFLINE       Type = 102 // 好友下线
	Type_FRIEND_INPUT_ING     Type = 103 // 好友正在输入
	Type_FRIEND_INPUT_END     Type = 104 // 好友停止输入
	Type_FRIEND_ADD_REQUEST   Type = 105 // 好友添加请求
	Type_FRIEND_ADD_REPLY     Type = 106 // 好友添加回复
	Type_GROUP_MESSAGE        Type = 200 // 群消息
	Type_GROUP_MEMBER_ONLINE  Type = 201 // 群成员在线
	Type_GROUP_MEMBER_OFFLINE Type = 202 // 群成员在线
	Type_GROUP_MEMBER_JOIN    Type = 203 // 群成员加入
	Type_GROUP_MEMBER_LEAVE   Type = 204 // 群成员离开
	Type_GROUP_MEMBER_MUTE    Type = 205 // 群成员禁言
	Type_GROUP_MEMBER_UNMUTE  Type = 206 // 群成员解除禁言
	Type_ROOM_MESSAGE         Type = 300 // 聊天室消息
	Type_ROOM_MEMBER_JOIN     Type = 301 // 聊天室成员加入
	Type_ROOM_MEMBER_LEAVE    Type = 302 // 聊天室成员离开
	Type_BOTTLE_MESSAGE       Type = 400 // 飘流瓶消息
	Type_SYSTEM_NOTICE        Type = 900 // 系统通知
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:   "TEXT",
		1:   "STICKER",
		2:   "QUOTED",
		3:   "IMAGE",
		4:   "AUDIO",
		5:   "VIDEO",
		6:   "FILE",
		7:   "LOCATION",
		8:   "CARD",
		9:   "VOICE_CALL",
		10:  "VIDEO_CALL",
		101: "FRIEND_ONLINE",
		102: "FRIEND_OFFLINE",
		103: "FRIEND_INPUT_ING",
		104: "FRIEND_INPUT_END",
		105: "FRIEND_ADD_REQUEST",
		106: "FRIEND_ADD_REPLY",
		200: "GROUP_MESSAGE",
		201: "GROUP_MEMBER_ONLINE",
		202: "GROUP_MEMBER_OFFLINE",
		203: "GROUP_MEMBER_JOIN",
		204: "GROUP_MEMBER_LEAVE",
		205: "GROUP_MEMBER_MUTE",
		206: "GROUP_MEMBER_UNMUTE",
		300: "ROOM_MESSAGE",
		301: "ROOM_MEMBER_JOIN",
		302: "ROOM_MEMBER_LEAVE",
		400: "BOTTLE_MESSAGE",
		900: "SYSTEM_NOTICE",
	}
	Type_value = map[string]int32{
		"TEXT":                 0,
		"STICKER":              1,
		"QUOTED":               2,
		"IMAGE":                3,
		"AUDIO":                4,
		"VIDEO":                5,
		"FILE":                 6,
		"LOCATION":             7,
		"CARD":                 8,
		"VOICE_CALL":           9,
		"VIDEO_CALL":           10,
		"FRIEND_ONLINE":        101,
		"FRIEND_OFFLINE":       102,
		"FRIEND_INPUT_ING":     103,
		"FRIEND_INPUT_END":     104,
		"FRIEND_ADD_REQUEST":   105,
		"FRIEND_ADD_REPLY":     106,
		"GROUP_MESSAGE":        200,
		"GROUP_MEMBER_ONLINE":  201,
		"GROUP_MEMBER_OFFLINE": 202,
		"GROUP_MEMBER_JOIN":    203,
		"GROUP_MEMBER_LEAVE":   204,
		"GROUP_MEMBER_MUTE":    205,
		"GROUP_MEMBER_UNMUTE":  206,
		"ROOM_MESSAGE":         300,
		"ROOM_MEMBER_JOIN":     301,
		"ROOM_MEMBER_LEAVE":    302,
		"BOTTLE_MESSAGE":       400,
		"SYSTEM_NOTICE":        900,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_rpc_push_type_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_internal_rpc_push_type_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_internal_rpc_push_type_proto_rawDescGZIP(), []int{0}
}

var File_internal_rpc_push_type_proto protoreflect.FileDescriptor

var file_internal_rpc_push_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x2a, 0xa4, 0x04, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x49, 0x43, 0x4b,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x55, 0x44, 0x49, 0x4f, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x52,
	0x44, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x4e,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44,
	0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x66, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x52,
	0x49, 0x45, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x49, 0x4e, 0x47, 0x10, 0x67,
	0x12, 0x14, 0x0a, 0x10, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54,
	0x5f, 0x45, 0x4e, 0x44, 0x10, 0x68, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44,
	0x5f, 0x41, 0x44, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x69, 0x12, 0x14,
	0x0a, 0x10, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x52, 0x45, 0x50,
	0x4c, 0x59, 0x10, 0x6a, 0x12, 0x12, 0x0a, 0x0d, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0xc8, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10,
	0xc9, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0xca, 0x01, 0x12, 0x16, 0x0a,
	0x11, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4a, 0x4f,
	0x49, 0x4e, 0x10, 0xcb, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d,
	0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0xcc, 0x01, 0x12, 0x16,
	0x0a, 0x11, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4d,
	0x55, 0x54, 0x45, 0x10, 0xcd, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f,
	0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4d, 0x55, 0x54, 0x45, 0x10, 0xce, 0x01,
	0x12, 0x11, 0x0a, 0x0c, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x10, 0xac, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0xad, 0x02, 0x12, 0x16, 0x0a, 0x11, 0x52, 0x4f,
	0x4f, 0x4d, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10,
	0xae, 0x02, 0x12, 0x13, 0x0a, 0x0e, 0x42, 0x4f, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x90, 0x03, 0x12, 0x12, 0x0a, 0x0d, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x10, 0x84, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2f,
	0x70, 0x75, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_push_type_proto_rawDescOnce sync.Once
	file_internal_rpc_push_type_proto_rawDescData = file_internal_rpc_push_type_proto_rawDesc
)

func file_internal_rpc_push_type_proto_rawDescGZIP() []byte {
	file_internal_rpc_push_type_proto_rawDescOnce.Do(func() {
		file_internal_rpc_push_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_push_type_proto_rawDescData)
	})
	return file_internal_rpc_push_type_proto_rawDescData
}

var file_internal_rpc_push_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_rpc_push_type_proto_goTypes = []interface{}{
	(Type)(0), // 0: auth.Type
}
var file_internal_rpc_push_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_rpc_push_type_proto_init() }
func file_internal_rpc_push_type_proto_init() {
	if File_internal_rpc_push_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_rpc_push_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_rpc_push_type_proto_goTypes,
		DependencyIndexes: file_internal_rpc_push_type_proto_depIdxs,
		EnumInfos:         file_internal_rpc_push_type_proto_enumTypes,
	}.Build()
	File_internal_rpc_push_type_proto = out.File
	file_internal_rpc_push_type_proto_rawDesc = nil
	file_internal_rpc_push_type_proto_goTypes = nil
	file_internal_rpc_push_type_proto_depIdxs = nil
}
