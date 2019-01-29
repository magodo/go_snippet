package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// gen() generate some numbers and feeds to channel(unbuffered)
func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				switch ctx.Err() {
				case context.Canceled:
					fmt.Println("gen: canceld")
				case context.DeadlineExceeded:
					fmt.Println("gen: timeout")
				}
				return
			}
		}
	}()
	return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-ctx.Done():
				switch ctx.Err() {
				case context.Canceled:
					fmt.Println("sq: canceld")
				case context.DeadlineExceeded:
					fmt.Println("sq: timeout")
				}
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-ctx.Done():
				switch ctx.Err() {
				case context.Canceled:
					fmt.Println("merge: canceld")
				case context.DeadlineExceeded:
					fmt.Println("merge: timeout")
				}
				return
			}
		}
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
func mergeBuggy(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			//time.Sleep(time.Second) // NOTE: this will incur sq() range loop run into "default" branch
			out <- n
		}
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	in := gen(ctx, 2, 3, 4, 5)

	// fan-out
	c1 := sq(ctx, in)
	c2 := sq(ctx, in)

	// fan-in

	out := merge(ctx, c1, c2)
	//out := mergeBuggy(c1, c2)
	for n := range out {
		fmt.Println(n)
	}

	//fmt.Println(<-out)

	// ctx will be canceled
	return
}
