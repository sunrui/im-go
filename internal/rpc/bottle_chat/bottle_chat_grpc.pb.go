/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:07:37
 */

package bottle_chat

import (
	"pkg/rpc/proto/bottle_chat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server 消息服务
type Server struct {
	bottle_chat.UnimplementedBottleChatServer
}

// Subscribe 订阅消息
func (Server) Subscribe(*bottle_chat.SubscribeRequest, bottle_chat.BottleChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
