/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 18:03:32
 */

package token

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

func RequestIDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value("request-id")
		if requestID, ok := value.(string); ok && requestID != "" {
			println("requestId = ", requestID)
		} else {
			second := fmt.Sprintf("%d", time.Now().UnixMilli())
			println(second)
			//md["request-id"] = []string{second}
			md.Append("request-id", second)
		}
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

func RequestIDServerInterceptor() grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		//AuthInterceptor(ctx, req, info, handler)

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}
		// Set request ID for context.
		requestIDs := md["request-id"]
		if len(requestIDs) >= 1 {
			ctx = context.WithValue(ctx, "request-id", requestIDs[0])
			return handler(ctx, req)
		}

		// Generate request ID and set context if not exists.
		requestID := time.Now().Unix()
		ctx = context.WithValue(ctx, "request-id", requestID)
		return handler(ctx, req)
	}
}
