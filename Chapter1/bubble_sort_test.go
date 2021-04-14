package chapter1

import (
	"sort"
	"testing"
)

// // TestCorrectSort calls merge_sort with array of numbers
func TestBubbleSort(t *testing.T) {
	testing_array := []int{9, 8, 7, 100, 12}

	BubbleSort(testing_array)

	if !sort.IntsAreSorted(testing_array) {
		t.Fatalf("Array not sorted properly. Result array = %v", testing_array)
	}
}

// TestCorrectSortWithDuplication calls merge_sort with array of numbers
// with duplications
func TestBubbleSortWithDuplication(t *testing.T) {
	testing_array := []int{9, 8, 7, 100, 12, 11, 12, 100, 9, 0, -1, 12}

	BubbleSort(testing_array)

	if !sort.IntsAreSorted(testing_array) {
		t.Fatalf("Array not sorted properly. Result array = %v", testing_array)
	}
}
