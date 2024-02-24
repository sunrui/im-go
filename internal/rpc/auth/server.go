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
	"internal/rpc/common"
)

// ImAuthServer 认证服务
type ImAuthServer struct {
	UnimplementedAuthServer
}

func (ImAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	println(fmt.Sprintf("unary Login message %q\n", req.Token))

	return &LoginReply{
		Reply: &common.Reply{
			Code:    common.ReplyCode_NO_AUTH,
			Message: nil,
		},
	}, nil
}

func (ImAuthServer) GetState(context.Context, *wrapperspb.Int32Value) (*GetStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

func (ImAuthServer) Logout(context.Context, *wrapperspb.Int32Value) (*common.Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func (ImAuthServer) mustEmbedImAuthServer() {}
