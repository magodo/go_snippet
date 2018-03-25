package main

import "time"

func foo() int {
	ch := make(chan int)
	go func() {
		select {
		case ch <- 1:
		default:
		}
	}()
	time.Sleep(time.Second)
	return <-ch
}

func main() {
	print(foo())
}
