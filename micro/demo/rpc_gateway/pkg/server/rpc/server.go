package rpc

import (
	"log"
	"time"

	rpc "github.com/magodo/rpc_gateway/internal/api/rpc"
	srv "github.com/magodo/rpc_gateway/internal/service"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

func Serve() {
	service := grpc.NewService(
		micro.Name("magodo.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	rpc.RegisterSayHandler(service.Server(), new(srv.Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
