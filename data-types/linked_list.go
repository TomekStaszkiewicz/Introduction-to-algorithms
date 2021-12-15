package dataTypes

import "fmt"

type LinkedListItem struct {
	Value int
	next  *LinkedListItem
}

type LinkedList struct {
	head *LinkedListItem
}

func CreateList() *LinkedList {
	list := LinkedList{}

	return &list
}

func (l *LinkedList) GetLength() int {
	length := 0
	currentNode := l.head
	for currentNode.next != nil {
		length += 1
		currentNode = currentNode.next
	}

	return length
}

func (l *LinkedList) Add(val int) {
	item := LinkedListItem{
		Value: val,
	}
	if l.head == nil {
		l.head = &item
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &item
	}
}

func (l *LinkedList) Head() *LinkedListItem {
	return l.head
}

func (l *LinkedList) PrintValues() {
	currNode := l.head

	for currNode != nil {
		fmt.Println(currNode.Value)
		currNode = currNode.next
	}
}
