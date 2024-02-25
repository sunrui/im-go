/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 01:15:59
 */

package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ImplAuthServer 认证服务
type ImplAuthServer struct {
	UnimplementedAuthServer
}

// Login 登录
func (ImplAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	println(fmt.Sprintf("unary Login message %q\n", req.Token))

	return &LoginReply{}, nil
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
