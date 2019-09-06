#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Mon 13 May 2019 02:05:39 PM CST
# Description:
#########################################################################

usage() {
    cat << EOF
Usage: ./${MYNAME} [options] name age

Options:
    -h|--help           show this message

Arguments:
    name
    age
EOF
}

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MYNAME="$(basename "${BASH_SOURCE[0]}")"

main() {
    name=$1
    age=$2
    cat << EOF
name: $name
age : $age
EOF
    env
    sleep 100
}
