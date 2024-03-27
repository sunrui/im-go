//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 10:57:20

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/proto/message/message_video_call.proto

package message

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

// 视频通话
type MessageVideoCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalleeId               string `protobuf:"bytes,1,opt,name=callee_id,json=calleeId,proto3" json:"callee_id,omitempty"`                                                // 对方的唯一标识，如用户ID或电话号码
	StartTimeUnixTimestamp int64  `protobuf:"varint,2,opt,name=start_time_unix_timestamp,json=startTimeUnixTimestamp,proto3" json:"start_time_unix_timestamp,omitempty"` // 通话开始时间戳（Unix时间戳）
	DurationSeconds        int32  `protobuf:"varint,3,opt,name=duration_seconds,json=durationSeconds,proto3" json:"duration_seconds,omitempty"`                          // 通话持续时长（秒）
	Protocol               string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`                                                                // 通话使用的协议类型（例如：SIP、WebRTC等）
	SessionUrlOrToken      string `protobuf:"bytes,5,opt,name=session_url_or_token,json=sessionUrlOrToken,proto3" json:"session_url_or_token,omitempty"`                 // （可选）如果是VoIP服务，可能是用于建立通话会话的服务器URL或令牌
}

func (x *MessageVideoCall) Reset() {
	*x = MessageVideoCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_message_message_video_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageVideoCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageVideoCall) ProtoMessage() {}

func (x *MessageVideoCall) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_message_message_video_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageVideoCall.ProtoReflect.Descriptor instead.
func (*MessageVideoCall) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_message_message_video_call_proto_rawDescGZIP(), []int{0}
}

func (x *MessageVideoCall) GetCalleeId() string {
	if x != nil {
		return x.CalleeId
	}
	return ""
}

func (x *MessageVideoCall) GetStartTimeUnixTimestamp() int64 {
	if x != nil {
		return x.StartTimeUnixTimestamp
	}
	return 0
}

func (x *MessageVideoCall) GetDurationSeconds() int32 {
	if x != nil {
		return x.DurationSeconds
	}
	return 0
}

func (x *MessageVideoCall) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *MessageVideoCall) GetSessionUrlOrToken() string {
	if x != nil {
		return x.SessionUrlOrToken
	}
	return ""
}

var File_pkg_rpc_proto_message_message_video_call_proto protoreflect.FileDescriptor

var file_pkg_rpc_proto_message_message_video_call_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x2f, 0x0a, 0x14, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6f, 0x72, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x4f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x17, 0x5a, 0x15,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_proto_message_message_video_call_proto_rawDescOnce sync.Once
	file_pkg_rpc_proto_message_message_video_call_proto_rawDescData = file_pkg_rpc_proto_message_message_video_call_proto_rawDesc
)

func file_pkg_rpc_proto_message_message_video_call_proto_rawDescGZIP() []byte {
	file_pkg_rpc_proto_message_message_video_call_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_proto_message_message_video_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_proto_message_message_video_call_proto_rawDescData)
	})
	return file_pkg_rpc_proto_message_message_video_call_proto_rawDescData
}

var file_pkg_rpc_proto_message_message_video_call_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_rpc_proto_message_message_video_call_proto_goTypes = []interface{}{
	(*MessageVideoCall)(nil), // 0: pkg.rpc.proto.message.MessageVideoCall
}
var file_pkg_rpc_proto_message_message_video_call_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_rpc_proto_message_message_video_call_proto_init() }
func file_pkg_rpc_proto_message_message_video_call_proto_init() {
	if File_pkg_rpc_proto_message_message_video_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_proto_message_message_video_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageVideoCall); i {
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
			RawDescriptor: file_pkg_rpc_proto_message_message_video_call_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_rpc_proto_message_message_video_call_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_proto_message_message_video_call_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_proto_message_message_video_call_proto_msgTypes,
	}.Build()
	File_pkg_rpc_proto_message_message_video_call_proto = out.File
	file_pkg_rpc_proto_message_message_video_call_proto_rawDesc = nil
	file_pkg_rpc_proto_message_message_video_call_proto_goTypes = nil
	file_pkg_rpc_proto_message_message_video_call_proto_depIdxs = nil
}
