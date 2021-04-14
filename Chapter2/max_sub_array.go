package chapter2

import (
	"math"
)

type SubArrResult struct {
	leftIndex  int
	rightIndex int
	sum        int
}

func MaxSubArray(A []int, l int, h int) SubArrResult {
	if l == h {
		return SubArrResult{
			leftIndex:  l,
			rightIndex: h,
			sum:        A[l],
		}
	} else {
		m := (l + h) / 2
		resLeft := MaxSubArray(A, l, m)
		resRight := MaxSubArray(A, m+1, h)
		resMiddle := findMiddleMaxSubArray(A, l, m, h)

		if resLeft.sum >= resRight.sum && resLeft.sum >= resMiddle.sum {
			return resLeft
		} else if resRight.sum >= resLeft.sum && resRight.sum >= resMiddle.sum {
			return resRight
		} else {
			return resMiddle
		}
	}
}

func findMiddleMaxSubArray(A []int, start int, middle int, end int) SubArrResult {
	leftSum := int(math.Inf(-1))
	leftIndex := 0
	totalSum := 0

	for i := middle; i >= start; i-- {
		totalSum += A[i]

		if totalSum > leftSum {
			leftSum = totalSum
			leftIndex = i
		}
	}

	rightSum := int(math.Inf(-1))
	rightIndex := 0
	totalSum = 0

	for i := middle + 1; i <= end; i++ {
		totalSum += A[i]

		if totalSum > rightSum {
			rightSum = totalSum
			rightIndex = i
		}
	}

	result := SubArrResult{
		leftIndex:  leftIndex,
		rightIndex: rightIndex,
		sum:        leftSum + rightSum,
	}

	return result
}
