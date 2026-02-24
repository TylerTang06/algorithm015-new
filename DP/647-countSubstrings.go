package week06

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}

	// dp[i,i] = true
	// dp[i,j] = dp[i+1,j-1] if s[i] == s[j]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	res := 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] == true {
				res++
			}
		}
	}

	return res + len(s)
}
