/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:07:01
 */

package group_chat

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pkg/rpc/proto/group_chat"
)

// Server 消息服务
type Server struct {
	group_chat.UnimplementedGroupChatServer
}

// Subscribe 订阅消息
func (Server) Subscribe(*group_chat.SubscribeRequest, group_chat.GroupChat_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
