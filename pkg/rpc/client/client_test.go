/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:19:24
 */

package client

import (
	"fmt"
	"testing"
	"time"

	"pkg/rpc/proto/bottle_chat"
	"pkg/rpc/proto/group_chat"
	"pkg/rpc/proto/message"
	"pkg/rpc/proto/p2p_chat"
	"pkg/rpc/proto/room_chat"
)

type Notify struct{}

func (n Notify) OnAuthFailed(reason string) {
	println(fmt.Sprintf("OnAuthFailed，reason: %s", reason))
}

func (n Notify) OnError(err error) {
	println(fmt.Sprintf("OnError, err: %s", err.Error()))
}

func (n Notify) OnDisconnect() {
	println("OnDisconnect")
}

func (n Notify) onClose() {
	println("OnClose")
}

func (n Notify) OnMessage(reply *message.SubscribeReply) {
	println(fmt.Sprintf("OnMessage, reply: %s", reply.Chat.String()))
}

func (n Notify) OnP2PChat(reply *p2p_chat.SubscribeReply) {
	println(fmt.Sprintf("OnMessage, reply: %s", reply.String()))
}

func (n Notify) OnGroupChat(reply *group_chat.SubscribeReply) {
	println(fmt.Sprintf("OnMessage, reply: %s", reply.String()))
}

func (n Notify) OnRoomChat(reply *room_chat.SubscribeReply) {
	println(fmt.Sprintf("OnMessage, reply: %s", reply.String()))
}

func (n Notify) OnBottleChat(reply *bottle_chat.SubscribeReply) {
	println(fmt.Sprintf("OnMessage, reply: %s", reply.String()))
}

func TestClient_Start(t *testing.T) {
	client := NewClient(Config{
		CertFile:    "../../../build/x509/ca_cert.pem",
		ServerName:  "*.honeysense.com",
		Ip:          "127.0.0.1",
		Port:        2024,
		AccessToken: "some-secret-token",
	}, &Notify{})

	client.Start()

	time.Sleep(time.Second)

	for {
		if reply, err := client.ChatTo(&message.ToRequest{}); err != nil {
			println(err.Error())
		} else {
			println("reply", reply.SequenceId, reply.Status, reply.Comment)
		}

		time.Sleep(5 * time.Second)
	}
}
