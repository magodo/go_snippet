package main

import "testing"

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	x, _ := s.Peek()
	if x != 2 {
		t.Fatal("peek failed")
	}

	x, _ = s.Pop()
	if x != 2 {
		t.Fatal("pop failed")
	}

	if !testEq(s.values, []int{1}) {
		t.Fatal("pop result failed")
	}
}
