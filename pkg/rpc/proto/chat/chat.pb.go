//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-07 23:31:57

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/proto/chat/chat.proto

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

// 状态
type ToReply_Status int32

const (
	ToReply_SENT   ToReply_Status = 0  // 已发送
	ToReply_REJECT ToReply_Status = 22 // 拒收
)

// Enum value maps for ToReply_Status.
var (
	ToReply_Status_name = map[int32]string{
		0:  "SENT",
		22: "REJECT",
	}
	ToReply_Status_value = map[string]int32{
		"SENT":   0,
		"REJECT": 22,
	}
)

func (x ToReply_Status) Enum() *ToReply_Status {
	p := new(ToReply_Status)
	*p = x
	return p
}

func (x ToReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_rpc_proto_chat_chat_proto_enumTypes[0].Descriptor()
}

func (ToReply_Status) Type() protoreflect.EnumType {
	return &file_pkg_rpc_proto_chat_chat_proto_enumTypes[0]
}

func (x ToReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToReply_Status.Descriptor instead.
func (ToReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{3, 0}
}

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
		mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[0]
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
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[1]
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
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (m *Content) GetType() isContent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Content) GetText() *ChatText {
	if x, ok := x.GetType().(*Content_Text); ok {
		return x.Text
	}
	return nil
}

func (x *Content) GetSticker() *ChatSticker {
	if x, ok := x.GetType().(*Content_Sticker); ok {
		return x.Sticker
	}
	return nil
}

func (x *Content) GetQuoted() *ChatQuoted {
	if x, ok := x.GetType().(*Content_Quoted); ok {
		return x.Quoted
	}
	return nil
}

func (x *Content) GetImage() *ChatImage {
	if x, ok := x.GetType().(*Content_Image); ok {
		return x.Image
	}
	return nil
}

func (x *Content) GetAudio() *ChatAudio {
	if x, ok := x.GetType().(*Content_Audio); ok {
		return x.Audio
	}
	return nil
}

func (x *Content) GetVideo() *ChatVideo {
	if x, ok := x.GetType().(*Content_Video); ok {
		return x.Video
	}
	return nil
}

func (x *Content) GetFile() *ChatFile {
	if x, ok := x.GetType().(*Content_File); ok {
		return x.File
	}
	return nil
}

func (x *Content) GetLocation() *ChatLocation {
	if x, ok := x.GetType().(*Content_Location); ok {
		return x.Location
	}
	return nil
}

func (x *Content) GetCard() *ChatCard {
	if x, ok := x.GetType().(*Content_Card); ok {
		return x.Card
	}
	return nil
}

func (x *Content) GetWithdraw() *ChatWithdraw {
	if x, ok := x.GetType().(*Content_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (x *Content) GetVoice() *ChatVoice {
	if x, ok := x.GetType().(*Content_Voice); ok {
		return x.Voice
	}
	return nil
}

func (x *Content) GetVoiceCall() *ChatVoiceCall {
	if x, ok := x.GetType().(*Content_VoiceCall); ok {
		return x.VoiceCall
	}
	return nil
}

func (x *Content) GetVideoCall() *ChatVideoCall {
	if x, ok := x.GetType().(*Content_VideoCall); ok {
		return x.VideoCall
	}
	return nil
}

type isContent_Type interface {
	isContent_Type()
}

type Content_Text struct {
	Text *ChatText `protobuf:"bytes,10,opt,name=text,proto3,oneof"` // 文字
}

type Content_Sticker struct {
	Sticker *ChatSticker `protobuf:"bytes,11,opt,name=sticker,proto3,oneof"` // 表情包
}

type Content_Quoted struct {
	Quoted *ChatQuoted `protobuf:"bytes,12,opt,name=quoted,proto3,oneof"` // 引用
}

type Content_Image struct {
	Image *ChatImage `protobuf:"bytes,13,opt,name=image,proto3,oneof"` // 图片
}

type Content_Audio struct {
	Audio *ChatAudio `protobuf:"bytes,14,opt,name=audio,proto3,oneof"` // 音频
}

type Content_Video struct {
	Video *ChatVideo `protobuf:"bytes,15,opt,name=video,proto3,oneof"` // 视频
}

type Content_File struct {
	File *ChatFile `protobuf:"bytes,16,opt,name=file,proto3,oneof"` // 文件
}

type Content_Location struct {
	Location *ChatLocation `protobuf:"bytes,17,opt,name=location,proto3,oneof"` // 定位
}

type Content_Card struct {
	Card *ChatCard `protobuf:"bytes,18,opt,name=card,proto3,oneof"` // 卡片
}

type Content_Withdraw struct {
	Withdraw *ChatWithdraw `protobuf:"bytes,19,opt,name=withdraw,proto3,oneof"` // 撤回
}

type Content_Voice struct {
	Voice *ChatVoice `protobuf:"bytes,20,opt,name=voice,proto3,oneof"` // 语音信息
}

type Content_VoiceCall struct {
	VoiceCall *ChatVoiceCall `protobuf:"bytes,21,opt,name=voiceCall,proto3,oneof"` // 语音通话
}

type Content_VideoCall struct {
	VideoCall *ChatVideoCall `protobuf:"bytes,22,opt,name=videoCall,proto3,oneof"` // 视频通话
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
type ToRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *Source  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"` // 来源
	Chat   *Content `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`     // 内容
}

func (x *ToRequest) Reset() {
	*x = ToRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToRequest) ProtoMessage() {}

func (x *ToRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToRequest.ProtoReflect.Descriptor instead.
func (*ToRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ToRequest) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ToRequest) GetChat() *Content {
	if x != nil {
		return x.Chat
	}
	return nil
}

// 发送回复
type ToReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceId string         `protobuf:"bytes,1,opt,name=sequenceId,proto3" json:"sequenceId,omitempty"`                                 // 序号
	Status     ToReply_Status `protobuf:"varint,3,opt,name=status,proto3,enum=pkg.rpc.proto.chat.ToReply_Status" json:"status,omitempty"` // 状态
	Comment    string         `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`                                       // 备注
}

func (x *ToReply) Reset() {
	*x = ToReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToReply) ProtoMessage() {}

func (x *ToReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToReply.ProtoReflect.Descriptor instead.
func (*ToReply) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ToReply) GetSequenceId() string {
	if x != nil {
		return x.SequenceId
	}
	return ""
}

func (x *ToReply) GetStatus() ToReply_Status {
	if x != nil {
		return x.Status
	}
	return ToReply_SENT
}

func (x *ToReply) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

// 订阅请求
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[4]
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
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{4}
}

// 订阅回复
type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *Source  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"` // 来源
	Chat   *Content `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`     // 内容
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_proto_msgTypes[5]
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
	return file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeReply) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *SubscribeReply) GetChat() *Content {
	if x != nil {
		return x.Chat
	}
	return nil
}

var File_pkg_rpc_proto_chat_chat_proto protoreflect.FileDescriptor

var file_pkg_rpc_proto_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x0a, 0x02, 0x49,
	0x64, 0x22, 0x86, 0x06, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x3b, 0x0a, 0x07, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x35, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x48, 0x00, 0x52,
	0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x32, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x3e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52,
	0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x00, 0x52, 0x08, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x09,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x12,
	0x41, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61,
	0x6c, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x70, 0x0a, 0x09, 0x54, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x63,
	0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x9f, 0x01, 0x0a,
	0x07, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x54, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1e,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x16, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x75, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x32, 0xa1, 0x01, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x40, 0x0a, 0x02, 0x54, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x54, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x54, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x24, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x14, 0x5a,
	0x12, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_proto_chat_chat_proto_rawDescOnce sync.Once
	file_pkg_rpc_proto_chat_chat_proto_rawDescData = file_pkg_rpc_proto_chat_chat_proto_rawDesc
)

func file_pkg_rpc_proto_chat_chat_proto_rawDescGZIP() []byte {
	file_pkg_rpc_proto_chat_chat_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_proto_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_proto_chat_chat_proto_rawDescData)
	})
	return file_pkg_rpc_proto_chat_chat_proto_rawDescData
}

var file_pkg_rpc_proto_chat_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_rpc_proto_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_rpc_proto_chat_chat_proto_goTypes = []interface{}{
	(ToReply_Status)(0),      // 0: pkg.rpc.proto.chat.ToReply.Status
	(*Source)(nil),           // 1: pkg.rpc.proto.chat.Source
	(*Content)(nil),          // 2: pkg.rpc.proto.chat.Content
	(*ToRequest)(nil),        // 3: pkg.rpc.proto.chat.ToRequest
	(*ToReply)(nil),          // 4: pkg.rpc.proto.chat.ToReply
	(*SubscribeRequest)(nil), // 5: pkg.rpc.proto.chat.SubscribeRequest
	(*SubscribeReply)(nil),   // 6: pkg.rpc.proto.chat.SubscribeReply
	(*ChatText)(nil),         // 7: pkg.rpc.proto.chat.ChatText
	(*ChatSticker)(nil),      // 8: pkg.rpc.proto.chat.ChatSticker
	(*ChatQuoted)(nil),       // 9: pkg.rpc.proto.chat.ChatQuoted
	(*ChatImage)(nil),        // 10: pkg.rpc.proto.chat.ChatImage
	(*ChatAudio)(nil),        // 11: pkg.rpc.proto.chat.ChatAudio
	(*ChatVideo)(nil),        // 12: pkg.rpc.proto.chat.ChatVideo
	(*ChatFile)(nil),         // 13: pkg.rpc.proto.chat.ChatFile
	(*ChatLocation)(nil),     // 14: pkg.rpc.proto.chat.ChatLocation
	(*ChatCard)(nil),         // 15: pkg.rpc.proto.chat.ChatCard
	(*ChatWithdraw)(nil),     // 16: pkg.rpc.proto.chat.ChatWithdraw
	(*ChatVoice)(nil),        // 17: pkg.rpc.proto.chat.ChatVoice
	(*ChatVoiceCall)(nil),    // 18: pkg.rpc.proto.chat.ChatVoiceCall
	(*ChatVideoCall)(nil),    // 19: pkg.rpc.proto.chat.ChatVideoCall
}
var file_pkg_rpc_proto_chat_chat_proto_depIdxs = []int32{
	7,  // 0: pkg.rpc.proto.chat.Content.text:type_name -> pkg.rpc.proto.chat.ChatText
	8,  // 1: pkg.rpc.proto.chat.Content.sticker:type_name -> pkg.rpc.proto.chat.ChatSticker
	9,  // 2: pkg.rpc.proto.chat.Content.quoted:type_name -> pkg.rpc.proto.chat.ChatQuoted
	10, // 3: pkg.rpc.proto.chat.Content.image:type_name -> pkg.rpc.proto.chat.ChatImage
	11, // 4: pkg.rpc.proto.chat.Content.audio:type_name -> pkg.rpc.proto.chat.ChatAudio
	12, // 5: pkg.rpc.proto.chat.Content.video:type_name -> pkg.rpc.proto.chat.ChatVideo
	13, // 6: pkg.rpc.proto.chat.Content.file:type_name -> pkg.rpc.proto.chat.ChatFile
	14, // 7: pkg.rpc.proto.chat.Content.location:type_name -> pkg.rpc.proto.chat.ChatLocation
	15, // 8: pkg.rpc.proto.chat.Content.card:type_name -> pkg.rpc.proto.chat.ChatCard
	16, // 9: pkg.rpc.proto.chat.Content.withdraw:type_name -> pkg.rpc.proto.chat.ChatWithdraw
	17, // 10: pkg.rpc.proto.chat.Content.voice:type_name -> pkg.rpc.proto.chat.ChatVoice
	18, // 11: pkg.rpc.proto.chat.Content.voiceCall:type_name -> pkg.rpc.proto.chat.ChatVoiceCall
	19, // 12: pkg.rpc.proto.chat.Content.videoCall:type_name -> pkg.rpc.proto.chat.ChatVideoCall
	1,  // 13: pkg.rpc.proto.chat.ToRequest.source:type_name -> pkg.rpc.proto.chat.Source
	2,  // 14: pkg.rpc.proto.chat.ToRequest.chat:type_name -> pkg.rpc.proto.chat.Content
	0,  // 15: pkg.rpc.proto.chat.ToReply.status:type_name -> pkg.rpc.proto.chat.ToReply.Status
	1,  // 16: pkg.rpc.proto.chat.SubscribeReply.source:type_name -> pkg.rpc.proto.chat.Source
	2,  // 17: pkg.rpc.proto.chat.SubscribeReply.chat:type_name -> pkg.rpc.proto.chat.Content
	3,  // 18: pkg.rpc.proto.chat.Chat.To:input_type -> pkg.rpc.proto.chat.ToRequest
	5,  // 19: pkg.rpc.proto.chat.Chat.Subscribe:input_type -> pkg.rpc.proto.chat.SubscribeRequest
	4,  // 20: pkg.rpc.proto.chat.Chat.To:output_type -> pkg.rpc.proto.chat.ToReply
	6,  // 21: pkg.rpc.proto.chat.Chat.Subscribe:output_type -> pkg.rpc.proto.chat.SubscribeReply
	20, // [20:22] is the sub-list for method output_type
	18, // [18:20] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_pkg_rpc_proto_chat_chat_proto_init() }
func file_pkg_rpc_proto_chat_chat_proto_init() {
	if File_pkg_rpc_proto_chat_chat_proto != nil {
		return
	}
	file_pkg_rpc_proto_chat_chat_text_proto_init()
	file_pkg_rpc_proto_chat_chat_sticker_proto_init()
	file_pkg_rpc_proto_chat_chat_quoted_proto_init()
	file_pkg_rpc_proto_chat_chat_image_proto_init()
	file_pkg_rpc_proto_chat_chat_audio_proto_init()
	file_pkg_rpc_proto_chat_chat_video_proto_init()
	file_pkg_rpc_proto_chat_chat_file_proto_init()
	file_pkg_rpc_proto_chat_chat_location_proto_init()
	file_pkg_rpc_proto_chat_chat_card_proto_init()
	file_pkg_rpc_proto_chat_chat_withdraw_proto_init()
	file_pkg_rpc_proto_chat_chat_voice_proto_init()
	file_pkg_rpc_proto_chat_chat_voice_call_proto_init()
	file_pkg_rpc_proto_chat_chat_video_call_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_proto_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_rpc_proto_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_rpc_proto_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToRequest); i {
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
		file_pkg_rpc_proto_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToReply); i {
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
		file_pkg_rpc_proto_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_rpc_proto_chat_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	file_pkg_rpc_proto_chat_chat_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Source_UserId)(nil),
		(*Source_GroupId)(nil),
		(*Source_RoomId)(nil),
		(*Source_BottleId)(nil),
	}
	file_pkg_rpc_proto_chat_chat_proto_msgTypes[1].OneofWrappers = []interface{}{
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
			RawDescriptor: file_pkg_rpc_proto_chat_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rpc_proto_chat_chat_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_proto_chat_chat_proto_depIdxs,
		EnumInfos:         file_pkg_rpc_proto_chat_chat_proto_enumTypes,
		MessageInfos:      file_pkg_rpc_proto_chat_chat_proto_msgTypes,
	}.Build()
	File_pkg_rpc_proto_chat_chat_proto = out.File
	file_pkg_rpc_proto_chat_chat_proto_rawDesc = nil
	file_pkg_rpc_proto_chat_chat_proto_goTypes = nil
	file_pkg_rpc_proto_chat_chat_proto_depIdxs = nil
}
