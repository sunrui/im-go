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

	"pkg/rpc/starter/auth"

	"pkg/rpc/starter/interceptor"

	"pkg/rpc/message/bottle_chat"

	"pkg/rpc/message/room_chat"

	"pkg/rpc/message/group_chat"

	"pkg/rpc/message/p2p_chat"

	"pkg/rpc/message/chat"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

// Client 客户端
type Client struct {
	config   Config           // 配置
	notifier Notifier         // 通知
	conn     *grpc.ClientConn // 连接

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
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: client.config.AccessToken,
				}),
		}),
		grpc.WithTransportCredentials(func() credentials.TransportCredentials {
			if transportCredentials, err := credentials.NewClientTLSFromFile(
				client.config.CertFile, client.config.ServerName); err != nil {
				client.notifier.OnError(err)
				return nil
			} else {
				return transportCredentials
			}
		}()),
		grpc.WithChainUnaryInterceptor(interceptor.NewSequenceInterpreter().Client()),
	}

	if conn, err := grpc.Dial(fmt.Sprintf("%s:%d", client.config.Ip, client.config.Port), opts...); err != nil {
		client.notifier.OnError(err)
		return
	} else {
		client.conn = conn
		ctx := context.Background()

		go func() {
			client.authClient = auth.NewAuthClient(client.conn)
			if authSubscribeClient, err := client.authClient.Subscribe(ctx, &auth.SubscribeRequest{}); err != nil {
				client.notifier.OnError(err)
				return
			} else {
				for {
					if reply, err := authSubscribeClient.Recv(); err != nil {
						client.notifier.OnError(err)
					} else {
						if err == io.EOF {
							client.notifier.onClose()
							break
						}

						println(reply.ProtoMessage, reply.Message, reply.GetMessage())
					}
				}
			}
		}()

		go func() {
			client.chatClient = chat.NewChatClient(conn)
			if chatSubscribeClient, err := client.chatClient.Subscribe(ctx, &chat.SubscribeRequest{}); err != nil {
				client.notifier.OnError(err)
				return
			} else {
				for {
					if reply, err := chatSubscribeClient.Recv(); err != nil {
						client.notifier.OnError(err)
					} else {
						if err == io.EOF {
							client.notifier.onClose()
							break
						}

						client.notifier.OnChat(reply)
					}
				}
			}
		}()

		go func() {
			client.p2pChatClient = p2p_chat.NewP2PChatClient(conn)
			if p2pChatSubscribeClient, err := client.p2pChatClient.Subscribe(ctx, &p2p_chat.SubscribeRequest{}); err != nil {
				client.notifier.OnError(err)
				return
			} else {
				for {
					if reply, err := p2pChatSubscribeClient.Recv(); err != nil {
						client.notifier.OnError(err)
					} else {
						if err == io.EOF {
							client.notifier.onClose()
							break
						}

						client.notifier.OnP2PChat(reply)
					}
				}
			}
		}()

		go func() {
			client.groupChatClient = group_chat.NewGroupChatClient(conn)
			if groupChatSubscribeClient, err := client.groupChatClient.Subscribe(ctx, &group_chat.SubscribeRequest{}); err != nil {
				client.notifier.OnError(err)
				return
			} else {
				for {
					if reply, err := groupChatSubscribeClient.Recv(); err != nil {
						client.notifier.OnError(err)
					} else {
						if err == io.EOF {
							client.notifier.onClose()
							break
						}

						client.notifier.OnGroupChat(reply)
					}
				}
			}
		}()

		go func() {
			client.roomChatClient = room_chat.NewRoomChatClient(conn)
			if roomChatSubscribeClient, err := client.roomChatClient.Subscribe(ctx, &room_chat.SubscribeRequest{}); err != nil {
				client.notifier.OnError(err)
				return
			} else {
				for {
					if reply, err := roomChatSubscribeClient.Recv(); err != nil {
						client.notifier.OnError(err)
					} else {
						if err == io.EOF {
							client.notifier.onClose()
							break
						}

						client.notifier.OnRoomChat(reply)
					}
				}
			}
		}()

		go func() {
			client.bottleChatClient = bottle_chat.NewBottleChatClient(conn)
			if bottleChatSubscribeClient, err := client.bottleChatClient.Subscribe(ctx, &bottle_chat.SubscribeRequest{}); err != nil {
				client.notifier.OnError(err)
				return
			} else {
				for {
					if reply, err := bottleChatSubscribeClient.Recv(); err != nil {
						client.notifier.OnError(err)
					} else {
						if err == io.EOF {
							client.notifier.onClose()
							break
						}

						client.notifier.OnBottleChat(reply)
					}
				}
			}
		}()
	}
}

// ChatTo 聊天
func (client Client) ChatTo(chatToRequest *chat.ToRequest) (*chat.ToReply, error) {
	return client.chatClient.To(context.Background(), chatToRequest)
}

// Close 关闭
func (client Client) Close() {
	if err := client.conn.Close(); err != nil {
		client.notifier.OnError(err)
	}
}
