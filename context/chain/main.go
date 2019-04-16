package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//ctx1, _ := context.WithCancel(ctx)
	go func() {
		<-ctx.Done()
		fmt.Println("gr1 canceld")
	}()

	//ctx2, _ := context.WithCancel(ctx)
	go func() {
		<-ctx.Done()
		fmt.Println("gr2 canceld")
	}()

	fmt.Println("press 'enter' to cancel...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

	time.Sleep(time.Second)
}
