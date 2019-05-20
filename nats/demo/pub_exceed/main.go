package main

import (
	"fmt"
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// nats defaul outgoing buffer is fixed to be 32768
	// we configure the each payload to consume all buffer
	payload := ""
	for i := 0; i < 32768; i++ {
		payload += "x"
	}

	t := time.Tick(time.Millisecond * 10)
	for _ = range t {
		err := nc.Publish("foo", []byte(payload))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("send a message")
	}

	select {}
}
