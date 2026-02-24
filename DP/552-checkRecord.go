package week06

// dp, it is totally diffrent
func checkRecord(n int) int {
	if n == 1 {
		return 3
	}

	// dp[i][j][k] 0<=j<=1 0<=k<=2 i=i-j-k
	dp00, dp10, dp01 := 1, 1, 1
	dp11, dp02, dp12 := 0, 0, 0
	mod := 1000000007
	for i := 2; i <= n; i++ {
		cur00, cur10, cur01 := dp00, dp10, dp01
		cur11, cur02, cur12 := dp11, dp02, dp12

		// +P
		// dp[i][0][0] = dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]
		// dp[i][1][0] = dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2]
		dp00 = (cur00 + cur01 + cur02) % mod
		dp10 = (cur10 + cur11 + cur12) % mod
		// +L
		// dp[i][0][1] = dp[i-1][0][0]
		// dp[i][1][1] = dp[i-1][1][0]
		// dp[i][0][2] = dp[i-1][0][1]
		// dp[i][1][2] = dp[i-1][1][1]
		dp01, dp11, dp02, dp12 = cur00, cur10, cur01, cur11
		// +A
		// dp[i][1][0] = dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]
		dp10 += (cur00 + cur01 + cur02) % mod
	}

	// res = dp[i][0][0] + dp[i][1][0] + dp[i][0][1] + dp[i][1][1] + dp[i][0][2] + dp[i][1][2] + dp[i][1][0]
	res := (dp00 + dp10 + dp01 + dp11 + dp02 + dp12) % mod

	return res
}
