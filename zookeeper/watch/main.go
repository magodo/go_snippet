package main

import (
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	zkConn, chEvt, err := zk.Connect(strings.Split("172.18.0.4:2181,172.18.0.2:2181,172.18.0.3:2181", ","), 10*time.Second) // NOTE: expiration timeout
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case evt := <-chEvt:
				log.Printf("[Session Watcher] %s", spew.Sprint(evt))
			}
		}
	}()

	_, err = zkConn.Create("/foo", []byte{}, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatal(err)
	}

	_, _, ch, err := zkConn.ExistsW("/foo")
	if err != nil {
		log.Fatal(err)
	}

	select {
	case evt := <-ch:
		log.Printf("[Existence Watcher] %s", spew.Sdump(evt))
	}

	log.Println("press Ctrl-C to quit...")
	c := make(chan int)
	<-c
}
