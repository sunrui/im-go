//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-07 23:31:57

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internal/rpc/message/chat/chat.proto

package chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	content "internal/rpc/message/content"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 状态
type SendReply_Status int32

const (
	SendReply_SENT   SendReply_Status = 0  // 已发送
	SendReply_REJECT SendReply_Status = 22 // 拒收
)

// Enum value maps for SendReply_Status.
var (
	SendReply_Status_name = map[int32]string{
		0:  "SENT",
		22: "REJECT",
	}
	SendReply_Status_value = map[string]int32{
		"SENT":   0,
		"REJECT": 22,
	}
)

func (x SendReply_Status) Enum() *SendReply_Status {
	p := new(SendReply_Status)
	*p = x
	return p
}

func (x SendReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_rpc_message_chat_chat_proto_enumTypes[0].Descriptor()
}

func (SendReply_Status) Type() protoreflect.EnumType {
	return &file_internal_rpc_message_chat_chat_proto_enumTypes[0]
}

func (x SendReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendReply_Status.Descriptor instead.
func (SendReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_chat_proto_rawDescGZIP(), []int{2, 0}
}

// 聊天内容
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Content_Text
	//	*Content_Sticker
	//	*Content_Quoted
	//	*Content_Image
	//	*Content_Audio
	//	*Content_Video
	//	*Content_File
	//	*Content_Location
	//	*Content_Card
	//	*Content_Withdraw
	//	*Content_Voice
	//	*Content_VoiceCall
	//	*Content_VideoCall
	Type isContent_Type `protobuf_oneof:"Type"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (m *Content) GetType() isContent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Content) GetText() *content.Text {
	if x, ok := x.GetType().(*Content_Text); ok {
		return x.Text
	}
	return nil
}

func (x *Content) GetSticker() *content.Sticker {
	if x, ok := x.GetType().(*Content_Sticker); ok {
		return x.Sticker
	}
	return nil
}

func (x *Content) GetQuoted() *content.Quoted {
	if x, ok := x.GetType().(*Content_Quoted); ok {
		return x.Quoted
	}
	return nil
}

func (x *Content) GetImage() *content.Image {
	if x, ok := x.GetType().(*Content_Image); ok {
		return x.Image
	}
	return nil
}

func (x *Content) GetAudio() *content.Audio {
	if x, ok := x.GetType().(*Content_Audio); ok {
		return x.Audio
	}
	return nil
}

func (x *Content) GetVideo() *content.Video {
	if x, ok := x.GetType().(*Content_Video); ok {
		return x.Video
	}
	return nil
}

func (x *Content) GetFile() *content.File {
	if x, ok := x.GetType().(*Content_File); ok {
		return x.File
	}
	return nil
}

func (x *Content) GetLocation() *content.Location {
	if x, ok := x.GetType().(*Content_Location); ok {
		return x.Location
	}
	return nil
}

func (x *Content) GetCard() *content.Card {
	if x, ok := x.GetType().(*Content_Card); ok {
		return x.Card
	}
	return nil
}

func (x *Content) GetWithdraw() *content.Withdraw {
	if x, ok := x.GetType().(*Content_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (x *Content) GetVoice() *content.Voice {
	if x, ok := x.GetType().(*Content_Voice); ok {
		return x.Voice
	}
	return nil
}

func (x *Content) GetVoiceCall() *content.VoiceCall {
	if x, ok := x.GetType().(*Content_VoiceCall); ok {
		return x.VoiceCall
	}
	return nil
}

func (x *Content) GetVideoCall() *content.VideoCall {
	if x, ok := x.GetType().(*Content_VideoCall); ok {
		return x.VideoCall
	}
	return nil
}

type isContent_Type interface {
	isContent_Type()
}

type Content_Text struct {
	Text *content.Text `protobuf:"bytes,10,opt,name=text,proto3,oneof"` // 文字
}

type Content_Sticker struct {
	Sticker *content.Sticker `protobuf:"bytes,11,opt,name=sticker,proto3,oneof"` // 表情包
}

type Content_Quoted struct {
	Quoted *content.Quoted `protobuf:"bytes,12,opt,name=quoted,proto3,oneof"` // 引用
}

type Content_Image struct {
	Image *content.Image `protobuf:"bytes,13,opt,name=image,proto3,oneof"` // 图片
}

type Content_Audio struct {
	Audio *content.Audio `protobuf:"bytes,14,opt,name=audio,proto3,oneof"` // 音频
}

type Content_Video struct {
	Video *content.Video `protobuf:"bytes,15,opt,name=video,proto3,oneof"` // 视频
}

type Content_File struct {
	File *content.File `protobuf:"bytes,16,opt,name=file,proto3,oneof"` // 文件
}

type Content_Location struct {
	Location *content.Location `protobuf:"bytes,17,opt,name=location,proto3,oneof"` // 定位
}

type Content_Card struct {
	Card *content.Card `protobuf:"bytes,18,opt,name=card,proto3,oneof"` // 卡片
}

type Content_Withdraw struct {
	Withdraw *content.Withdraw `protobuf:"bytes,19,opt,name=withdraw,proto3,oneof"` // 撤回
}

type Content_Voice struct {
	Voice *content.Voice `protobuf:"bytes,20,opt,name=voice,proto3,oneof"` // 语音信息
}

type Content_VoiceCall struct {
	VoiceCall *content.VoiceCall `protobuf:"bytes,21,opt,name=voiceCall,proto3,oneof"` // 语音通话
}

type Content_VideoCall struct {
	VideoCall *content.VideoCall `protobuf:"bytes,22,opt,name=videoCall,proto3,oneof"` // 视频通话
}

func (*Content_Text) isContent_Type() {}

func (*Content_Sticker) isContent_Type() {}

func (*Content_Quoted) isContent_Type() {}

func (*Content_Image) isContent_Type() {}

func (*Content_Audio) isContent_Type() {}

func (*Content_Video) isContent_Type() {}

func (*Content_File) isContent_Type() {}

func (*Content_Location) isContent_Type() {}

func (*Content_Card) isContent_Type() {}

func (*Content_Withdraw) isContent_Type() {}

func (*Content_Voice) isContent_Type() {}

func (*Content_VoiceCall) isContent_Type() {}

func (*Content_VideoCall) isContent_Type() {}

// 发送请求
type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  *Source  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`   // 来源
	Content *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 内容
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendRequest) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *SendRequest) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

// 发送回复
type SendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceId int32            `protobuf:"varint,1,opt,name=sequenceId,proto3" json:"sequenceId,omitempty"`                                         // 序号
	Status     SendReply_Status `protobuf:"varint,3,opt,name=status,proto3,enum=internal.rpc.message.chat.SendReply_Status" json:"status,omitempty"` // 状态
	Comment    string           `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`                                                // 备注
}

func (x *SendReply) Reset() {
	*x = SendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReply) ProtoMessage() {}

func (x *SendReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReply.ProtoReflect.Descriptor instead.
func (*SendReply) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *SendReply) GetSequenceId() int32 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *SendReply) GetStatus() SendReply_Status {
	if x != nil {
		return x.Status
	}
	return SendReply_SENT
}

func (x *SendReply) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

// 接收请求
type ReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveRequest) Reset() {
	*x = ReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveRequest) ProtoMessage() {}

func (x *ReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveRequest.ProtoReflect.Descriptor instead.
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_chat_proto_rawDescGZIP(), []int{3}
}

// 接收回复
type ReceiveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  *Source  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`   // 来源
	Content *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 内容
}

func (x *ReceiveReply) Reset() {
	*x = ReceiveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveReply) ProtoMessage() {}

func (x *ReceiveReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_message_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveReply.ProtoReflect.Descriptor instead.
func (*ReceiveReply) Descriptor() ([]byte, []int) {
	return file_internal_rpc_message_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ReceiveReply) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ReceiveReply) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_internal_rpc_message_chat_chat_proto protoreflect.FileDescriptor

var file_internal_rpc_message_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x1a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x06, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x41, 0x0a, 0x07, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x06, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x3b,
	0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x38, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52,
	0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48,
	0x00, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x3b, 0x0a, 0x05, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x47, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x00, 0x52,
	0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x09,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x16, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x32, 0xbd, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x54, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x29,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x30, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_message_chat_chat_proto_rawDescOnce sync.Once
	file_internal_rpc_message_chat_chat_proto_rawDescData = file_internal_rpc_message_chat_chat_proto_rawDesc
)

func file_internal_rpc_message_chat_chat_proto_rawDescGZIP() []byte {
	file_internal_rpc_message_chat_chat_proto_rawDescOnce.Do(func() {
		file_internal_rpc_message_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_message_chat_chat_proto_rawDescData)
	})
	return file_internal_rpc_message_chat_chat_proto_rawDescData
}

var file_internal_rpc_message_chat_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_rpc_message_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_rpc_message_chat_chat_proto_goTypes = []interface{}{
	(SendReply_Status)(0),     // 0: internal.rpc.message.chat.SendReply.Status
	(*Content)(nil),           // 1: internal.rpc.message.chat.Content
	(*SendRequest)(nil),       // 2: internal.rpc.message.chat.SendRequest
	(*SendReply)(nil),         // 3: internal.rpc.message.chat.SendReply
	(*ReceiveRequest)(nil),    // 4: internal.rpc.message.chat.ReceiveRequest
	(*ReceiveReply)(nil),      // 5: internal.rpc.message.chat.ReceiveReply
	(*content.Text)(nil),      // 6: internal.rpc.message.content.Text
	(*content.Sticker)(nil),   // 7: internal.rpc.message.content.Sticker
	(*content.Quoted)(nil),    // 8: internal.rpc.message.content.Quoted
	(*content.Image)(nil),     // 9: internal.rpc.message.content.Image
	(*content.Audio)(nil),     // 10: internal.rpc.message.content.Audio
	(*content.Video)(nil),     // 11: internal.rpc.message.content.Video
	(*content.File)(nil),      // 12: internal.rpc.message.content.File
	(*content.Location)(nil),  // 13: internal.rpc.message.content.Location
	(*content.Card)(nil),      // 14: internal.rpc.message.content.Card
	(*content.Withdraw)(nil),  // 15: internal.rpc.message.content.Withdraw
	(*content.Voice)(nil),     // 16: internal.rpc.message.content.Voice
	(*content.VoiceCall)(nil), // 17: internal.rpc.message.content.VoiceCall
	(*content.VideoCall)(nil), // 18: internal.rpc.message.content.VideoCall
	(*Source)(nil),            // 19: internal.rpc.message.chat.Source
}
var file_internal_rpc_message_chat_chat_proto_depIdxs = []int32{
	6,  // 0: internal.rpc.message.chat.Content.text:type_name -> internal.rpc.message.content.Text
	7,  // 1: internal.rpc.message.chat.Content.sticker:type_name -> internal.rpc.message.content.Sticker
	8,  // 2: internal.rpc.message.chat.Content.quoted:type_name -> internal.rpc.message.content.Quoted
	9,  // 3: internal.rpc.message.chat.Content.image:type_name -> internal.rpc.message.content.Image
	10, // 4: internal.rpc.message.chat.Content.audio:type_name -> internal.rpc.message.content.Audio
	11, // 5: internal.rpc.message.chat.Content.video:type_name -> internal.rpc.message.content.Video
	12, // 6: internal.rpc.message.chat.Content.file:type_name -> internal.rpc.message.content.File
	13, // 7: internal.rpc.message.chat.Content.location:type_name -> internal.rpc.message.content.Location
	14, // 8: internal.rpc.message.chat.Content.card:type_name -> internal.rpc.message.content.Card
	15, // 9: internal.rpc.message.chat.Content.withdraw:type_name -> internal.rpc.message.content.Withdraw
	16, // 10: internal.rpc.message.chat.Content.voice:type_name -> internal.rpc.message.content.Voice
	17, // 11: internal.rpc.message.chat.Content.voiceCall:type_name -> internal.rpc.message.content.VoiceCall
	18, // 12: internal.rpc.message.chat.Content.videoCall:type_name -> internal.rpc.message.content.VideoCall
	19, // 13: internal.rpc.message.chat.SendRequest.source:type_name -> internal.rpc.message.chat.Source
	1,  // 14: internal.rpc.message.chat.SendRequest.content:type_name -> internal.rpc.message.chat.Content
	0,  // 15: internal.rpc.message.chat.SendReply.status:type_name -> internal.rpc.message.chat.SendReply.Status
	19, // 16: internal.rpc.message.chat.ReceiveReply.source:type_name -> internal.rpc.message.chat.Source
	1,  // 17: internal.rpc.message.chat.ReceiveReply.content:type_name -> internal.rpc.message.chat.Content
	2,  // 18: internal.rpc.message.chat.Chat.Send:input_type -> internal.rpc.message.chat.SendRequest
	4,  // 19: internal.rpc.message.chat.Chat.Receive:input_type -> internal.rpc.message.chat.ReceiveRequest
	3,  // 20: internal.rpc.message.chat.Chat.Send:output_type -> internal.rpc.message.chat.SendReply
	5,  // 21: internal.rpc.message.chat.Chat.Receive:output_type -> internal.rpc.message.chat.ReceiveReply
	20, // [20:22] is the sub-list for method output_type
	18, // [18:20] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_internal_rpc_message_chat_chat_proto_init() }
func file_internal_rpc_message_chat_chat_proto_init() {
	if File_internal_rpc_message_chat_chat_proto != nil {
		return
	}
	file_internal_rpc_message_chat_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_rpc_message_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_internal_rpc_message_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_internal_rpc_message_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReply); i {
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
		file_internal_rpc_message_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveRequest); i {
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
		file_internal_rpc_message_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveReply); i {
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
	file_internal_rpc_message_chat_chat_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Content_Text)(nil),
		(*Content_Sticker)(nil),
		(*Content_Quoted)(nil),
		(*Content_Image)(nil),
		(*Content_Audio)(nil),
		(*Content_Video)(nil),
		(*Content_File)(nil),
		(*Content_Location)(nil),
		(*Content_Card)(nil),
		(*Content_Withdraw)(nil),
		(*Content_Voice)(nil),
		(*Content_VoiceCall)(nil),
		(*Content_VideoCall)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_rpc_message_chat_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_rpc_message_chat_chat_proto_goTypes,
		DependencyIndexes: file_internal_rpc_message_chat_chat_proto_depIdxs,
		EnumInfos:         file_internal_rpc_message_chat_chat_proto_enumTypes,
		MessageInfos:      file_internal_rpc_message_chat_chat_proto_msgTypes,
	}.Build()
	File_internal_rpc_message_chat_chat_proto = out.File
	file_internal_rpc_message_chat_chat_proto_rawDesc = nil
	file_internal_rpc_message_chat_chat_proto_goTypes = nil
	file_internal_rpc_message_chat_chat_proto_depIdxs = nil
}
