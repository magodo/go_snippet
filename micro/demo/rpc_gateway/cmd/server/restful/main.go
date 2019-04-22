package main

import (
	"flag"

	"github.com/magodo/rpc_gateway/pkg/server/restful"
)

var (
	endpoint *string = flag.String("endpoint", "localhost:8080", "rpc endpoint address")
)

func main() {
	flag.Parse()
	restful.Serve(*endpoint)
}
