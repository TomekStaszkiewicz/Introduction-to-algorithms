package chapter1

import (
	"sort"
	"testing"
)

// // TestCorrectSort calls merge_sort with array of numbers
func TestMergeSort(t *testing.T) {
	testing_array := []int{9, 8, 7, 100, 12}

	MergeSort(testing_array, 0, len(testing_array)-1)

	if !sort.IntsAreSorted(testing_array) {
		t.Fatalf("Array not sorted properly. Result array = %v", testing_array)
	}
}

// TestCorrectSortWithDuplication calls merge_sort with array of numbers
// with duplications
func TestMergeSortWithDuplication(t *testing.T) {
	testing_array := []int{9, 8, 7, 100, 12, 11, 12, 100, 9, 0, -1, 12}

	MergeSort(testing_array, 0, len(testing_array)-1)

	if !sort.IntsAreSorted(testing_array) {
		t.Fatalf("Array not sorted properly. Result array = %v", testing_array)
	}
}
