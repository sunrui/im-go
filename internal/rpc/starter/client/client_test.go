/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:19:24
 */

package client

import (
	"testing"

	"internal/rpc/message/bottle_chat"
	"internal/rpc/message/chat"
	"internal/rpc/message/group_chat"
	"internal/rpc/message/p2p_chat"
	"internal/rpc/message/room_chat"
)

type Notify struct{}

func (n Notify) OnAuthFailed(reason string) {
	println("OnAuthFailed，reason: %s", reason)
}

func (n Notify) OnError(err error) {
	println("OnError, err: %s", err.Error())
}

func (n Notify) OnDisconnect() {
	println("OnDisconnect")
}

func (n Notify) OnChat(reply *chat.ReceiveReply) {
	println("OnChat, reply", reply.Content.String())
}

func (n Notify) OnP2pChat(reply *p2p_chat.ReceiveReply) {
	println("OnChat, reply", reply.String())
}

func (n Notify) OnGroupChat(reply *group_chat.ReceiveReply) {
	println("OnChat, reply", reply.String())
}

func (n Notify) OnRoomChat(reply *room_chat.ReceiveReply) {
	println("OnChat, reply", reply.String())
}

func (n Notify) OnBottleChat(reply *bottle_chat.ReceiveReply) {
	println("OnChat, reply", reply.String())
}

func TestClient_Start(t *testing.T) {
	client := NewClient(Config{
		CertFile:    "../../../../build/x509/ca_cert.pem",
		ServerName:  "*.honeysense.com",
		Ip:          "127.0.0.1",
		Port:        2024,
		AccessToken: "some-secret-token",
	}, &Notify{})

	client.Start()

	reply, err := client.ChatTo(&chat.ToRequest{
		Source: &chat.Source{
			SequenceId: "1",
		},
	})

	println(reply, err)
}
