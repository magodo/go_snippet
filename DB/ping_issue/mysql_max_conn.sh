#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Thu 21 Feb 2019 10:40:30 AM CST
# Description:
#########################################################################

usage() {
    cat << EOF
Usage: ./${MYNAME} [options] max_conn

Options:
    -h|--help           show this message
Arguments:
    max_conn
EOF
}

main() {
    conn_cmd=(mysql '-h' '127.0.0.1' '-u' 'root' '-p123' 'mysql')
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

    max_conn=$1

    # here decrease one to eliminate this connection itself
    current_conn="$(( $("${conn_cmd[@]}" <<< 'select * from information_schema.PROCESSLIST where host != ""' | sed -n '2,$p' | wc -l)-1 ))"
    "${conn_cmd[@]}" <<< "set global max_connections=$((max_conn))"

    pids=()
    # here increase one because mysql has one more reserved connection for super user
    for i in $(seq $(( "$max_conn" - "$current_conn" + 1 ))); do
        coproc PROC"$i" { "${conn_cmd[@]}" ;}
        pids+=($!)
    done

    read -n 1 -r -s -p 'Enter any key to continue...'
    
    for pid in "${pids[@]}"; do
        kill $pid
    done
}

main "$@"
