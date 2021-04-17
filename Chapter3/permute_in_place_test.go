package chapter3

import (
	"reflect"
	"testing"
)

// PermuteInPlaceTest tests permuteInPlace function
func TestPermuteInPlace(t *testing.T) {
	tab := []int{0, 1, 2, 3, 4}
	originalTab := make([]int, len(tab))

	copy(originalTab, tab)
	PermuteInPlace(tab)
	if reflect.DeepEqual(originalTab, tab) {
		t.Fatalf("Array not permutted!")
	}
}
