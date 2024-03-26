//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 10:34:41

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internal/rpc/message/content/video.proto

package content

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

// 视频
type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Uri:
	//
	//	*Video_AliasName
	//	*Video_Hash
	//	*Video_Url
	Uri isVideo_Uri `protobuf_oneof:"Uri"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_content_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_content_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_content_video_proto_rawDescGZIP(), []int{0}
}

func (m *Video) GetUri() isVideo_Uri {
	if m != nil {
		return m.Uri
	}
	return nil
}

func (x *Video) GetAliasName() string {
	if x, ok := x.GetUri().(*Video_AliasName); ok {
		return x.AliasName
	}
	return ""
}

func (x *Video) GetHash() string {
	if x, ok := x.GetUri().(*Video_Hash); ok {
		return x.Hash
	}
	return ""
}

func (x *Video) GetUrl() string {
	if x, ok := x.GetUri().(*Video_Url); ok {
		return x.Url
	}
	return ""
}

type isVideo_Uri interface {
	isVideo_Uri()
}

type Video_AliasName struct {
	AliasName string `protobuf:"bytes,1,opt,name=alias_name,json=aliasName,proto3,oneof"` // 别名
}

type Video_Hash struct {
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3,oneof"` // 哈希
}

type Video_Url struct {
	Url string `protobuf:"bytes,3,opt,name=url,proto3,oneof"` // 地址
}

func (*Video_AliasName) isVideo_Uri() {}

func (*Video_Hash) isVideo_Uri() {}

func (*Video_Url) isVideo_Uri() {}

var File_internal_rpc_message_content_video_proto protoreflect.FileDescriptor

var file_internal_rpc_message_content_video_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x05, 0x0a, 0x03,
	0x55, 0x72, 0x69, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_message_content_video_proto_rawDescOnce sync.Once
	file_internal_rpc_message_content_video_proto_rawDescData = file_internal_rpc_message_content_video_proto_rawDesc
)

func file_internal_rpc_message_content_video_proto_rawDescGZIP() []byte {
	file_internal_rpc_message_content_video_proto_rawDescOnce.Do(func() {
		file_internal_rpc_message_content_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_message_content_video_proto_rawDescData)
	})
	return file_internal_rpc_message_content_video_proto_rawDescData
}

var file_internal_rpc_message_content_video_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_rpc_message_content_video_proto_goTypes = []interface{}{
	(*Video)(nil), // 0: internal.rpc.message.content.Video
}
var file_internal_rpc_message_content_video_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_rpc_message_content_video_proto_init() }
func file_internal_rpc_message_content_video_proto_init() {
	if File_internal_rpc_message_content_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_rpc_message_content_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
	file_internal_rpc_message_content_video_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Video_AliasName)(nil),
		(*Video_Hash)(nil),
		(*Video_Url)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_rpc_message_content_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_rpc_message_content_video_proto_goTypes,
		DependencyIndexes: file_internal_rpc_message_content_video_proto_depIdxs,
		MessageInfos:      file_internal_rpc_message_content_video_proto_msgTypes,
	}.Build()
	File_internal_rpc_message_content_video_proto = out.File
	file_internal_rpc_message_content_video_proto_rawDesc = nil
	file_internal_rpc_message_content_video_proto_goTypes = nil
	file_internal_rpc_message_content_video_proto_depIdxs = nil
}
