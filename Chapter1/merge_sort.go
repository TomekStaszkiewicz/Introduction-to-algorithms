package chapter1

// MergeSort function sorts an array of integers using merge algorith,
func MergeSort(A []int, p int, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSort(A, p, q)
	MergeSort(A, q+1, r)
	merge(A, p, q, r)
}

func merge(A []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q
	var leftTable []int
	var rightTable []int

	for i := 0; i < n1; i++ {
		leftTable = append(leftTable, A[p+i])
	}
	for i := 0; i < n2; i++ {
		rightTable = append(rightTable, A[q+1+i])
	}

	i := 0
	j := 0
	k := p
	for k <= r {
		if i >= n1 {
			A[k] = rightTable[j]
			j++
		} else if j >= n2 {
			A[k] = leftTable[i]
			i++
		} else if leftTable[i] <= rightTable[j] {
			A[k] = leftTable[i]
			i++
		} else {
			A[k] = rightTable[j]
			j++
		}
		k++
	}

}
