#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Sat 20 Apr 2019 09:37:24 PM CST
# Description:
#########################################################################

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MYNAME="$(basename "${BASH_SOURCE[0]}")"

PROJECT_DIR="$MYDIR/../"
RPC_GEN_DIR="$PROJECT_DIR/internal/api/rpc"
RESTFUL_GEN_DIR="$PROJECT_DIR/internal/api/restful"

protoc -I /usr/local/include -I "$PROJECT_DIR"/api/proto -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis \
    --go_out=plugins=grpc:"$RPC_GEN_DIR" \
    --micro_out="$RPC_GEN_DIR" \
    "$PROJECT_DIR/api/proto/greeter.proto"

protoc -I /usr/local/include -I "$PROJECT_DIR"/api/proto -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis \
    --go_out=plugins=grpc:"$RESTFUL_GEN_DIR" \
    --grpc-gateway_out=logtostderr=true:"$RESTFUL_GEN_DIR" \
    "$PROJECT_DIR/api/proto/greeter.proto"
