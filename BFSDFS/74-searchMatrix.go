package week04

// binary search
// other idea, for example, narrow the range of searching
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[mid/n][mid%n] == target {
			return true
		}
		if matrix[mid/n][mid%n] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
