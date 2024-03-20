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
	"strings"
	"sync"
	"time"

	"internal/rpc/interceptor"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
	println(fmt.Sprintf("unary Login message %q\n", req.Token))

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

// GetState 获取状态
func (ImplAuthServer) GetState(context.Context, *wrapperspb.Int32Value) (*GetStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

// Logout 登出
func (ImplAuthServer) Logout(context.Context, *wrapperspb.Int32Value) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

var WaitGroup sync.WaitGroup

func (ImplAuthServer) Subscribe(req *NotifyRequest, stream Auth_SubscribeServer) error {
	log.Printf("Received %v", req.GetMessage())

	for {
		WaitGroup.Add(1)

		println("start push message")
		message := fmt.Sprintf("message - %s", time.Now().Format("2006-01-02 15:04:05"))

		// 通过 send 方法不断推送数据
		err := stream.Send(&NotifyResponse{
			Message: message,
		})

		ip := getClientIP(stream.Context())
		println("send message to ip: ", ip, " - ", message)

		time.Sleep(time.Second)

		if err != nil {
			log.Fatalf("Send error:%v", err)
		}

		WaitGroup.Done()
		WaitGroup.Wait()
	}

	return nil
}

func (ImplAuthServer) mustEmbedImAuthServer() {}
