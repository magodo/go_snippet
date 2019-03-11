package main

import "errors"

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("no element")
	}
	return s.values[0], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Push(x int) {
	s.values = append([]int{x}, s.values...)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("no element")
	}
	x := s.values[0]
	s.values = s.values[1:]
	return x, nil
}
