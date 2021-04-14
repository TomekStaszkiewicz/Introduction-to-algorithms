package chapter2

import (
	"testing"
)

// Test basic example of max_sub_array algorithm
func TestMaxSubArray(t *testing.T) {

	testing_array := []int{0, -10, 10, 20}

	result := MaxSubArray(testing_array, 0, 3)

	if result.sum != 30 {
		t.Fatalf("Did not find correct subarray. Result sum = %v", result.sum)
	}
}

// TestWithSameNumbers tests MaxSubArray algorithm with the same numbers
func TestWithSameNumbers(t *testing.T) {

	testing_array := []int{-10, -10, -10, -10}

	result := MaxSubArray(testing_array, 0, 3)

	if result.sum != -10 {
		t.Fatalf("Did not find correct subarray. Result sum = %v", result.sum)
	}
}

// TestWithBigArray tests MaxSubArray with big array
func TestWithBigArray(t *testing.T) {
	testing_array := []int{0, 0, 0, 0, 0, 0, 10, -10, -100,
		100, 9, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 21, -30, 100}

	result := MaxSubArray(testing_array, 0, 26)

	if result.sum != 321 {
		t.Fatalf("Did not find correct subarray. Result sum = %v", result.sum)
	}
}
