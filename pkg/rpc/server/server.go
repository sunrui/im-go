/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 16:14:35
 */

package server

import (
	"crypto/tls"
	"fmt"
	"net"

	"pkg/rpc/proto/bottle_chat"

	"pkg/rpc/proto/room_chat"

	"pkg/rpc/proto/group_chat"

	"pkg/rpc/proto/p2p_chat"

	"pkg/rpc/proto/message"

	"pkg/rpc/auth"
	"pkg/rpc/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// Server 服务
type Server struct {
	config    Config    // 配置
	notifier  Notifier  // 通知
	registrar Registrar // 注册
}

// NewServer 创建服务
func NewServer(config Config, notifier Notifier, registrar Registrar) Server {
	return Server{
		config:    config,
		notifier:  notifier,
		registrar: registrar,
	}
}

// Start 启动
func (server Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", server.config.Port))
	if err != nil {
		server.notifier.OnError(err)
		return
	}

	cert, err := tls.LoadX509KeyPair(server.config.CertFile, server.config.KeyFile)
	if err != nil {
		server.notifier.OnError(err)
		return
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.ChainUnaryInterceptor(interceptor.ServerAuth, interceptor.NewSequenceInterpreter().Server()),
	}

	s := grpc.NewServer(opts...)

	auth.RegisterAuthServer(s, &auth.ImplAuthServer{})

	message.RegisterMessageServer(s, server.registrar.MessageServer)
	p2p_chat.RegisterP2PChatServer(s, server.registrar.P2pChatServer)
	group_chat.RegisterGroupChatServer(s, server.registrar.GroupChatServer)
	room_chat.RegisterRoomChatServer(s, server.registrar.RoomChatServer)
	bottle_chat.RegisterBottleChatServer(s, server.registrar.BottleChatServer)
	reflection.Register(s)

	defer func() {
		s.Stop()
		_ = listen.Close()
	}()

	println(fmt.Sprintf("listen at %d", server.config.Port))
	if err = s.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
