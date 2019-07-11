// +build ignore

package main

import (
	"cache"
	"cache/memcached"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	c := memcached.NewMemcachedClient("0.0.0.0:32769")
	c.Set(&cache.Record{Key: "bar", Value: nil, Expiry: 5 * time.Second})
	time.Sleep(2)
	rec, err := c.Get("bar")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(rec)
}
