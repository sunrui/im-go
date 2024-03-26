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
	BottleChat_Receive_FullMethodName = "/internal.rpc.message.bottle_chat.BottleChat/Receive"
)

// BottleChatClient is the client API for BottleChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BottleChatClient interface {
	// 接收
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (BottleChat_ReceiveClient, error)
}

type bottleChatClient struct {
	cc grpc.ClientConnInterface
}

func NewBottleChatClient(cc grpc.ClientConnInterface) BottleChatClient {
	return &bottleChatClient{cc}
}

func (c *bottleChatClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (BottleChat_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &BottleChat_ServiceDesc.Streams[0], BottleChat_Receive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bottleChatReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BottleChat_ReceiveClient interface {
	Recv() (*ReceiveReply, error)
	grpc.ClientStream
}

type bottleChatReceiveClient struct {
	grpc.ClientStream
}

func (x *bottleChatReceiveClient) Recv() (*ReceiveReply, error) {
	m := new(ReceiveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BottleChatServer is the server API for BottleChat service.
// All implementations must embed UnimplementedBottleChatServer
// for forward compatibility
type BottleChatServer interface {
	// 接收
	Receive(*ReceiveRequest, BottleChat_ReceiveServer) error
	mustEmbedUnimplementedBottleChatServer()
}

// UnimplementedBottleChatServer must be embedded to have forward compatible implementations.
type UnimplementedBottleChatServer struct {
}

func (UnimplementedBottleChatServer) Receive(*ReceiveRequest, BottleChat_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedBottleChatServer) mustEmbedUnimplementedBottleChatServer() {}

// UnsafeBottleChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BottleChatServer will
// result in compilation errors.
type UnsafeBottleChatServer interface {
	mustEmbedUnimplementedBottleChatServer()
}

func RegisterBottleChatServer(s grpc.ServiceRegistrar, srv BottleChatServer) {
	s.RegisterService(&BottleChat_ServiceDesc, srv)
}

func _BottleChat_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BottleChatServer).Receive(m, &bottleChatReceiveServer{stream})
}

type BottleChat_ReceiveServer interface {
	Send(*ReceiveReply) error
	grpc.ServerStream
}

type bottleChatReceiveServer struct {
	grpc.ServerStream
}

func (x *bottleChatReceiveServer) Send(m *ReceiveReply) error {
	return x.ServerStream.SendMsg(m)
}

// BottleChat_ServiceDesc is the grpc.ServiceDesc for BottleChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BottleChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.rpc.message.bottle_chat.BottleChat",
	HandlerType: (*BottleChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _BottleChat_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/rpc/message/bottle_chat/bottle_chat.proto",
}
