/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:04:46
 */

package p2p_chat

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pkg/rpc/proto/p2p_chat"
)

// Server 消息服务
type Server struct {
	p2p_chat.UnimplementedP2PChatServer
}

// Subscribe 订阅消息
func (Server) Subscribe(*p2p_chat.SubscribeRequest, p2p_chat.P2PChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
