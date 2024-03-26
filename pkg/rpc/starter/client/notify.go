/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:26:22
 */

package client

import (
	"pkg/rpc/chat/bottle_chat"
	"pkg/rpc/chat/group_chat"
	"pkg/rpc/chat/message"
	"pkg/rpc/chat/p2p_chat"
	"pkg/rpc/chat/room_chat"
)

type Notifier interface {
	OnAuthFailed(reason string)
	OnError(err error)
	OnDisconnect()
	onClose()

	OnChat(reply *message.SubscribeReply)
	OnP2PChat(reply *p2p_chat.SubscribeReply)
	OnGroupChat(reply *group_chat.SubscribeReply)
	OnRoomChat(reply *room_chat.SubscribeReply)
	OnBottleChat(reply *bottle_chat.SubscribeReply)
}
