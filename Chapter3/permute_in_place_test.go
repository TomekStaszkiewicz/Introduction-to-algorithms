package chapter3

import (
	"reflect"
	"testing"
)

// PermuteInPlaceTest tests permuteInPlace function
func TestPermuteInPlace(t *testing.T) {
	tab := []int{0, 1, 2, 3, 4, 5, 1, 2, 9, 8}
	originalTab := make([]int, len(tab))

	copy(originalTab, tab)
	PermuteInPlace(tab)
	if reflect.DeepEqual(originalTab, tab) {
		t.Fatalf("Array not permutted!")
	}
	allElementsAfterpermutation := true

	for _, value := range originalTab {
		allElementsAfterpermutation = findInSlice(tab, value)
	}

	if !allElementsAfterpermutation {
		t.Fatalf("Array does not contain all the original elements!\nOriginal=%v permutted=%v", originalTab, tab)
	}
}

func findInSlice(tab []int, value int) bool {
	for _, v := range tab {
		if v == value {
			return true
		}
	}
	return false
}
