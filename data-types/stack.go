package dataTypes

import (
	"errors"
)

type Stack []int

func CreateStack() *Stack {
	x := Stack{}
	return &x
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("stack is empty")
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return val, nil
}
