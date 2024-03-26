/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:19:24
 */

package client

import (
	"testing"
	"time"

	"pkg/rpc/chat/bottle_chat"
	"pkg/rpc/chat/group_chat"
	"pkg/rpc/chat/message"
	"pkg/rpc/chat/p2p_chat"
	"pkg/rpc/chat/room_chat"
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

func (n Notify) onClose() {
	println("OnClose")
}

func (n Notify) OnChat(reply *message.SubscribeReply) {
	println("OnChat, reply", reply.Chat.String())
}

func (n Notify) OnP2PChat(reply *p2p_chat.SubscribeReply) {
	println("OnChat, reply", reply.String())
}

func (n Notify) OnGroupChat(reply *group_chat.SubscribeReply) {
	println("OnChat, reply", reply.String())
}

func (n Notify) OnRoomChat(reply *room_chat.SubscribeReply) {
	println("OnChat, reply", reply.String())
}

func (n Notify) OnBottleChat(reply *bottle_chat.SubscribeReply) {
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

	// reply, err := client.ChatTo(&message.ToRequest{
	// 	Source: &message.Source{
	// 		SequenceId: "1",
	// 	},
	// })
	//
	// println(reply, err)
	time.Sleep(time.Hour)
}
