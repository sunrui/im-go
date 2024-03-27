/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:44:57
 */

package server

import (
	"pkg/rpc/proto/bottle_chat"
	"pkg/rpc/proto/group_chat"
	"pkg/rpc/proto/message"
	"pkg/rpc/proto/p2p_chat"
	"pkg/rpc/proto/room_chat"
)

// Registrar 注册
type Registrar struct {
	MessageServer    message.MessageServer        // 消息服务
	P2pChatServer    p2p_chat.P2PChatServer       // 点对点聊天服务
	GroupChatServer  group_chat.GroupChatServer   // 群聊天服务
	RoomChatServer   room_chat.RoomChatServer     // 聊天室服务
	BottleChatServer bottle_chat.BottleChatServer // 漂流瓶服务
}
