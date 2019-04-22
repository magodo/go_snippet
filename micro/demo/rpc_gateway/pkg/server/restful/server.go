package restful

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	greeter "github.com/magodo/rpc_gateway/internal/api/restful"
	"google.golang.org/grpc"
)

func Serve(endpoint string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := greeter.RegisterSayHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err = http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
