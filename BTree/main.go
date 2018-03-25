package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	n1, n2 := 0, 0
	var num int
	var buf1, buf2 []int

	for n1 < 10 || n2 < 10 {
		select {
		case num = <-ch1:
			buf1 = append(buf1, num)
			n1 += 1
		case num = <-ch2:
			buf2 = append(buf2, num)
			n2 += 1
		}
	}

	fmt.Print(buf1)
	fmt.Print(buf2)

	// check equality
	if len(buf1) == len(buf2) {
		for i := 0; i < len(buf1); i++ {
			if buf1[i] != buf2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
