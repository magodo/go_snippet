package main

import (
	"log"

	"foo/internal"
	pb "foo/proto/greeter"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

const (
	serviceName    string = "greeter"
	serviceVersion string = "v0.0.0"
)

func main() {

	// prepare service
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
	)
	name := ""
	srv.Init(
		micro.Flags(
			cli.StringFlag{
				Name:   "greeter_name",
				Usage:  "greeter name",
				EnvVar: "GREETER_NAME",
			},
		),
		micro.Action(func(c *cli.Context) {
			if name = c.String("greeter_name"); name == "" {
				log.Fatal("greeter_name not specified")
			}
		}),
	)
	pb.RegisterGreeterServiceHandler(srv.Server(), &internal.Handler{Name: name})

	// run
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
