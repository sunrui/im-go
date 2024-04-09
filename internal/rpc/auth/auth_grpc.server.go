/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 01:15:59
 */

package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"pkg/rpc/interceptor"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// ImplAuthServer 认证服务
type ImplAuthServer struct {
	UnimplementedAuthServer
}

// isAuthenticated validates the authorization.
func isAuthenticated(authorization []string) (username string, err error) {
	if len(authorization) < 1 {
		return "", errors.New("received empty authorization token from client")
	}
	name := strings.TrimPrefix(authorization[0], "Bearer ")

	return name, nil
}

func getClientIP(ctx context.Context) string {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		println("[getClientIP] invoke FromContext() failed")
		return ""
	}

	if pr.Addr == net.Addr(nil) {
		println("[getClientIP] peer.Addr is nil")
		return ""
	}
	addSlice := strings.Split(pr.Addr.String(), ":")

	return addSlice[0] + ":" + addSlice[1]
}

// Login 登录
func (ImplAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	println(fmt.Sprintf("unary Login chat %q\n", req.Token))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "method GetState not implemented")
	}

	requestId := md[interceptor.SequenceTag]
	println("Sequence = ", requestId[0])

	userId, err := isAuthenticated(md["authorization"])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// if md["userId"] == nil {
	// 	println("context userId is nil, set is to " + userId)
	// 	md["userId"] = make([]string, 1)
	// 	md["userId"][0] = userId
	// 	ctx = context.WithValue(ctx, "userId", userId)
	// } else {
	// 	println("context userId value = %s", md["userId"][0])
	// }

	return &LoginReply{
		UserId: userId,
		Ip:     getClientIP(ctx),
	}, nil
}

func (ImplAuthServer) Subscribe(req *SubscribeRequest, stream Auth_SubscribeServer) error {
	log.Printf("Received %v", req.GetMessage())

	for {
		println("start notify chat")
		message := fmt.Sprintf("chat - %s", time.Now().Format("2006-01-02 15:04:05"))

		// 通过 send 方法不断推送数据
		err := stream.Send(&SubscribeResponse{
			Message: message,
		})

		ip := getClientIP(stream.Context())
		println("send chat to ip: ", ip, " - ", message)

		time.Sleep(time.Second)

		if err != nil {
			log.Fatalf("Send error:%v", err)
		}
	}

	return nil
}

func (ImplAuthServer) mustEmbedImAuthServer() {}
