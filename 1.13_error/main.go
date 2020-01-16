package main

import (
	"errors"
	"fmt"
	"log"
)

// Define error, need implement two methods:
// - Unwrap()
// - Error()
type MyError struct {
	Err error
}

func (e *MyError) Unwrap() error {
	return e.Err
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %s", e.Err.Error())
}

var MyErrorFoo = &MyError{Err: errors.New("foo")}

func main() {
	e := fmt.Errorf("some: %w", MyErrorFoo)
	if errors.Is(e, MyErrorFoo) {
		log.Println("yes")
	}
}
