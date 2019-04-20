package main

import (
	"flag"

	"github.com/magodo/go_snippet/grpc_gateway/internal/server"
)

var (
	httpAddr = flag.String("http-addr", "localhost:9090", "http address")
	rpcAddr  = flag.String("rpc-addr", "localhost:9091", "http address")
)

func main() {
	flag.Parse()
	go server.RunRpc(*rpcAddr)
	server.RunHttp(*httpAddr, *rpcAddr)
}
