package dataTypes

import "fmt"

type TreeNode struct {
	value      int
	leftChild  *TreeNode
	rightChild *TreeNode
}

func CreateTree(value int) *TreeNode {
	return &TreeNode{value: value}
}

func (parentNode *TreeNode) Add(value int) {
	if value < parentNode.value {
		if parentNode.leftChild == nil {
			parentNode.leftChild = &TreeNode{value: value}
		} else {
			parentNode.leftChild.Add(value)
		}
	} else {
		if parentNode.rightChild == nil {
			parentNode.rightChild = &TreeNode{value: value}
		} else {
			parentNode.rightChild.Add(value)
		}
	}
}

func (node *TreeNode) Print() {
	fmt.Println(node.value)
	if node.leftChild != nil {
		node.leftChild.Print()
	}
	if node.rightChild != nil {
		node.rightChild.Print()
	}
}
