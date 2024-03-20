/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 01:31:46
 */

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"internal/rpc/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"internal/rpc/auth"
)

func main() {
	listen, err := net.Listen("tcp", ":2024")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	cert, err := tls.LoadX509KeyPair("../../build/x509/server_cert.pem", "../../build/x509/server_key.pem")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		// grpc.UnaryInterceptor(interceptor.ServerAuth),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		// grpc.UnaryInterceptor(interceptor.RequestIdServerInterceptor()),
		grpc.ChainUnaryInterceptor(interceptor.ServerAuth, interceptor.NewSequenceInterpreter().Server()),
	}

	s := grpc.NewServer(opts...)

	implAuthServer := auth.ImplAuthServer{}

	auth.RegisterAuthServer(s, &implAuthServer)
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
