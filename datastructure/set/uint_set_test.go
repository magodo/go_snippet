package main

import "testing"

func TestFoo(t *testing.T) {
	set := NewUintSet(64)

	set.Add(0)
	if !set.Contains(0) {
		t.Fatal("Add 0 failed")
	}
	set.Add(63)
	if !set.Contains(63) {
		t.Fatal("Add 63 failed")
	}

	set.Clear(63)
	if set.Contains(63) {
		t.Fatal("Clear 63 failed")
	}
}
