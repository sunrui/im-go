/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:02:16
 */

package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"internal/rpc/message/bottle_chat"

	"internal/rpc/message/room_chat"

	"internal/rpc/message/group_chat"

	"internal/rpc/message/p2p_chat"

	"internal/rpc/message/chat"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"internal/rpc/auth"
	"internal/rpc/interceptor"
)

// Client 客户端
type Client struct {
	config   Config   // 配置
	notifier Notifier // 通知

	authClient       auth.AuthClient              // 认证客户端
	chatClient       chat.ChatClient              // 聊天客户端
	p2pChatClient    p2p_chat.P2PChatClient       // 点对点聊天客户端
	groupChatClient  group_chat.GroupChatClient   // 群聊天客户端
	roomChatClient   room_chat.RoomChatClient     // 聊天室客户端
	bottleChatClient bottle_chat.BottleChatClient // 漂流瓶客户端
}

// NewClient 创建客户端
func NewClient(config Config, notifier Notifier) *Client {
	return &Client{
		config:   config,
		notifier: notifier,
	}
}

// Start 启动
func (client Client) Start() {
	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: client.config.AccessToken,
	})}
	creds, err := credentials.NewClientTLSFromFile(client.config.CertFile, client.config.ServerName)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(creds),
		grpc.WithChainUnaryInterceptor(interceptor.NewSequenceInterpreter().Client()),
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", client.config.Ip, client.config.Port), opts...)
	if err != nil {
		println(err.Error())
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	// 实例化一个client对象，传入参数conn
	client.authClient = auth.NewAuthClient(conn)
	client.chatClient = chat.NewChatClient(conn)
	client.p2pChatClient = p2p_chat.NewP2PChatClient(conn)
	client.groupChatClient = group_chat.NewGroupChatClient(conn)
	client.roomChatClient = room_chat.NewRoomChatClient(conn)
	client.bottleChatClient = bottle_chat.NewBottleChatClient(conn)

	authClient, err := client.authClient.Subscribe(context.Background(), &auth.NotifyRequest{
		UserId:  "userId",
		Message: "httpAddr",
	})
	if err != nil {
		println("authClient error = ", err.Error())
		return
	}

	for {
		notifyRes, err := authClient.Recv()

		if err == io.EOF {
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("Recv error:%v", err)
			break
		}

		println(notifyRes.Message)
	}

	time.Sleep(time.Hour)
}

// ChatTo 聊天
func (client Client) ChatTo(chatToRequest *chat.ToRequest) (*chat.ToReply, error) {
	return client.chatClient.To(context.Background(), chatToRequest)
}
