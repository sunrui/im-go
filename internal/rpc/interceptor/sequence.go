/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-19 22:22:21
 */

package interceptor

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const SequenceTag = "sequence"

// SequenceInterceptor 序号拦截器
type SequenceInterceptor struct{}

// NewSequenceInterpreter 新建序列号拦截器
func NewSequenceInterpreter() SequenceInterceptor {
	return SequenceInterceptor{}
}

// Client 客户端
func (sequenceInterceptor SequenceInterceptor) Client() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value(SequenceTag)
		if sequence, ok := value.(string); ok && sequence != "" {
			println("sequence = ", sequence)
		} else {
			md.Append(SequenceTag, fmt.Sprintf("%d", time.Now().UnixMilli()))
		}
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

// Server 服务器
func (sequenceInterceptor SequenceInterceptor) Server() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		requestIDs := md[SequenceTag]
		if len(requestIDs) >= 1 {
			ctx = context.WithValue(ctx, SequenceTag, requestIDs[0])
			return handler(ctx, req)
		}

		requestID := time.Now().Unix()
		ctx = context.WithValue(ctx, SequenceTag, requestID)
		return handler(ctx, req)
	}
}
