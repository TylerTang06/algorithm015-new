package week07

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	// dp[i] = dp[i-1] + dp[i-2]
	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		n2, n1 = n2+n1, n2
	}

	return n2
}
