package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/looplab/fsm"
)

type Foo interface {
	Bar()
}

type SM struct {
	Foo
	*fsm.FSM
}

type S1 bool

func (_ *S1) Bar() {
	fmt.Println("In S1")
}

type S2 bool

func (_ *S2) Bar() {
	fmt.Println("In S2")
}

type S3 bool

func (_ *S3) Bar() {
	fmt.Println("In S3")
}

func NewSM() *SM {
	sm := &SM{Foo: new(S1)}
	sm.FSM = fsm.NewFSM(
		"S1",
		fsm.Events{
			{Name: "12", Src: []string{"S1"}, Dst: "S2"},
			{Name: "23", Src: []string{"S2"}, Dst: "S3"},
			{Name: "21", Src: []string{"S2"}, Dst: "S1"},
		},
		fsm.Callbacks{
			"before_event": func(e *fsm.Event) {
				fmt.Println("before_event called")
				switch e.Event {
				case "23":
					e.Cancel(errors.New("foo"))
				}
				//spew.Dump(e)
			},
			"after_event": func(e *fsm.Event) {
				fmt.Println("after_event called")
				//spew.Dump(e)
			},
			"enter_state": func(e *fsm.Event) {
				fmt.Println("enter_state called")
				switch e.Event {
				case "12":
					sm.Foo = new(S2)
				case "23":
					sm.Foo = new(S3)
				case "21":
					sm.Foo = new(S1)
				}
			},
		},
	)
	return sm
}

func main() {
	sm := NewSM()
	sm.Bar()

	fmt.Println("1->2")
	err := sm.Event("12")
	if err != nil {
		log.Println(err)
	}

	sm.Bar()

	fmt.Println("2->3")
	err = sm.Event("23")
	if err != nil {
		log.Println(err)
	}
	sm.Bar()

	fmt.Println("2->1")
	err = sm.Event("21")
	if err != nil {
		log.Println(err)
	}
	sm.Bar()
}
