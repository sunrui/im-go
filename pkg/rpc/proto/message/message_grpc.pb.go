//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-07 23:31:57

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/rpc/proto/message/message.proto

package message

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
	Message_ChatTo_FullMethodName    = "/pkg.rpc.proto.message.Message/ChatTo"
	Message_Subscribe_FullMethodName = "/pkg.rpc.proto.message.Message/Subscribe"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	// 发送
	ChatTo(ctx context.Context, in *ToRequest, opts ...grpc.CallOption) (*ToReply, error)
	// 订阅
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Message_SubscribeClient, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) ChatTo(ctx context.Context, in *ToRequest, opts ...grpc.CallOption) (*ToReply, error) {
	out := new(ToReply)
	err := c.cc.Invoke(ctx, Message_ChatTo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Message_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Message_ServiceDesc.Streams[0], Message_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Message_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type messageSubscribeClient struct {
	grpc.ClientStream
}

func (x *messageSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	// 发送
	ChatTo(context.Context, *ToRequest) (*ToReply, error)
	// 订阅
	Subscribe(*SubscribeRequest, Message_SubscribeServer) error
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) ChatTo(context.Context, *ToRequest) (*ToReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatTo not implemented")
}
func (UnimplementedMessageServer) Subscribe(*SubscribeRequest, Message_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_ChatTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).ChatTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_ChatTo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).ChatTo(ctx, req.(*ToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServer).Subscribe(m, &messageSubscribeServer{stream})
}

type Message_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type messageSubscribeServer struct {
	grpc.ServerStream
}

func (x *messageSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.rpc.proto.message.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChatTo",
			Handler:    _Message_ChatTo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Message_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/rpc/proto/message/message.proto",
}
