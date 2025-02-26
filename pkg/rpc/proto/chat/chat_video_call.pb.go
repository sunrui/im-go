//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 10:57:20

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/proto/chat/chat_video_call.proto

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

// 定义一个名为VideoCallStatus的枚举类型，用于描述视频通话的不同状态。
type ChatVideoCall_Status int32

const (
	// 发出了视频通话请求，等待对方响应。
	ChatVideoCall_PENDING_REQUEST ChatVideoCall_Status = 0
	// 视频通话正在进行中，音频和视频流均正常传输。
	ChatVideoCall_ACTIVE ChatVideoCall_Status = 1
	// 对方拒绝了视频通话请求。
	ChatVideoCall_REJECTED ChatVideoCall_Status = 2
	// 视频通话请求在等待对方响应时超时。
	ChatVideoCall_TIMEOUT ChatVideoCall_Status = 3
	// 视频通话被系统规则或对方设置禁止。
	ChatVideoCall_FORBIDDEN ChatVideoCall_Status = 4
	// 主叫方在对方接受之前主动取消了视频通话。
	ChatVideoCall_CANCELLED_BY_CALLER ChatVideoCall_Status = 5
	// 被叫方在通话过程中主动结束视频通话。
	ChatVideoCall_ENDED_BY_CALLEE ChatVideoCall_Status = 6
	// 视频通话因网络问题、设备故障或其他技术原因意外中断。
	ChatVideoCall_TECHNICAL_ISSUE ChatVideoCall_Status = 7
	// 视频通话因违反服务条款、账户限制等原因被系统强制结束。
	ChatVideoCall_TERMINATED_BY_SYSTEM ChatVideoCall_Status = 8
	// 视频通话已进入保留状态，音频和/或视频暂时暂停。
	ChatVideoCall_ON_HOLD ChatVideoCall_Status = 9
	// 通话中的一方或双方已启用音频静音。
	ChatVideoCall_AUDIO_MUTED ChatVideoCall_Status = 10
	// 通话中的一方或双方已关闭视频流。
	ChatVideoCall_VIDEO_DISABLED ChatVideoCall_Status = 11
	// 视频通话正在进行转移操作，即将连接到另一个用户或设备。
	ChatVideoCall_TRANSFER_IN_PROGRESS ChatVideoCall_Status = 12
)

// Enum value maps for ChatVideoCall_Status.
var (
	ChatVideoCall_Status_name = map[int32]string{
		0:  "PENDING_REQUEST",
		1:  "ACTIVE",
		2:  "REJECTED",
		3:  "TIMEOUT",
		4:  "FORBIDDEN",
		5:  "CANCELLED_BY_CALLER",
		6:  "ENDED_BY_CALLEE",
		7:  "TECHNICAL_ISSUE",
		8:  "TERMINATED_BY_SYSTEM",
		9:  "ON_HOLD",
		10: "AUDIO_MUTED",
		11: "VIDEO_DISABLED",
		12: "TRANSFER_IN_PROGRESS",
	}
	ChatVideoCall_Status_value = map[string]int32{
		"PENDING_REQUEST":      0,
		"ACTIVE":               1,
		"REJECTED":             2,
		"TIMEOUT":              3,
		"FORBIDDEN":            4,
		"CANCELLED_BY_CALLER":  5,
		"ENDED_BY_CALLEE":      6,
		"TECHNICAL_ISSUE":      7,
		"TERMINATED_BY_SYSTEM": 8,
		"ON_HOLD":              9,
		"AUDIO_MUTED":          10,
		"VIDEO_DISABLED":       11,
		"TRANSFER_IN_PROGRESS": 12,
	}
)

func (x ChatVideoCall_Status) Enum() *ChatVideoCall_Status {
	p := new(ChatVideoCall_Status)
	*p = x
	return p
}

func (x ChatVideoCall_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatVideoCall_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_rpc_proto_chat_chat_video_call_proto_enumTypes[0].Descriptor()
}

func (ChatVideoCall_Status) Type() protoreflect.EnumType {
	return &file_pkg_rpc_proto_chat_chat_video_call_proto_enumTypes[0]
}

func (x ChatVideoCall_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatVideoCall_Status.Descriptor instead.
func (ChatVideoCall_Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescGZIP(), []int{0, 0}
}

// 视频通话
type ChatVideoCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalleeId               string `protobuf:"bytes,1,opt,name=callee_id,json=calleeId,proto3" json:"callee_id,omitempty"`                                                // 对方的唯一标识，如用户ID或电话号码
	StartTimeUnixTimestamp int64  `protobuf:"varint,2,opt,name=start_time_unix_timestamp,json=startTimeUnixTimestamp,proto3" json:"start_time_unix_timestamp,omitempty"` // 通话开始时间戳（Unix时间戳）
	DurationSeconds        int32  `protobuf:"varint,3,opt,name=duration_seconds,json=durationSeconds,proto3" json:"duration_seconds,omitempty"`                          // 通话持续时长（秒）
	Protocol               string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`                                                                // 通话使用的协议类型（例如：SIP、WebRTC等）
	SessionUrlOrToken      string `protobuf:"bytes,5,opt,name=session_url_or_token,json=sessionUrlOrToken,proto3" json:"session_url_or_token,omitempty"`                 // （可选）如果是VoIP服务，可能是用于建立通话会话的服务器URL或令牌
}

func (x *ChatVideoCall) Reset() {
	*x = ChatVideoCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_video_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatVideoCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatVideoCall) ProtoMessage() {}

func (x *ChatVideoCall) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_video_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatVideoCall.ProtoReflect.Descriptor instead.
func (*ChatVideoCall) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescGZIP(), []int{0}
}

func (x *ChatVideoCall) GetCalleeId() string {
	if x != nil {
		return x.CalleeId
	}
	return ""
}

func (x *ChatVideoCall) GetStartTimeUnixTimestamp() int64 {
	if x != nil {
		return x.StartTimeUnixTimestamp
	}
	return 0
}

func (x *ChatVideoCall) GetDurationSeconds() int32 {
	if x != nil {
		return x.DurationSeconds
	}
	return 0
}

func (x *ChatVideoCall) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ChatVideoCall) GetSessionUrlOrToken() string {
	if x != nil {
		return x.SessionUrlOrToken
	}
	return ""
}

var File_pkg_rpc_proto_chat_chat_video_call_proto protoreflect.FileDescriptor

var file_pkg_rpc_proto_chat_chat_video_call_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6b, 0x67, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x22, 0xde,
	0x03, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x19, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x16, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x2f, 0x0a, 0x14, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6f,
	0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x4f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xfc, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49,
	0x44, 0x44, 0x45, 0x4e, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x05, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x41, 0x4c, 0x4c,
	0x45, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x49, 0x43, 0x41,
	0x4c, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x52,
	0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x09,
	0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x4d, 0x55, 0x54, 0x45, 0x44, 0x10,
	0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0c, 0x42,
	0x14, 0x5a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescOnce sync.Once
	file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescData = file_pkg_rpc_proto_chat_chat_video_call_proto_rawDesc
)

func file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescGZIP() []byte {
	file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescData)
	})
	return file_pkg_rpc_proto_chat_chat_video_call_proto_rawDescData
}

var file_pkg_rpc_proto_chat_chat_video_call_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_rpc_proto_chat_chat_video_call_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_rpc_proto_chat_chat_video_call_proto_goTypes = []interface{}{
	(ChatVideoCall_Status)(0), // 0: pkg.rpc.proto.chat.ChatVideoCall.Status
	(*ChatVideoCall)(nil),     // 1: pkg.rpc.proto.chat.ChatVideoCall
}
var file_pkg_rpc_proto_chat_chat_video_call_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_rpc_proto_chat_chat_video_call_proto_init() }
func file_pkg_rpc_proto_chat_chat_video_call_proto_init() {
	if File_pkg_rpc_proto_chat_chat_video_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_proto_chat_chat_video_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatVideoCall); i {
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
			RawDescriptor: file_pkg_rpc_proto_chat_chat_video_call_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_rpc_proto_chat_chat_video_call_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_proto_chat_chat_video_call_proto_depIdxs,
		EnumInfos:         file_pkg_rpc_proto_chat_chat_video_call_proto_enumTypes,
		MessageInfos:      file_pkg_rpc_proto_chat_chat_video_call_proto_msgTypes,
	}.Build()
	File_pkg_rpc_proto_chat_chat_video_call_proto = out.File
	file_pkg_rpc_proto_chat_chat_video_call_proto_rawDesc = nil
	file_pkg_rpc_proto_chat_chat_video_call_proto_goTypes = nil
	file_pkg_rpc_proto_chat_chat_video_call_proto_depIdxs = nil
}
