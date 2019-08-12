package stack

import "testing"

func TestFoo(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Fatalf("expect size to be %d, but actual %d", 3, stack.Size())
	}

	if stack.Peek() != 3 {
		t.Fatalf("expect peek to be %d, but actual %d", 3, stack.Peek())
	}

	stack.Pop()
	top := stack.Pop()
	if top != 2 {
		t.Fatalf("expect top to be %d, but actual %d", 2, top)
	}
	stack.Pop()
	top = stack.Pop()
	if top != nil {
		t.Fatalf("expect top to be %v, but actual %v", nil, top)
	}
}
