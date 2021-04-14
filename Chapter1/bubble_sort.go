package chapter1

// BubbleSort sorts array of integers using bubble sort algorithm
func BubbleSort(A []int) {

	for i := range A {
		for j := 0; j < len(A)-1-i; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}

}
