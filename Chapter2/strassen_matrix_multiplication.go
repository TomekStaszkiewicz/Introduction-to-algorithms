package chapter2

// StrassenMatrixMultiplication uses Strassen algorithm to multiply two
// square matrixes
func StrassenMatrixMultiplication(A MatrixData, B MatrixData) MatrixData {

	if len(A.matrix) == 1 {
		return MatrixData{
			matrix: [][]int{
				{A.matrix[0][0] * B.matrix[0][0]},
			},
		}
	}
	C := createEmptyMatrix(len(A.matrix))
	A_11, A_12, A_21, A_22 := A.getDividedArray()
	B_11, B_12, B_21, B_22 := B.getDividedArray()
	S_1 := B_12.substract(B_22)
	S_2 := A_11.add(A_12)
	S_3 := A_21.add(A_22)
	S_4 := B_21.substract(B_11)
	S_5 := A_11.add(A_22)
	S_6 := B_11.add(B_22)
	S_7 := A_12.substract(A_22)
	S_8 := B_21.add(B_22)
	S_9 := A_11.substract(A_21)
	S_10 := B_11.add(B_12)

	P_1 := StrassenMatrixMultiplication(A_11, S_1)
	P_2 := StrassenMatrixMultiplication(S_2, B_22)
	P_3 := StrassenMatrixMultiplication(S_3, B_11)
	P_4 := StrassenMatrixMultiplication(A_22, S_4)
	P_5 := StrassenMatrixMultiplication(S_5, S_6)
	P_6 := StrassenMatrixMultiplication(S_7, S_8)
	P_7 := StrassenMatrixMultiplication(S_9, S_10)

	C_11 := P_5.add(P_4).substract(P_2).add(P_6)
	C_12 := P_1.add(P_2)
	C_21 := P_3.add(P_4)
	C_22 := P_5.add(P_1).substract(P_3).substract(P_7)

	C = merge_arrays(C_11, C_12, C_21, C_22)

	return C
}

func createEmptyMatrix(size int) MatrixData {
	arr := [][]int{}
	for i := 0; i < size; i++ {
		newRow := []int{}
		for j := 0; j < size; j++ {
			newRow = append(newRow, 0)
		}
		arr = append(arr, newRow)
	}

	return MatrixData{
		matrix: arr,
	}
}

type MatrixData struct {
	matrix [][]int
}

func (m MatrixData) multiply(A MatrixData) MatrixData {
	C := createEmptyMatrix(len(A.matrix))

	for i := range m.matrix {
		for j := range A.matrix {
			C.matrix[i][j] = 0
			for k := range A.matrix {
				C.matrix[i][j] += m.matrix[i][k] * A.matrix[k][j]
			}
		}
	}

	return C
}

func (m MatrixData) add(A MatrixData) MatrixData {
	C := createEmptyMatrix(len(A.matrix))

	for i := range m.matrix {
		for j := range A.matrix {
			C.matrix[i][j] = m.matrix[i][j] + A.matrix[i][j]
		}
	}

	return C
}

func (m MatrixData) substract(A MatrixData) MatrixData {
	C := createEmptyMatrix(len(A.matrix))

	for i := range m.matrix {
		for j := range A.matrix {
			C.matrix[i][j] = m.matrix[i][j] - A.matrix[i][j]
		}
	}

	return C
}

func (m MatrixData) getDividedArray() (MatrixData, MatrixData, MatrixData, MatrixData) {
	orgLength := len(m.matrix)
	size := orgLength / 2
	A_11 := createEmptyMatrix(size)
	A_12 := createEmptyMatrix(size)
	A_21 := createEmptyMatrix(size)
	A_22 := createEmptyMatrix(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A_11.matrix[i][j] = m.matrix[i][j]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A_12.matrix[i][j] = m.matrix[i][j+size]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A_21.matrix[i][j] = m.matrix[size+i][j]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A_22.matrix[i][j] = m.matrix[size+i][size+j]
		}
	}

	return A_11, A_12, A_21, A_22
}

func merge_arrays(
	A_11 MatrixData,
	A_12 MatrixData,
	A_21 MatrixData,
	A_22 MatrixData,
) MatrixData {
	size := len(A_11.matrix) * 2
	C := createEmptyMatrix(size)

	A_11_size := len(A_11.matrix)
	for i := 0; i < A_11_size; i++ {
		for j := 0; j < A_11_size; j++ {
			C.matrix[i][j] = A_11.matrix[i][j]
		}
	}

	A_12_size := len(A_12.matrix)
	for i := 0; i < A_12_size; i++ {
		for j := 0; j < A_12_size; j++ {
			C.matrix[i][j+A_11_size] = A_12.matrix[i][j]
		}
	}
	A_21_size := len(A_21.matrix)
	for i := 0; i < A_11_size; i++ {
		for j := 0; j < A_21_size; j++ {
			C.matrix[i+A_11_size][j] = A_21.matrix[i][j]
		}
	}
	A_22_size := len(A_22.matrix)
	for i := 0; i < A_22_size; i++ {
		for j := 0; j < A_22_size; j++ {
			C.matrix[i+A_11_size][j+A_11_size] = A_22.matrix[i][j]
		}
	}

	return C
}
