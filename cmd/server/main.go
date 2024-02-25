/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 01:31:46
 */

package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"internal/rpc/auth"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":2024")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()

	auth.RegisterAuthServer(s, &auth.ImplAuthServer{})
	reflection.Register(s)

	defer func() {
		s.Stop()
		_ = listen.Close()
	}()

	println("listen at 2024")

	if err = s.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
