package main

import (
	"bar/lib"
	"fmt"
)

func main() {
	//	errors.New("foo")
	lib.Hello()
	LocalFunc()
}

func LocalFunc() {
	fmt.Println("local")
}
