package week06

import (
	"math"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1
	}

	res, rows, cols := math.MinInt64, len(matrix), len(matrix[0])
	for l := 0; l < cols; l++ {
		rowSum := make([]int, rows)
		for r := l; r < cols; r++ {
			for i := 0; i < rows; i++ {
				rowSum[i] += matrix[i][r]
			}
			if curMax := curMaxSumArr(rowSum, k); res < curMax {
				res = curMax
			}
			if res == k {
				return k
			}
		}
	}

	return res
}

func curMaxSumArr(arr []int, k int) int {
	sum := arr[0]
	max := sum
	// O(n) can reduce a lot of time
	for i := 1; i < len(arr); i++ {
		if sum > 0 {
			sum += arr[i]
		} else {
			sum = arr[i]
		}
		if max < sum {
			max = sum
		}
	}
	if max <= k {
		return max
	}

	max = math.MinInt64
	// O(n^2)
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > max && sum <= k {
				max = sum
			}
			if max == k {
				return k
			}
		}
	}

	return max
}
