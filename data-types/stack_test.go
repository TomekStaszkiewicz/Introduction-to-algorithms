package dataTypes

import (
	"testing"
)

func TestCreate(t *testing.T) {
	stack := Create()

	if len(*stack) != 0 {
		t.Fatalf("Stack not created properly!")
	}
}

func TestPush(t *testing.T) {
	stack := Create()

	stack.Push(10)
	stack.Push(20)

	if len(*stack) != 2 {
		t.Fatalf("Incorrect push operation, len is: %v", len(*stack))
	}
}

func TestPop(t *testing.T) {
	stack := Create()

	stack.Push(10)
	stack.Push(20)
	val, _ := stack.Pop()
	if val != 20 {
		t.Fatalf("Expected 20, got %v in stack.Pop", val)
	}

	val, _ = stack.Pop()
	if val != 10 {
		t.Fatalf("Expected 10, got %v in stack.Pop", val)
	}

	val, err := stack.Pop()
	if err == nil {
		t.Fatalf("Pop did not throw error on empty stack")
	}
}
