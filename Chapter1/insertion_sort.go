package chapter1

// InsertionSort sorts given array of numbers
func InsertionSort(a []int) []int {
	var j int

	for iterator := range a {
		key := a[iterator]
		j = iterator - 1

		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}

	return a
}
