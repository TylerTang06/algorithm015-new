package week09

import "strconv"

func numDecodings(s string) int {
	if s == "" || s[0] == '0' {
		return 0
	}

	// dp[i] = dp[i-2] if s[i] == '0'
	// else
	// dp[i] = dp[i-1] + dp[i-2]  if 10 <= s[i-1,i] <= 26
	// else
	// dp[i] = dp[i-1]
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}

		num, _ := strconv.Atoi(string(s[i-1]) + string(s[i]))
		if num <= 26 && num >= 10 {
			if i == 1 {
				dp[i]++
			} else {
				dp[i] += dp[i-2]
			}
		}
	}

	return dp[len(s)-1]
}
