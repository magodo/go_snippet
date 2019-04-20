package server

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/magodo/go_snippet/grpc_gateway/internal/api/proto/bar"
	"github.com/magodo/go_snippet/grpc_gateway/internal/api/proto/foo"
	"google.golang.org/grpc"
)

func RunHttp(httpAddr string, rpcAddr string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := foo.RegisterFooServiceHandlerFromEndpoint(ctx, mux, rpcAddr, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = bar.RegisterBarServiceHandlerFromEndpoint(ctx, mux, rpcAddr, opts)
	if err != nil {
		log.Fatal(err)
	}
	return http.ListenAndServe(httpAddr, mux)
}
