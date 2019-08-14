package main

import "testing"

func TestDefaultDict(t *testing.T) {
	d := NewDefaultDict(0)
	if d.Get(0).(int) != 0 {
		t.Fatal("default value is not set")
	}
	d.Set(0, 1)
	if d.Get(0).(int) != 1 {
		t.Fatal("set value is not set")
	}
	d.Delete(0)
	if d.Get(0).(int) != 0 {
		t.Fatal("default value is not set (after restore)")
	}
}
