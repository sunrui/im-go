/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:26:22
 */

package client

import (
	"pkg/rpc/proto/bottle_chat"
	"pkg/rpc/proto/group_chat"
	"pkg/rpc/proto/message"
	"pkg/rpc/proto/p2p_chat"
	"pkg/rpc/proto/room_chat"
)

// Notifier 通知
type Notifier interface {
	OnAuthFailed(reason string) // 认证失败
	OnError(err error)          // 错误
	OnDisconnect()              // 断开连接
	onClose()                   // 关闭

	OnMessage(reply *message.SubscribeReply)        // 消息
	OnP2PChat(reply *p2p_chat.SubscribeReply)       // 点对点聊天
	OnGroupChat(reply *group_chat.SubscribeReply)   // 群聊天
	OnRoomChat(reply *room_chat.SubscribeReply)     // 聊天室聊天
	OnBottleChat(reply *bottle_chat.SubscribeReply) // 漂流瓶聊天
}
