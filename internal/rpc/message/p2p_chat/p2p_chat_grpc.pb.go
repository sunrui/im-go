//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 11:52:02

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: internal/rpc/message/p2p_chat/p2p_chat.proto

package p2p_chat

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
	P2PChat_Subscribe_FullMethodName = "/internal.rpc.message.p2p_chat.P2pChat/Subscribe"
)

// P2PChatClient is the client API for P2PChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PChatClient interface {
	// 订阅
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (P2PChat_SubscribeClient, error)
}

type p2PChatClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PChatClient(cc grpc.ClientConnInterface) P2PChatClient {
	return &p2PChatClient{cc}
}

func (c *p2PChatClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (P2PChat_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &P2PChat_ServiceDesc.Streams[0], P2PChat_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PChatSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P2PChat_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type p2PChatSubscribeClient struct {
	grpc.ClientStream
}

func (x *p2PChatSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// P2PChatServer is the server API for P2PChat service.
// All implementations must embed UnimplementedP2PChatServer
// for forward compatibility
type P2PChatServer interface {
	// 订阅
	Subscribe(*SubscribeRequest, P2PChat_SubscribeServer) error
	mustEmbedUnimplementedP2PChatServer()
}

// UnimplementedP2PChatServer must be embedded to have forward compatible implementations.
type UnimplementedP2PChatServer struct {
}

func (UnimplementedP2PChatServer) Subscribe(*SubscribeRequest, P2PChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedP2PChatServer) mustEmbedUnimplementedP2PChatServer() {}

// UnsafeP2PChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PChatServer will
// result in compilation errors.
type UnsafeP2PChatServer interface {
	mustEmbedUnimplementedP2PChatServer()
}

func RegisterP2PChatServer(s grpc.ServiceRegistrar, srv P2PChatServer) {
	s.RegisterService(&P2PChat_ServiceDesc, srv)
}

func _P2PChat_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P2PChatServer).Subscribe(m, &p2PChatSubscribeServer{stream})
}

type P2PChat_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type p2PChatSubscribeServer struct {
	grpc.ServerStream
}

func (x *p2PChatSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

// P2PChat_ServiceDesc is the grpc.ServiceDesc for P2PChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2PChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.rpc.message.p2p_chat.P2pChat",
	HandlerType: (*P2PChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _P2PChat_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/rpc/message/p2p_chat/p2p_chat.proto",
}
