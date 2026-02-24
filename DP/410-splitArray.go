package week06

import "math"

// dp, also can use binary search method
func splitArray(nums []int, m int) int {
	n := len(nums)
	if nums == nil || n < m {
		return math.MaxInt32
	}

	// init data structures
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	sub := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sub[i] = sub[i-1] + nums[i-1]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= minInt(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = minInt(dp[i][j], maxInt(dp[k][j-1], sub[i]-sub[k]))
			}
		}
	}

	return dp[n][m]
}

func minInt(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
