/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 16:14:45
 */

package server

import (
	"context"
	"testing"

	"pkg/rpc/proto/bottle_chat"

	"pkg/rpc/proto/room_chat"

	"pkg/rpc/proto/group_chat"

	"pkg/rpc/proto/p2p_chat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pkg/rpc/proto/message"
)

type Notify struct{}

func (n Notify) OnError(err error) {
	println(err.Error())
}

func (n Notify) onClose() {
	println("onClose")
}

type MessageServer struct {
	message.UnimplementedMessageServer
}

// To 发送消息
func (MessageServer) To(context.Context, *message.ToRequest) (*message.ToReply, error) {
	return &message.ToReply{
		SequenceId: "test_sequence_id",
		Status:     0,
		Comment:    "comment",
	}, nil

	return nil, status.Errorf(codes.Unimplemented, "method To not implemented")
}

// Subscribe 订阅消息
func (MessageServer) Subscribe(*message.SubscribeRequest, message.Message_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

type P2PChatServer struct {
	p2p_chat.UnimplementedP2PChatServer
}

// Subscribe 订阅消息
func (P2PChatServer) Subscribe(*p2p_chat.SubscribeRequest, p2p_chat.P2PChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

type GroupChatServer struct {
	group_chat.UnimplementedGroupChatServer
}

// Subscribe 订阅消息
func (GroupChatServer) Subscribe(*group_chat.SubscribeRequest, group_chat.GroupChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

type RoomChatServer struct {
	room_chat.UnimplementedRoomChatServer
}

// Subscribe 订阅消息
func (RoomChatServer) Subscribe(*room_chat.SubscribeRequest, room_chat.RoomChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

type BottleChatServer struct {
	bottle_chat.UnimplementedBottleChatServer
}

// Subscribe 订阅消息
func (BottleChatServer) Subscribe(*bottle_chat.SubscribeRequest, bottle_chat.BottleChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func TestServer(t *testing.T) {
	config := Config{
		CertFile: "../../../build/x509/server_cert.pem",
		KeyFile:  "../../../build/x509/server_key.pem",
		Port:     2024,
	}

	server := NewServer(config, Notify{}, Registrar{
		MessageServer:    MessageServer{},
		P2pChatServer:    P2PChatServer{},
		GroupChatServer:  GroupChatServer{},
		RoomChatServer:   RoomChatServer{},
		BottleChatServer: BottleChatServer{},
	})
	server.Start()
}
