/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-02-25 01:31:30
 */

package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"internal/rpc/auth"
	"log"
	"time"
)

func main() {
	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(fetchToken())}
	creds, err := credentials.NewClientTLSFromFile("../../build/x509/ca_cert.pem", "*.honeysense.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(creds),
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("127.0.0.1:2024", opts...)
	if err != nil {
		println(err.Error())
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	// 实例化一个client对象，传入参数conn
	c := auth.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		if r, err := c.Login(ctx, &auth.LoginRequest{
			Token: "1234",
		}); err != nil {
			println(err.Error())
		} else {
			println(fmt.Sprintf("response from server: %q", r))
		}
	}
}

// fetchToken simulates a token lookup and omits the details of proper token
// acquisition. For examples of how to acquire an OAuth2 token, see:
// https://godoc.org/golang.org/x/oauth2
func fetchToken() *oauth2.Token {
	println("fetch token.")

	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
