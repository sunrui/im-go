#!/usr/bin/env bash

cd ..

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/proto/chat/type.proto

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/proto/chat/type_text.proto

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/proto/chat/type_sticker.proto

protoc -I=./ \
       --go_out=./ --go_opt=paths=source_relative \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       ./internal/rpc/proto/chat/chat.proto
