package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)

	zkConn, chEvt, err := zk.Connect(strings.Split("172.18.0.4:2181,172.18.0.2:2181,172.18.0.3:2181", ","), 3*time.Second) // NOTE: expiration timeout
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case evt := <-chEvt:
				log.Printf("[Default Watcher] %s", spew.Sdump(evt))
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

	go func() {
		evt := <-ch
		log.Printf("[Existence Watcher] %s", spew.Sdump(evt))
	}()

	log.Println("press Ctrl-C to quit...")
	<-c
	zkConn.Close()
}
