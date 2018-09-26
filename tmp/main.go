package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var r = bufio.NewReader(os.Stdin)

func askUserEnter(msg string, out interface{}) {
	fmt.Print(msg)
	input, _ := r.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	switch out.(type) {
	case *string:
		*(out.(*string)) = input
	case *int:
		num, _ := strconv.Atoi(input)
		*(out.(*int)) = num
	case *uint64:
		num, _ := strconv.ParseUint(input, 10, 64)
		*(out.(*uint64)) = num
	}
}

func ChooseDBTypeID() (typeid int) {
	askUserEnter(`Choose DB type below: 
    1: 		mysql-5.5
    2: 		mysql-5.1
    6: 		mysql-5.6
    10:     mysql-5.7
    16:		postgresql-9.6
Enter index: `, &typeid)
	return
}

func ChooseHaInstanceID() (id string) {
	askUserEnter("Enter HA instance ID (len < 36): ", &id)
	return
}

func ChooseCPULimit() (limit uint64) {
	askUserEnter("Enter CPU core amount limit: ", &limit)
	return
}

type Do interface {
	Do()
}

func Logic(d Do) {
	d.Do()
}

type Derived1 struct {
}
type Derived2 struct {
}

func (d *Derived1) Do() {
	fmt.Println("1")
}
func (d *Derived2) Do() {
	fmt.Println("2")
}

func a(i int) {
	fmt.Println(i)
}

type Set struct {
	sync.Mutex
	m map[interface{}]bool
}

//
func (s *Set) Add(obj interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[obj] = true
}

// TryAdd trys to add an obj into set, if already exists, returns false.
func (s *Set) TryAdd(obj interface{}) (ok bool) {
	s.Lock()
	defer s.Unlock()
	_, exists := s.m[obj]
	if exists {
		return false
	}
	s.m[obj] = true
	return true
}

//
func (s *Set) Delete(obj interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, obj)
}

//
func (s *Set) Contains(obj interface{}) bool {
	s.Lock()
	defer s.Unlock()
	_, exists := s.m[obj]
	return exists
}

var set = &Set{
	m: make(map[interface{}]bool),
}

type Callable struct {
	i int
	f func()
}

func NewCallable() (c *Callable) {

	f := func() {
		fmt.Println(c.i)
	}
	return &Callable{
		i: 123,
		f: f,
	}
}

type Foo struct {
	I int
	D float64
}

func main() {
	f := &Foo{
		1,
		1.0,
	}

	str, _ := json.MarshalIndent(f, "", "  ")
	fmt.Println(string(str))
}
