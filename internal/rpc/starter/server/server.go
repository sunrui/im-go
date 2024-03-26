/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 16:14:35
 */

package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	auth2 "internal/rpc/starter/auth"
	interceptor2 "internal/rpc/starter/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func Server() {
	listen, err := net.Listen("tcp", ":2024")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	cert, err := tls.LoadX509KeyPair("../../../../build/x509/server_cert.pem", "../../../../build/x509/server_key.pem")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		// grpc.UnaryInterceptor(interceptor.ServerAuth),
		grpc.ChainUnaryInterceptor(interceptor2.ServerAuth, interceptor2.NewSequenceInterpreter().Server()),
	}

	s := grpc.NewServer(opts...)

	auth2.RegisterAuthServer(s, &auth2.ImplAuthServer{})
	reflection.Register(s)

	defer func() {
		s.Stop()
		_ = listen.Close()
	}()

	println("listen at 2024")

	// go func() {
	// for {
	// time.Sleep(time.Second * 5)
	// println("waitGroup add 1")
	// auth.WaitGroup.Add(1)
	// }
	// }()

	if err = s.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
