//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 13:59:16

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: internal/rpc/message/bottle_chat/bottle_chat.proto

package bottle_chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GroupChat_Receive_FullMethodName = "/internal.rpc.message.bottle_chat.GroupChat/Receive"
)

// GroupChatClient is the client API for GroupChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupChatClient interface {
	// 接收
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (GroupChat_ReceiveClient, error)
}

type groupChatClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupChatClient(cc grpc.ClientConnInterface) GroupChatClient {
	return &groupChatClient{cc}
}

func (c *groupChatClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (GroupChat_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &GroupChat_ServiceDesc.Streams[0], GroupChat_Receive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &groupChatReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GroupChat_ReceiveClient interface {
	Recv() (*ReceiveReply, error)
	grpc.ClientStream
}

type groupChatReceiveClient struct {
	grpc.ClientStream
}

func (x *groupChatReceiveClient) Recv() (*ReceiveReply, error) {
	m := new(ReceiveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GroupChatServer is the server API for GroupChat service.
// All implementations must embed UnimplementedGroupChatServer
// for forward compatibility
type GroupChatServer interface {
	// 接收
	Receive(*ReceiveRequest, GroupChat_ReceiveServer) error
	mustEmbedUnimplementedGroupChatServer()
}

// UnimplementedGroupChatServer must be embedded to have forward compatible implementations.
type UnimplementedGroupChatServer struct {
}

func (UnimplementedGroupChatServer) Receive(*ReceiveRequest, GroupChat_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedGroupChatServer) mustEmbedUnimplementedGroupChatServer() {}

// UnsafeGroupChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupChatServer will
// result in compilation errors.
type UnsafeGroupChatServer interface {
	mustEmbedUnimplementedGroupChatServer()
}

func RegisterGroupChatServer(s grpc.ServiceRegistrar, srv GroupChatServer) {
	s.RegisterService(&GroupChat_ServiceDesc, srv)
}

func _GroupChat_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GroupChatServer).Receive(m, &groupChatReceiveServer{stream})
}

type GroupChat_ReceiveServer interface {
	Send(*ReceiveReply) error
	grpc.ServerStream
}

type groupChatReceiveServer struct {
	grpc.ServerStream
}

func (x *groupChatReceiveServer) Send(m *ReceiveReply) error {
	return x.ServerStream.SendMsg(m)
}

// GroupChat_ServiceDesc is the grpc.ServiceDesc for GroupChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.rpc.message.bottle_chat.GroupChat",
	HandlerType: (*GroupChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _GroupChat_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/rpc/message/bottle_chat/bottle_chat.proto",
}
