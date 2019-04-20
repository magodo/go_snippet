#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Tue 16 Apr 2019 10:31:29 AM CST
# Description:
#########################################################################

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MYNAME="$(basename "${BASH_SOURCE[0]}")"

usage() {
    cat << EOF
Usage: ./${MYNAME} [options] api_version

Options:
    -h|--help           show this message

Arguments:
    api_version         version of api to use (e.g. v1)
EOF
}

main() {
    while :; do
        case $1 in
            -h|--help)
                usage
                exit 1
                ;;
            --)
                shift
                break
                ;;
            *)
                break
                ;;
        esac
        shift
    done
    local expect_n_arg
    expect_n_arg=1
    [[ $# = "$expect_n_arg" ]] || die "wrong arguments (expected: $expect_n_arg, got: $#)"

    api_version=$1

    proto_dir="$MYDIR"/../api/proto/"$api_version"
    output_dir="$MYDIR"/../pkg/api/"$api_version"
    swagger_output_dir="$MYDIR"/../api/swagger/"$api_version"
    protoc -I "$proto_dir" -I "$MYDIR/../third_party" --go_out=plugins=grpc:"$output_dir" todo-service.proto
    protoc -I "$proto_dir" -I "$MYDIR/../third_party" --grpc-gateway_out=logtostderr=true:"$output_dir" todo-service.proto
    protoc -I "$proto_dir" -I "$MYDIR/../third_party" --swagger_out=logtostderr=true:"$swagger_output_dir" todo-service.proto
}

main "$@"
