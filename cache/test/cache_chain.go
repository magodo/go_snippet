package main

import (
	"cache"
	"cache/memcached"
	"cache/memory"
	"cache/redis"
	"log"
	"time"
)

func main() {
	cmemcached := memcached.NewCache("0.0.0.0:32770")
	if cmemcached == nil {
		log.Fatal("failed to new memcached")
	}
	credis := redis.NewCache("0.0.0.0:32768", "")
	if credis == nil {
		log.Fatal("failed to new redis")
	}
	cmem := memory.NewCache()

	chains := cache.ChainCache(cmem, credis, cmemcached)

	record := &cache.Record{
		Key:    "foo",
		Value:  nil,
		Expiry: 10 * time.Second,
	}

	if err := cmemcached.Set(record); err != nil {
		log.Fatal(err)
	}

	if _, err := chains.Get("foo"); err != nil {
		log.Fatal(err)
	}

	if _, err := cmem.Get("foo"); err != nil {
		log.Fatal("memory ", err)
	}
	if _, err := credis.Get("foo"); err != nil {
		log.Fatal("redis ", err)
	}
}
