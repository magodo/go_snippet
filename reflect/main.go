package main

import (
	"errors"
	"log"
	"reflect"
)

func MakeAsyncFunction(ch chan error, f interface{}) interface{} {
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(
		vf.Type(),
		func(in []reflect.Value) []reflect.Value {
			go func() {
				out := vf.Call(in)
				var err error
				err, _ = out[0].Interface().(error) // nil interface can't do type assertion
				ch <- err
			}()

			// tricky way to construct `Value` with (value, type) of (nil, error)
			// see: https://stackoverflow.com/questions/51092352/how-can-i-instantiate-a-nil-error-using-golangs-reflect
			return []reflect.Value{reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())}
		},
	)
	return wrapperF.Interface()
}

func main() {
	c := make(chan error, 1)

	f := MakeAsyncFunction(c, foo).(func(int) error)
	err := f(1)
	if err != nil {
		log.Fatal(err)
	}
	err = <-c
	if err != nil {
		log.Fatal(err)
	}

	b := MakeAsyncFunction(c, bar).(func(int, int) error)
	err = b(1, 2)
	if err != nil {
		log.Fatal(err)
	}
	err = <-c
	if err != nil {
		log.Fatal(err)
	}
}

func foo(i int) error {
	log.Printf("foo: %d\n", i)
	return nil

}

func bar(i int, j int) error {
	log.Printf("bar: %d, %d\n", i, j)
	return errors.New("bar error")

}
