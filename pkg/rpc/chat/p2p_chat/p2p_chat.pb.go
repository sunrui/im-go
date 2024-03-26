//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 11:52:02

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/chat/p2p_chat/p2p_chat.proto

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
		mi := &file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[0]
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
	return file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescGZIP(), []int{0}
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
	//	*SubscribeReply_MessageDelivered
	//	*SubscribeReply_MessageRead
	//	*SubscribeReply_MessageReject
	//	*SubscribeReply_MessageWithdraw
	//	*SubscribeReply_FriendRequest
	//	*SubscribeReply_FriendReply
	Type isSubscribeReply_Type `protobuf_oneof:"Type"`
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[1]
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
	return file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescGZIP(), []int{1}
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

func (x *SubscribeReply) GetMessageDelivered() *MessageDelivered {
	if x, ok := x.GetType().(*SubscribeReply_MessageDelivered); ok {
		return x.MessageDelivered
	}
	return nil
}

func (x *SubscribeReply) GetMessageRead() *MessageRead {
	if x, ok := x.GetType().(*SubscribeReply_MessageRead); ok {
		return x.MessageRead
	}
	return nil
}

func (x *SubscribeReply) GetMessageReject() *MessageReject {
	if x, ok := x.GetType().(*SubscribeReply_MessageReject); ok {
		return x.MessageReject
	}
	return nil
}

func (x *SubscribeReply) GetMessageWithdraw() *MessageWithdraw {
	if x, ok := x.GetType().(*SubscribeReply_MessageWithdraw); ok {
		return x.MessageWithdraw
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

type SubscribeReply_MessageDelivered struct {
	MessageDelivered *MessageDelivered `protobuf:"bytes,6,opt,name=messageDelivered,proto3,oneof"` // 消息送达
}

type SubscribeReply_MessageRead struct {
	MessageRead *MessageRead `protobuf:"bytes,7,opt,name=messageRead,proto3,oneof"` // 消息已读
}

type SubscribeReply_MessageReject struct {
	MessageReject *MessageReject `protobuf:"bytes,8,opt,name=messageReject,proto3,oneof"` // 消息拒绝
}

type SubscribeReply_MessageWithdraw struct {
	MessageWithdraw *MessageWithdraw `protobuf:"bytes,9,opt,name=messageWithdraw,proto3,oneof"` // 消息撤回
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

func (*SubscribeReply_MessageDelivered) isSubscribeReply_Type() {}

func (*SubscribeReply_MessageRead) isSubscribeReply_Type() {}

func (*SubscribeReply_MessageReject) isSubscribeReply_Type() {}

func (*SubscribeReply_MessageWithdraw) isSubscribeReply_Type() {}

func (*SubscribeReply_FriendRequest) isSubscribeReply_Type() {}

func (*SubscribeReply_FriendReply) isSubscribeReply_Type() {}

var File_pkg_rpc_chat_p2p_chat_p2p_chat_proto protoreflect.FileDescriptor

var file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70,
	0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x22, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x80, 0x06, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a,
	0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x06,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x6f,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x55, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x48, 0x00, 0x52, 0x10, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x46, 0x0a,
	0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x12, 0x4c, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x52, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x4c, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32,
	0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00,
	0x52, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x32, 0x68, 0x0a, 0x07, 0x50, 0x32, 0x70, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x5d, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x27, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42,
	0x17, 0x5a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x70, 0x32, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescOnce sync.Once
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescData = file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDesc
)

func file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescGZIP() []byte {
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescData)
	})
	return file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDescData
}

var file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil), // 0: pkg.rpc.chat.p2p_chat.SubscribeRequest
	(*SubscribeReply)(nil),   // 1: pkg.rpc.chat.p2p_chat.SubscribeReply
	(*Online)(nil),           // 2: pkg.rpc.chat.p2p_chat.Online
	(*Offline)(nil),          // 3: pkg.rpc.chat.p2p_chat.Offline
	(*Inputting)(nil),        // 4: pkg.rpc.chat.p2p_chat.Inputting
	(*InputStop)(nil),        // 5: pkg.rpc.chat.p2p_chat.InputStop
	(*MessageDelivered)(nil), // 6: pkg.rpc.chat.p2p_chat.MessageDelivered
	(*MessageRead)(nil),      // 7: pkg.rpc.chat.p2p_chat.MessageRead
	(*MessageReject)(nil),    // 8: pkg.rpc.chat.p2p_chat.MessageReject
	(*MessageWithdraw)(nil),  // 9: pkg.rpc.chat.p2p_chat.MessageWithdraw
	(*FriendRequest)(nil),    // 10: pkg.rpc.chat.p2p_chat.FriendRequest
	(*FriendReply)(nil),      // 11: pkg.rpc.chat.p2p_chat.FriendReply
}
var file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_depIdxs = []int32{
	2,  // 0: pkg.rpc.chat.p2p_chat.SubscribeReply.online:type_name -> pkg.rpc.chat.p2p_chat.Online
	3,  // 1: pkg.rpc.chat.p2p_chat.SubscribeReply.offline:type_name -> pkg.rpc.chat.p2p_chat.Offline
	4,  // 2: pkg.rpc.chat.p2p_chat.SubscribeReply.inputting:type_name -> pkg.rpc.chat.p2p_chat.Inputting
	5,  // 3: pkg.rpc.chat.p2p_chat.SubscribeReply.inputStop:type_name -> pkg.rpc.chat.p2p_chat.InputStop
	6,  // 4: pkg.rpc.chat.p2p_chat.SubscribeReply.messageDelivered:type_name -> pkg.rpc.chat.p2p_chat.MessageDelivered
	7,  // 5: pkg.rpc.chat.p2p_chat.SubscribeReply.messageRead:type_name -> pkg.rpc.chat.p2p_chat.MessageRead
	8,  // 6: pkg.rpc.chat.p2p_chat.SubscribeReply.messageReject:type_name -> pkg.rpc.chat.p2p_chat.MessageReject
	9,  // 7: pkg.rpc.chat.p2p_chat.SubscribeReply.messageWithdraw:type_name -> pkg.rpc.chat.p2p_chat.MessageWithdraw
	10, // 8: pkg.rpc.chat.p2p_chat.SubscribeReply.friendRequest:type_name -> pkg.rpc.chat.p2p_chat.FriendRequest
	11, // 9: pkg.rpc.chat.p2p_chat.SubscribeReply.friendReply:type_name -> pkg.rpc.chat.p2p_chat.FriendReply
	0,  // 10: pkg.rpc.chat.p2p_chat.P2pChat.Subscribe:input_type -> pkg.rpc.chat.p2p_chat.SubscribeRequest
	1,  // 11: pkg.rpc.chat.p2p_chat.P2pChat.Subscribe:output_type -> pkg.rpc.chat.p2p_chat.SubscribeReply
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_init() }
func file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_init() {
	if File_pkg_rpc_chat_p2p_chat_p2p_chat_proto != nil {
		return
	}
	file_pkg_rpc_chat_p2p_chat_online_proto_init()
	file_pkg_rpc_chat_p2p_chat_offline_proto_init()
	file_pkg_rpc_chat_p2p_chat_inputting_proto_init()
	file_pkg_rpc_chat_p2p_chat_input_stop_proto_init()
	file_pkg_rpc_chat_p2p_chat_message_delivered_proto_init()
	file_pkg_rpc_chat_p2p_chat_message_read_proto_init()
	file_pkg_rpc_chat_p2p_chat_message_reject_proto_init()
	file_pkg_rpc_chat_p2p_chat_message_withdraw_proto_init()
	file_pkg_rpc_chat_p2p_chat_friend_request_proto_init()
	file_pkg_rpc_chat_p2p_chat_friend_reply_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SubscribeReply_Online)(nil),
		(*SubscribeReply_Offline)(nil),
		(*SubscribeReply_Inputting)(nil),
		(*SubscribeReply_InputStop)(nil),
		(*SubscribeReply_MessageDelivered)(nil),
		(*SubscribeReply_MessageRead)(nil),
		(*SubscribeReply_MessageReject)(nil),
		(*SubscribeReply_MessageWithdraw)(nil),
		(*SubscribeReply_FriendRequest)(nil),
		(*SubscribeReply_FriendReply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_msgTypes,
	}.Build()
	File_pkg_rpc_chat_p2p_chat_p2p_chat_proto = out.File
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_rawDesc = nil
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_goTypes = nil
	file_pkg_rpc_chat_p2p_chat_p2p_chat_proto_depIdxs = nil
}
