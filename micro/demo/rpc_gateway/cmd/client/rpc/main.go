package main

import (
	"context"
	"fmt"
	"log"
	"time"

	rpc "github.com/magodo/rpc_gateway/internal/api/rpc"
	"github.com/micro/cli"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

var (
	serviceName string
)

func main() {
	service := grpc.NewService(
		micro.Name("magodo.service"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init(
		micro.Flags(cli.StringFlag{
			Name:        "service_name",
			Value:       "magodo.greeter",
			Destination: &serviceName,
		}),
	)

	service.Init()
	cli := rpc.NewSayService(serviceName, service.Client())
	resp, err := cli.Hello(context.Background(), &rpc.Request{
		Name: "magodo",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
