package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	counter := 0
	sub, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Received a message: %s\n", string(msg.Data))
		counter++
		if counter == 2 {
			wg.Done()
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	nc.Publish("foo", []byte("1"))
	nc.Publish("foo", []byte("1"))

	_ = sub
	//sub.Unsubscribe() // block
	nc.Drain()

	wg.Wait()
}
