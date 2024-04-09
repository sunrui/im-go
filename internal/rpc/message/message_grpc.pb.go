/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:04:27
 */

package message

import (
	"context"
	"pkg/rpc/proto/chat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server 消息服务
type Server struct {
	chat.UnimplementedMessageServer
}

// ChatTo 发送消息
func (Server) ChatTo(context.Context, *chat.ToRequest) (*chat.ToReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatTo not implemented")
}

// Subscribe 订阅消息
func (Server) Subscribe(*chat.SubscribeRequest, chat.Message_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
