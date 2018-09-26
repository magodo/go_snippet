package main

import (
	"fmt"
	"time"

	"github.com/looplab/fsm"
)

type Mover interface {
	MoveFromAToB() time.Duration
}

type Caterpillar bool

func (_ *Caterpillar) MoveFromAToB() time.Duration {
	return time.Hour
}

type Pupa bool

func (_ *Pupa) MoveFromAToB() time.Duration {
	return 1000 * time.Hour
}

type Butterfly bool

func (_ *Butterfly) MoveFromAToB() time.Duration {
	return time.Second
}

type Insect struct {
	Mover
	*fsm.FSM
}

func (insect *Insect) enterState(e *fsm.Event) {
	switch e.Event {
	case "upgrade":
		insect.Mover = new(Pupa)
	case "ultimate upgrade":
		insect.Mover = new(Butterfly)
	}
}

func NewInsect() *Insect {
	insect := &Insect{Mover: new(Caterpillar)}
	insect.FSM = fsm.NewFSM(
		"毛虫",
		fsm.Events{
			{Name: "upgrade", Src: []string{"毛虫"}, Dst: "蛹"},
			{Name: "ultimate upgrade", Src: []string{"蛹"}, Dst: "蝴蝶"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { insect.enterState(e) },
		},
	)
	return insect
}

func main() {
	insect := NewInsect()
	fmt.Printf("Move from A to B takes: %v\n", insect.MoveFromAToB())

	fmt.Println("Upgrade!!!")
	insect.Event("upgrade")
	fmt.Printf("Move from A to B takes: %v\n", insect.MoveFromAToB())

	fmt.Println("Ultimate upgrade!!!")
	insect.Event("ultimate upgrade")
	fmt.Printf("Move from A to B takes: %v\n", insect.MoveFromAToB())
}
