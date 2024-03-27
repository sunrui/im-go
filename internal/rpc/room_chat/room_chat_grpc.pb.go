/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:06:15
 */

package room_chat

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pkg/rpc/proto/room_chat"
)

// Server 消息服务
type Server struct {
	room_chat.UnimplementedRoomChatServer
}

// Subscribe 订阅消息
func (Server) Subscribe(*room_chat.SubscribeRequest, room_chat.RoomChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
