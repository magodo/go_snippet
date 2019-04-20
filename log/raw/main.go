package main

import (
	"fmt"
	"runtime"
)

func Foo() {
	defer un(trace(get_func_name()))
	fmt.Println("Hello")
}

func get_func_name() string {
	pc := make([]uintptr, 2)
	n := runtime.Callers(2, pc)
	if n == 0 {
		return ""
	}
	frames := runtime.CallersFrames(pc[:n])
	frame, ok := frames.Next()
	if !ok {
		return ""
	}
	return frame.Function
}

func trace(name string) string {
	fmt.Println(name + ": ENTER")
	return name
}

func un(name string) {
	fmt.Println(name + ": LEAVE")
}

func main() {
	Foo()
}
