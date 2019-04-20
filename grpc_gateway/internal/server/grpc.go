package server

import (
	"log"
	"net"

	barPb "github.com/magodo/go_snippet/grpc_gateway/internal/api/proto/bar"
	fooPb "github.com/magodo/go_snippet/grpc_gateway/internal/api/proto/foo"
	"github.com/magodo/go_snippet/grpc_gateway/internal/api/service/bar"
	"github.com/magodo/go_snippet/grpc_gateway/internal/api/service/foo"
	"google.golang.org/grpc"
)

func RunRpc(rpcAddr string) {
	l, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	rpcServer := grpc.NewServer()
	fooPb.RegisterFooServiceServer(rpcServer, &foo.FooService{})
	barPb.RegisterBarServiceServer(rpcServer, &bar.BarService{})
	rpcServer.Serve(l)
}
