//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 11:52:02

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/proto/p2p_chat/p2p_chat.proto

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

// 订阅请求
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[0]
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
	return file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescGZIP(), []int{0}
}

// 订阅回复
type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"` // 用户 id
	// Types that are assignable to Type:
	//
	//	*SubscribeReply_Online
	//	*SubscribeReply_Offline
	//	*SubscribeReply_Inputting
	//	*SubscribeReply_InputStop
	//	*SubscribeReply_ChatDelivered
	//	*SubscribeReply_ChatRead
	//	*SubscribeReply_ChatReject
	//	*SubscribeReply_ChatWithdraw
	//	*SubscribeReply_FriendRequest
	//	*SubscribeReply_FriendReply
	Type isSubscribeReply_Type `protobuf_oneof:"Type"`
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[1]
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
	return file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeReply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (m *SubscribeReply) GetType() isSubscribeReply_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *SubscribeReply) GetOnline() *Online {
	if x, ok := x.GetType().(*SubscribeReply_Online); ok {
		return x.Online
	}
	return nil
}

func (x *SubscribeReply) GetOffline() *Offline {
	if x, ok := x.GetType().(*SubscribeReply_Offline); ok {
		return x.Offline
	}
	return nil
}

func (x *SubscribeReply) GetInputting() *Inputting {
	if x, ok := x.GetType().(*SubscribeReply_Inputting); ok {
		return x.Inputting
	}
	return nil
}

func (x *SubscribeReply) GetInputStop() *InputStop {
	if x, ok := x.GetType().(*SubscribeReply_InputStop); ok {
		return x.InputStop
	}
	return nil
}

func (x *SubscribeReply) GetChatDelivered() *ChatDelivered {
	if x, ok := x.GetType().(*SubscribeReply_ChatDelivered); ok {
		return x.ChatDelivered
	}
	return nil
}

func (x *SubscribeReply) GetChatRead() *ChatRead {
	if x, ok := x.GetType().(*SubscribeReply_ChatRead); ok {
		return x.ChatRead
	}
	return nil
}

func (x *SubscribeReply) GetChatReject() *ChatReject {
	if x, ok := x.GetType().(*SubscribeReply_ChatReject); ok {
		return x.ChatReject
	}
	return nil
}

func (x *SubscribeReply) GetChatWithdraw() *ChatWithdraw {
	if x, ok := x.GetType().(*SubscribeReply_ChatWithdraw); ok {
		return x.ChatWithdraw
	}
	return nil
}

func (x *SubscribeReply) GetFriendRequest() *FriendRequest {
	if x, ok := x.GetType().(*SubscribeReply_FriendRequest); ok {
		return x.FriendRequest
	}
	return nil
}

func (x *SubscribeReply) GetFriendReply() *FriendReply {
	if x, ok := x.GetType().(*SubscribeReply_FriendReply); ok {
		return x.FriendReply
	}
	return nil
}

type isSubscribeReply_Type interface {
	isSubscribeReply_Type()
}

type SubscribeReply_Online struct {
	Online *Online `protobuf:"bytes,10,opt,name=online,proto3,oneof"` // 上线
}

type SubscribeReply_Offline struct {
	Offline *Offline `protobuf:"bytes,11,opt,name=offline,proto3,oneof"` // 下线
}

type SubscribeReply_Inputting struct {
	Inputting *Inputting `protobuf:"bytes,4,opt,name=inputting,proto3,oneof"` // 输入中
}

type SubscribeReply_InputStop struct {
	InputStop *InputStop `protobuf:"bytes,5,opt,name=inputStop,proto3,oneof"` // 输入停止
}

type SubscribeReply_ChatDelivered struct {
	ChatDelivered *ChatDelivered `protobuf:"bytes,6,opt,name=chatDelivered,proto3,oneof"` // 消息已送达
}

type SubscribeReply_ChatRead struct {
	ChatRead *ChatRead `protobuf:"bytes,7,opt,name=chatRead,proto3,oneof"` // 消息已读
}

type SubscribeReply_ChatReject struct {
	ChatReject *ChatReject `protobuf:"bytes,8,opt,name=chatReject,proto3,oneof"` // 消息拒绝
}

type SubscribeReply_ChatWithdraw struct {
	ChatWithdraw *ChatWithdraw `protobuf:"bytes,9,opt,name=chatWithdraw,proto3,oneof"` // 消息撤回
}

type SubscribeReply_FriendRequest struct {
	FriendRequest *FriendRequest `protobuf:"bytes,2,opt,name=friendRequest,proto3,oneof"` // 好友请求
}

type SubscribeReply_FriendReply struct {
	FriendReply *FriendReply `protobuf:"bytes,3,opt,name=friendReply,proto3,oneof"` // 好友回复
}

func (*SubscribeReply_Online) isSubscribeReply_Type() {}

func (*SubscribeReply_Offline) isSubscribeReply_Type() {}

func (*SubscribeReply_Inputting) isSubscribeReply_Type() {}

func (*SubscribeReply_InputStop) isSubscribeReply_Type() {}

func (*SubscribeReply_ChatDelivered) isSubscribeReply_Type() {}

func (*SubscribeReply_ChatRead) isSubscribeReply_Type() {}

func (*SubscribeReply_ChatReject) isSubscribeReply_Type() {}

func (*SubscribeReply_ChatWithdraw) isSubscribeReply_Type() {}

func (*SubscribeReply_FriendRequest) isSubscribeReply_Type() {}

func (*SubscribeReply_FriendReply) isSubscribeReply_Type() {}

var File_pkg_rpc_proto_p2p_chat_p2p_chat_proto protoreflect.FileDescriptor

var file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x1a,
	0x23, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x73, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x6b, 0x67,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xe6, 0x05, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a,
	0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32,
	0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52,
	0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x53, 0x74, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x48, 0x00, 0x52,
	0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x4d, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x74,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x61, 0x64, 0x48, 0x00, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x61, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32,
	0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x4a, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x00, 0x52, 0x0c, 0x63,
	0x68, 0x61, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x4d, 0x0a, 0x0d, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x32, 0x6a, 0x0a, 0x07, 0x50,
	0x32, 0x70, 0x43, 0x68, 0x61, 0x74, 0x12, 0x5f, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x28, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x32,
	0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescOnce sync.Once
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescData = file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDesc
)

func file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescGZIP() []byte {
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescData)
	})
	return file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDescData
}

var file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil), // 0: pkg.rpc.proto.p2p_chat.SubscribeRequest
	(*SubscribeReply)(nil),   // 1: pkg.rpc.proto.p2p_chat.SubscribeReply
	(*Online)(nil),           // 2: pkg.rpc.proto.p2p_chat.Online
	(*Offline)(nil),          // 3: pkg.rpc.proto.p2p_chat.Offline
	(*Inputting)(nil),        // 4: pkg.rpc.proto.p2p_chat.Inputting
	(*InputStop)(nil),        // 5: pkg.rpc.proto.p2p_chat.InputStop
	(*ChatDelivered)(nil),    // 6: pkg.rpc.proto.p2p_chat.ChatDelivered
	(*ChatRead)(nil),         // 7: pkg.rpc.proto.p2p_chat.ChatRead
	(*ChatReject)(nil),       // 8: pkg.rpc.proto.p2p_chat.ChatReject
	(*ChatWithdraw)(nil),     // 9: pkg.rpc.proto.p2p_chat.ChatWithdraw
	(*FriendRequest)(nil),    // 10: pkg.rpc.proto.p2p_chat.FriendRequest
	(*FriendReply)(nil),      // 11: pkg.rpc.proto.p2p_chat.FriendReply
}
var file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_depIdxs = []int32{
	2,  // 0: pkg.rpc.proto.p2p_chat.SubscribeReply.online:type_name -> pkg.rpc.proto.p2p_chat.Online
	3,  // 1: pkg.rpc.proto.p2p_chat.SubscribeReply.offline:type_name -> pkg.rpc.proto.p2p_chat.Offline
	4,  // 2: pkg.rpc.proto.p2p_chat.SubscribeReply.inputting:type_name -> pkg.rpc.proto.p2p_chat.Inputting
	5,  // 3: pkg.rpc.proto.p2p_chat.SubscribeReply.inputStop:type_name -> pkg.rpc.proto.p2p_chat.InputStop
	6,  // 4: pkg.rpc.proto.p2p_chat.SubscribeReply.chatDelivered:type_name -> pkg.rpc.proto.p2p_chat.ChatDelivered
	7,  // 5: pkg.rpc.proto.p2p_chat.SubscribeReply.chatRead:type_name -> pkg.rpc.proto.p2p_chat.ChatRead
	8,  // 6: pkg.rpc.proto.p2p_chat.SubscribeReply.chatReject:type_name -> pkg.rpc.proto.p2p_chat.ChatReject
	9,  // 7: pkg.rpc.proto.p2p_chat.SubscribeReply.chatWithdraw:type_name -> pkg.rpc.proto.p2p_chat.ChatWithdraw
	10, // 8: pkg.rpc.proto.p2p_chat.SubscribeReply.friendRequest:type_name -> pkg.rpc.proto.p2p_chat.FriendRequest
	11, // 9: pkg.rpc.proto.p2p_chat.SubscribeReply.friendReply:type_name -> pkg.rpc.proto.p2p_chat.FriendReply
	0,  // 10: pkg.rpc.proto.p2p_chat.P2pChat.Subscribe:input_type -> pkg.rpc.proto.p2p_chat.SubscribeRequest
	1,  // 11: pkg.rpc.proto.p2p_chat.P2pChat.Subscribe:output_type -> pkg.rpc.proto.p2p_chat.SubscribeReply
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_init() }
func file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_init() {
	if File_pkg_rpc_proto_p2p_chat_p2p_chat_proto != nil {
		return
	}
	file_pkg_rpc_proto_p2p_chat_online_proto_init()
	file_pkg_rpc_proto_p2p_chat_offline_proto_init()
	file_pkg_rpc_proto_p2p_chat_inputting_proto_init()
	file_pkg_rpc_proto_p2p_chat_input_stop_proto_init()
	file_pkg_rpc_proto_p2p_chat_chat_delivered_proto_init()
	file_pkg_rpc_proto_p2p_chat_chat_read_proto_init()
	file_pkg_rpc_proto_p2p_chat_chat_reject_proto_init()
	file_pkg_rpc_proto_p2p_chat_chat_withdraw_proto_init()
	file_pkg_rpc_proto_p2p_chat_friend_request_proto_init()
	file_pkg_rpc_proto_p2p_chat_friend_reply_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SubscribeReply_Online)(nil),
		(*SubscribeReply_Offline)(nil),
		(*SubscribeReply_Inputting)(nil),
		(*SubscribeReply_InputStop)(nil),
		(*SubscribeReply_ChatDelivered)(nil),
		(*SubscribeReply_ChatRead)(nil),
		(*SubscribeReply_ChatReject)(nil),
		(*SubscribeReply_ChatWithdraw)(nil),
		(*SubscribeReply_FriendRequest)(nil),
		(*SubscribeReply_FriendReply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_msgTypes,
	}.Build()
	File_pkg_rpc_proto_p2p_chat_p2p_chat_proto = out.File
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_rawDesc = nil
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_goTypes = nil
	file_pkg_rpc_proto_p2p_chat_p2p_chat_proto_depIdxs = nil
}
