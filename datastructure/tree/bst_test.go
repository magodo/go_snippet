package tree

import (
	"fmt"
	"log"
	"testing"
)

func TestBst(t *testing.T) {
	tree := new(Bst)
	tree.RegisterCompFunc(func(a interface{}, b interface{}) int {
		aa, ok := a.(int)
		if !ok {
			panic(fmt.Sprintf("%v is not int", a))
		}
		bb, ok := b.(int)
		if !ok {
			panic(fmt.Sprintf("%v is not int", b))
		}

		switch {
		case aa > bb:
			return 1
		case aa < bb:
			return -1
		default:
			return 0
		}
	})

	tree.Insert(4)

	tree.Insert(2)
	tree.Insert(6)

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)

	l := []int{}
	tree.Traverse(func(node *Node) {
		v, _ := node.Value.(int)
		l = append(l, v)
	})
	for i := 0; i < 7; i++ {
		if l[i] != i+1 {
			t.Fatal("insert not work as expected")
		}
	}

	node, _ := tree.Find(5)
	if node.Value != 5 {
		t.Fatal("find not work as expected")
	}

	tree.Traverse(func(node *Node) {
		fmt.Printf("find nearest ascendant for %v: ", node.Value)
		n, err := node.FindNearestAscendant()
		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			fmt.Printf("%v", n.Value)
		}
		fmt.Println()
	})

	tree.Traverse(func(node *Node) {
		fmt.Printf("find nearest descendant for %v: ", node.Value)
		n, err := node.FindNearestDescendant()
		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			fmt.Printf("%v", n.Value)
		}
		fmt.Println()
	})

	n, err := tree.FindNearestDescendantNodeAgainstValue(3)
	if err != nil {
		log.Fatal(err)
	}
	if n.Value.(int) != 2 {
		t.Fatal("find nearest descendant node against 3 failed")
	}

	n, err = tree.FindNearestAscendantNodeAgainstValue(3)
	if err != nil {
		log.Fatal(err)
	}
	if n.Value.(int) != 4 {
		t.Fatal("find nearest descendant node against 3 failed")
	}

	tree.Traverse(func(node *Node) {
		v, _ := node.Value.(int)
		fmt.Println(v)
	})

	tree.Delete(1)
	tree.Delete(2)
	tree.Delete(3)
	tree.Delete(4)
	tree.Delete(6)
	tree.Delete(7)
	if !(tree.Root.Value == interface{}(5) && tree.Root.Left == nil && tree.Root.Right == nil) {
		t.Fatal("delete not work as expected")
	}
}
