package week06

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// dp[i][j] = min{d[i-1][j], dp[i][j-1], dp[i-1][j-1]} + 1
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j]
					if dp[i][j] > dp[i][j-1] {
						dp[i][j] = dp[i][j-1]
					}
					if dp[i][j] > dp[i-1][j-1] {
						dp[i][j] = dp[i-1][j-1]
					}
				}
				dp[i][j]++
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max * max
}
