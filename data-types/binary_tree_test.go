package dataTypes

import (
	"testing"
)

func TestCreateTree(t *testing.T) {
	tree := CreateTree(10)

	if tree.value != 10 || tree.leftChild != nil || tree.rightChild != nil {
		t.Errorf("Tree created incorrectly")
	}
}

func TestAddNodeToTree(t *testing.T) {
	tree := CreateTree(10)

	tree.Add(11)
	tree.Add(9)

	tree.Print()
}
