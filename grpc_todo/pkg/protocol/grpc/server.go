package grpc

import (
	"context"
	"github.com/magodo/go_snippet/grpc_todo/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

// RunServer run grpc server
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register services
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutdown on user interrupt")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(l)
}
