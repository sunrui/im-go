/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-26 20:53:03
 */

package interceptor

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 请求标记
const requestIdTag = "request-id"

func RequestIdClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value(requestIdTag)
		if requestID, ok := value.(string); ok && requestID != "" {
			println("requestId = ", requestID)
		} else {
			second := fmt.Sprintf("%d", time.Now().UnixMilli())
			println(second)
			// md[requestIdTag] = []string{second}
			md.Append(requestIdTag, second)
		}
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

func RequestIdServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// ServerAuth(ctx, req, info, handler)

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}
		// Set request ID for context.
		requestIDs := md[requestIdTag]
		if len(requestIDs) >= 1 {
			ctx = context.WithValue(ctx, requestIdTag, requestIDs[0])
			return handler(ctx, req)
		}

		// Generate request ID and set context if not exists.
		requestID := time.Now().Unix()
		ctx = context.WithValue(ctx, requestIdTag, requestID)
		return handler(ctx, req)
	}
}
