/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 01:31:30
 */

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"internal/rpc/auth"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:2024", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		println(err.Error())
	}

	defer func() {
		_ = conn.Close()
	}()

	// 实例化一个client对象，传入参数conn
	c := auth.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 调用SayHello方法，以请求服务，然后得到响应消息
	r, err := c.Login(ctx, &auth.LoginRequest{
		Token: "1234",
	})
	if err != nil {
		println(err.Error())
	} else {
		println(fmt.Sprintf("response from server: %q", r))
	}
}
