package wgo

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Wait will wait for all outstanding goroutines to finish, unless "d" timeout(and this func returns false)
// If "d" is 0 (i.e. time.Duration(0)), then it will wait forever.
func Wait(d time.Duration) bool {
	if d == time.Duration(0) {
		wg.Wait()
		return true
	}
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return true
	case <-time.After(d):
		return false
	}
}

// Launch just launches a go routine, except the go routine is tracked by the global wg.
func Launch(f func()) {
	log.Println("add 1")
	wg.Add(1)
	go func() {
		defer func() {
			log.Println("done")
			wg.Done()
		}()
		f()
	}()
}
