package week04

/*
74. Search a 2D Matrix
给你一个满足下述两条属性的 m x n 整数矩阵：

每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
*/
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
