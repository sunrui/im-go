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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net"
	"strings"
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

	println(fmt.Sprintf("%+v", pr))

	if pr.Addr == net.Addr(nil) {
		println("[getClientIP] peer.Addr is nil")
		return ""
	}
	addSlice := strings.Split(pr.Addr.String(), ":")
	return addSlice[0]
}

// Login 登录
func (ImplAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	println(fmt.Sprintf("unary Login message %q\n", req.Token))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "method GetState not implemented")
	}

	requestId := md["request-id"]
	println("requestId = ", requestId[0])

	userName, err := isAuthenticated(md["authorization"])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if ctx.Value("userId") == nil {
		println("context userId is nil, set is to " + userName)
		ctx = context.WithValue(ctx, "userId", userName)
	} else {
		println("context userId value = %s", ctx.Value("userId"))
	}

	return &LoginReply{
		UserId: userName,
		Ip:     getClientIP(ctx),
	}, nil
}

// GetState 获取状态
func (ImplAuthServer) GetState(context.Context, *wrapperspb.Int32Value) (*GetStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

// Logout 登出
func (ImplAuthServer) Logout(context.Context, *wrapperspb.Int32Value) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func (ImplAuthServer) mustEmbedImAuthServer() {}
