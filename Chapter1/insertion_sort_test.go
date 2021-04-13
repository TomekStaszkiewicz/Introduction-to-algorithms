package chapter1

import (
	"sort"
	"testing"
)

// TestCorrectSort calls insertion_sort with array of numbers
func TestInsertionSort(t *testing.T) {
	testing_array := []int{9, 8, 7, 100, 12}

	result_array := InsertionSort(testing_array)

	if !sort.IntsAreSorted(result_array) {
		t.Fatalf("Array not sorted properly. Result array = %v", result_array)
	}
}

// TestCorrectSortWithDuplication calls insertion_sort with array of numbers
// with duplications
func TestInsertionSortWithDuplication(t *testing.T) {
	testing_array := []int{9, 8, 7, 100, 12, 12, 100, 9, 0, -1, 12}

	result_array := InsertionSort(testing_array)

	if !sort.IntsAreSorted(result_array) {
		t.Fatalf("Array not sorted properly. Result array = %v", result_array)
	}
}
