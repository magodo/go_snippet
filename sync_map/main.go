package main

import (
	"sync"
	"time"
)

// plain map is not safe for concurrent access
// failed with error:
//
// 		fatal error: concurrent map read and map write
func plain_map() {

	m := map[int]int{1: 1}

	read := func() {
		_ = m[1]
	}

	write := func() {
		m[2] = 2
	}

	Read := func() {
		for {
			read()
		}
	}

	Write := func() {
		for {
			write()
		}
	}

	go Read()
	go Write()
	time.Sleep(time.Minute)
}

// use read-write mutex lock to protect concurrent access to plain map
func plain_map_with_lock() {
	lock := sync.RWMutex{}
	m := map[int]int{1: 1}

	read := func() {
		lock.RLock()
		defer lock.RUnlock()
		_ = m[1]
	}

	write := func() {
		lock.Lock()
		defer lock.Unlock()
		m[2] = 2
	}

	Read := func() {
		for {
			read()
		}
	}

	Write := func() {
		for {
			write()
		}
	}

	go Read()
	go Write()
	time.Sleep(time.Minute)
}

// sync.Map is optimiazed type to avoid using lock, with following features:
// 	(1) when the entry for a given key is only ever written once but read many times, as in caches that only grow
// 	(2) when multiple goroutines read, write, and overwrite entries for disjoint sets of keys.
// In these two cases, use of a Map may significantly reduce lock contention compared to a Go map paired with a separate Mutex or RWMutex.
func sync_map() {

	m := sync.Map{}
	m.Store(1, 1)

	read := func() {
		_, _ = m.Load(1)
	}

	write := func() {
		m.Store(2, 2)
	}

	Read := func() {
		for {
			read()
		}
	}

	Write := func() {
		for {
			write()
		}
	}

	go Read()
	go Write()
	time.Sleep(time.Minute)
}

// plain map holding sync map as value, write is done only once for plain map, at very begining
// later, the read and writes are only for the embedded sync.Map
func plain_map_embed_sync_map() {
	m := map[string]*sync.Map{
		"a": &sync.Map{},
	}

	m["a"].Store(1, 1)

	read := func() {
		_, _ = m["a"].Load(1)
	}

	write := func() {
		m["a"].Store(2, 2)
	}

	Read := func() {
		for {
			read()
		}
	}

	Write := func() {
		for {
			write()
		}
	}
	go Read()
	go Write()
	time.Sleep(time.Minute)

}

func main() {
	//plain_map()
	//plain_map_with_lock()
	//sync_map()
	//plain_map_embed_sync_map()
}
