/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:26:22
 */

package client

import (
	"pkg/rpc/message/bottle_chat"
	"pkg/rpc/message/chat"
	"pkg/rpc/message/group_chat"
	"pkg/rpc/message/p2p_chat"
	"pkg/rpc/message/room_chat"
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
