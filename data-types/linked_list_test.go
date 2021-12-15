package dataTypes

import (
	"testing"
)

func TestCreateList(t *testing.T) {
	list := CreateList()

	if list.Head() != nil {
		t.Errorf("List created incorrectly")
	}
}

func TestAdd(t *testing.T) {
	list := CreateList()

	list.Add(11)

	head := list.Head()

	if head.Value != 11 {
		t.Errorf("Error when adding element. Expected 11, got %v", head.Value)
	}
}
