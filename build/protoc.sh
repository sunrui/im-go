#!/usr/bin/env bash

cd ..

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/starter/auth/*.proto

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/message/*/*.proto