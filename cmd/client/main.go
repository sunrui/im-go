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
	// 连接gRPC服务器
	conn, err := grpc.Dial("127.0.0.1:2024", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		println(err.Error())
	}
	defer conn.Close()

	// 实例化一个client对象，传入参数conn
	c := auth.NewAuthClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//延迟关闭请求会话
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
