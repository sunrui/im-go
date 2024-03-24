#!/usr/bin/env bash

cd ..

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/auth/auth.proto

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/proto/chat/*.proto
