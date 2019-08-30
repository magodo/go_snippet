package array

import (
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		grid       [][]int
		fill       int
		targetGrid [][]int
	}{
		{
			[][]int{
				[]int{1, 1},
				[]int{1, 1},
			},
			0,
			[][]int{
				[]int{0, 0, 0},
				[]int{0, 1, 1},
				[]int{0, 1, 1},
			},
		},
	}

	for _, c := range cases {
		out := Expand(c.grid, c.fill)
		if !reflect.DeepEqual(out, c.targetGrid) {
			t.Fatalf("expect: %v\ngot: %v\n", c.targetGrid, out)
		}
	}
}
