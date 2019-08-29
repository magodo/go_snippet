package main

import (
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		arr    []int
		sorted []int
	}{
		{
			arr:    []int{2, 4, 7, 2, 3, 9},
			sorted: []int{2, 2, 3, 4, 7, 9},
		},
		{
			arr:    []int{5, 4, 3, 2, 1},
			sorted: []int{1, 2, 3, 4, 5},
		},
		{
			arr:    []int{1},
			sorted: []int{1},
		},
		{
			arr:    []int{},
			sorted: []int{},
		},
	}

	for _, c := range cases {
		QSort(c.arr)
		if !reflect.DeepEqual(c.arr, c.sorted) {
			t.Fatalf("expect: %v\nactual: %v\n", c.sorted, c.arr)
		}
	}
}
