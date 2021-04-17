package chapter3

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	tab := []int{0, 1, 2, 3, 4}

	if reflect.DeepEqual(tab, SortingPermutation(tab)) {
		t.Fatalf("Array not permutted!")
	}
}
