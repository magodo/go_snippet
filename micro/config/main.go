//+build ignore

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
)

type BazConf struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Unit string `json:"unit"`
}

func main() {
	consulSource := consul.NewSource(
		consul.WithAddress("localhost:8500"),
		consul.WithPrefix("foo/bar"),
		consul.StripPrefix(true),
	)

	conf := config.NewConfig(
		config.WithSource(consulSource),
	)

	var baz BazConf
	for {
		if err := conf.Get("baz").Scan(&baz); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", baz)
		time.Sleep(time.Second)
	}
}
