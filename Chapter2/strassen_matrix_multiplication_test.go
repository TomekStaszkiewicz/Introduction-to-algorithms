package chapter2

import (
	"reflect"
	"testing"
)

// TestBasicMultiplication calls Stassen matrix multiplication algorithm
func TestBasicMultiplication(t *testing.T) {
	m1 := MatrixData{
		matrix: [][]int{
			{0, 1, 2, 3},
			{5, 6, 7, 8},
			{-1, -1, -1, -1},
			{9, 8, 7, 6},
		},
	}
	m2 := MatrixData{
		matrix: [][]int{
			{1, -1, 2, 3},
			{2, 3, 0, 0},
			{1, 2, 3, 4},
			{9, 8, 7, 1},
		},
	}

	expected_value := [][]int{
		{31, 31, 27, 11},
		{96, 91, 87, 51},
		{-13, -12, -12, -8},
		{86, 77, 81, 61},
	}

	result := StrassenMatrixMultiplication(m1, m2)

	if !reflect.DeepEqual(result.matrix, expected_value) {
		t.Fatalf("Wrong array result = %v", result)
	}
}
