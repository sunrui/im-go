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

	OnChat(reply *chat.ReceiveReply)
	OnP2pChat(reply *p2p_chat.ReceiveReply)
	OnGroupChat(reply *group_chat.ReceiveReply)
	OnRoomChat(reply *room_chat.ReceiveReply)
	OnBottleChat(reply *bottle_chat.ReceiveReply)
}
