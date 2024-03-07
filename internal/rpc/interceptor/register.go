/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-01 19:15:55
 */

package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func ChainUnary(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	// n := len(interceptors)
	// if n > 1 {
	// lastI := n - 1
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// var (
		// 	chainHandler grpc.UnaryInvoker
		// 	curI         int
		// )

		// func(ctx context.Context, method string, req, reply any, cc *ClientConn, opts ...CallOption) error

		// chainHandler = func(...) error {
		// 	if curI == lastI {
		// 		return invoker(...)
		// 	}
		// 	curI++
		// 	err := interceptors[curI](...)
		// 	curI--
		// 	return err
		// }
		// // 上述所省略的入参与该方法调用的入参一致
		// return interceptors[0](ctx, method, req, reply, cc, chainHandler, opts...)

		return nil
	}
}
