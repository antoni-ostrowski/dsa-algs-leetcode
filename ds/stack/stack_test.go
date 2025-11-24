package stack

import "testing"

func TestStack(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	if v, _ := stack.Pop(); v != 2 {
		t.Fatalf("Pop: expected 2, got %v", v)
	}

}
