/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:26:22
 */

package client

import (
	"internal/rpc/message/bottle_chat"
	"internal/rpc/message/chat"
	"internal/rpc/message/group_chat"
	"internal/rpc/message/p2p_chat"
	"internal/rpc/message/room_chat"
)

type Notifier interface {
	OnAuthFailed(reason string)
	OnError(err error)
	OnDisconnect()
	onClose()

	OnChat(reply *chat.SubscribeReply)
	OnP2PChat(reply *p2p_chat.SubscribeReply)
	OnGroupChat(reply *group_chat.SubscribeReply)
	OnRoomChat(reply *room_chat.SubscribeReply)
	OnBottleChat(reply *bottle_chat.SubscribeReply)
}
