package main

import (
	"log"
	"github.com/magodo/go_snippet/grpc_todo/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		log.Fatal(err)
	}
}