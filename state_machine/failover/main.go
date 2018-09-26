package main

import (
	"fmt"

	"github.com/looplab/fsm"
	"github.com/pkg/errors"
)

// State represents each state during failover
type State interface {
	MonitorState(sm *StateMachine)
	String() string
}

// Actor abstract all actions during failover, meant to be implemented by different DB
type Actor interface {
	ActNormal()
	ActPreFailover()
	ActPostFailover()
}

type DummyActor struct{}

func (a *DummyActor) ActNormal() {
	fmt.Println("normal")
}
func (a *DummyActor) ActPreFailover() {
	fmt.Println("pre failover")
}
func (a *DummyActor) ActPostFailover() {
	fmt.Println("post failover")
}

type StateMachine struct {
	Actor
	State
	*fsm.FSM
}

func NewStateMachine() (sm *StateMachine) {
	context := fsm.NewFSM(
		"normal",
		fsm.Events{
			{Name: "master_down", Src: []string{"normal"}, Dst: "pre_failover"},
			{Name: "master_up", Src: []string{"pre_failover"}, Dst: "normal"},
			{Name: "failover", Src: []string{"pre_failover"}, Dst: "post_failover"},
			{Name: "slave_up", Src: []string{"post_failover"}, Dst: "post_normal"},
		},
		fsm.Callbacks{
			"before_event": func(e *fsm.Event) {
				var err error
				switch e.Event {
				case "master_down":
					err = sm.TransitFromNormalToPreFailover()
				case "master_up":
					err = sm.TransitFromPreFailoverToNormal()
				case "failover":
					err = sm.TransitFromPreFailoverToPostFailover()
				case "slave_up":
					err = sm.TransitFromPostFailoverToNormal()
				}
				if err != nil {
					err = errors.Wrapf(err, "state transition failed(triggered by %s)", err)
					e.Cancel(err)
				}
			},
			"enter_state": func(e *fsm.Event) {
				switch e.Event {
				case "master_down":
					sm.State = new(StatePreFailover)
				case "master_up":
					sm.State = new(StateNormal)
				case "failover":
					sm.State = new(StatePostFailover)
				case "slave_up":
					sm.State = new(StateNormal)
				}
			},
		},
	)

	return &StateMachine{
		Actor: &DummyActor{},
		State: new(StateNormal),
		FSM:   context,
	}
}

/////////////////////// STATES /////////////////////////////////////

// NORMAL STATE
type StateNormal bool

func (s *StateNormal) MonitorState(sm *StateMachine) {
	sm.Actor.ActNormal()

	err := sm.Event("master_down")
	if err != nil {
		err = errors.Wrap(err, "failed to submit event: master_down")
		fmt.Println(err)
	}
	return
}

func (s *StateNormal) String() string {
	return "normal"
}

// PRE_FAILOVER STATE
type StatePreFailover bool

func (s *StatePreFailover) MonitorState(sm *StateMachine) {
	sm.Actor.ActPreFailover()

	err := sm.Event("failover")
	if err != nil {
		err = errors.Wrap(err, "failed to submit event: failover")
		fmt.Println(err)
	}
	return
}

func (s *StatePreFailover) String() string {
	return "pre_failover"
}

// POST_FAILOVER STATE
type StatePostFailover bool

func (s *StatePostFailover) MonitorState(sm *StateMachine) {
	sm.Actor.ActPostFailover()

	err := sm.Event("slave_up")
	if err != nil {
		err = errors.Wrap(err, "failed to submit event: slave_up")
		fmt.Println(err)
	}
	return
}

func (s *StatePostFailover) String() string {
	return "post_failover"
}

/////////////////////// TRANSITIONS /////////////////////////////////////

func (sm *StateMachine) TransitFromNormalToPreFailover() (err error) {
	return nil
}

func (sm *StateMachine) TransitFromPreFailoverToNormal() (err error) {
	return nil
}

func (sm *StateMachine) TransitFromPreFailoverToPostFailover() (err error) {
	return nil
}

func (sm *StateMachine) TransitFromPostFailoverToNormal() (err error) {
	return nil
}

func main() {
	sm := NewStateMachine()
	sm.State.MonitorState(sm)
	sm.State.MonitorState(sm)
	sm.State.MonitorState(sm)
}
