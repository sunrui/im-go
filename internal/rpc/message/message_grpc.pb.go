/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:04:27
 */

package message

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pkg/rpc/proto/message"
)

// Server 消息服务
type Server struct {
	message.UnimplementedMessageServer
}

// To 发送消息
func (Server) To(context.Context, *message.ToRequest) (*message.ToReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method To not implemented")
}

// Subscribe 订阅消息
func (Server) Subscribe(*message.SubscribeRequest, message.Message_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
